package login

import (
	"RainbowRunner/internal/message"
	"RainbowRunner/internal/serverconfig"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

func StartLoginServer() {
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", serverconfig.Config.Network.LoginServerPort))

	if err != nil {
		panic(err)
	}

	defer func() {
		err := listen.Close()
		if err != nil {
			panic(err)
		}
	}()

	for {
		conn, err := listen.Accept()

		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	parser := message.NewAuthMessageParser(conn)
	buf := make([]byte, 1024*10)

	fmt.Println("Client connected")

	_, err := conn.Write([]byte{
		3, 0, // Length
		0, // Message Type
	})

	if err != nil {
		panic(err)
	}

	for {
		read, err := conn.Read(buf)

		if err != nil {
			log.Info(fmt.Sprintf("failed to read from connection: %e\n", err))
			break
		}

		parser.Parse(buf, read)
	}

	err = conn.Close()

	if err != nil {
		return
	}
}
