package admin

import (
	"RainbowRunner/internal/game"
	byter "RainbowRunner/pkg/byter"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type CommandRequest struct {
	Type    string `json:"type"`
	Data    string `json:"data"`
	Channel string `json:"channel"`
	Command string `json:"command"`
}

func HandleRequest(w http.ResponseWriter, req *http.Request) {
	buf := make([]byte, req.ContentLength)

	_, err := req.Body.Read(buf)

	headers := w.Header()
	headers.Set("Access-Control-Allow-Origin", "*")

	if err != nil && !errors.Is(err, io.EOF) {
		fmt.Printf("[commander][error] invalid request\n")
		w.WriteHeader(500)
		return
	}

	commandRequest := &CommandRequest{}

	err = json.Unmarshal(buf, commandRequest)

	if err != nil {
		fmt.Printf("[commander][error] invalid request body\n")
		w.WriteHeader(500)
		return
	}

	if commandRequest.Type == "hex" {
		str := commandRequest.Data
		hexData, err := game.CommandStringToBytes(str)

		if err != nil {
			fmt.Printf("[commander][error] invalid command string\n")
			w.WriteHeader(500)
			return
		}

		var conn *game.RRConn
		var ok bool

		if conn, ok = game.Connections[0]; !ok {
			fmt.Printf("[commander][error] no clients\n")
			w.WriteHeader(500)
			return
		}

		var buf []byte

		channel, err := strconv.Atoi(commandRequest.Channel)

		if conn, ok = game.Connections[0]; !ok {
			fmt.Printf("[commander][error] invalid channel %v\n", channel)
			w.WriteHeader(500)
			return
		}

		if channel == 0 {
			buf = make([]byte, 0)
		} else {
			buf = []byte{byte(channel)}
		}

		fmt.Printf("[commander] executing command\n")

		body := byter.NewLEByter(buf)
		body.WriteBytes(hexData)
		game.WriteCompressedA(conn, 0x01, 0x0f, body)
	} else if commandRequest.Type == "cmd" {
		SendCommand(commandRequest.Command)
	}

	w.WriteHeader(200)
}

func SendCommand(command string) {

}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func StartAdminServer() {
	http.HandleFunc("/command", HandleRequest)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
