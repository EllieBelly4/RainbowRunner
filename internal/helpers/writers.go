package helpers

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/logging"
	"RainbowRunner/pkg/byter"
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
)

func WriteCompressedASimple(conn connections.Connection, b *byter.Byter) {
	pc, file, line, ok := runtime.Caller(1)
	callerInfo := "unk"

	if logging.LoggingOpts.LogGenericSent {
		if ok {
			details := runtime.FuncForPC(pc)
			callerInfo = fmt.Sprintf("%s() %s:%d", details.Name(), file, line)
		}

		logrus.Infof("WriteCompressedA: %s", callerInfo)
	}

	WriteCompressedA(conn, 0x01, 0x0f, b)
}

func WriteCompressedA(conn connections.Connection, dest uint8, messageType uint8, body *byter.Byter) {
	response := byter.NewLEByter(make([]byte, 0, 7))

	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(body.Data())
	w.Close()

	response.WriteByte(0x0a)                 // Packet Type
	response.WriteUInt24(uint(conn.GetID())) // Unk
	response.WriteUInt32(uint32(len(b.Bytes())) + 7)
	// Source
	response.WriteByte(dest)        // Unk
	response.WriteByte(messageType) // Unk
	response.WriteByte(0x00)
	response.WriteUInt32(uint32(len(body.Data())))

	if len(body.Buffer) >= 2 && logging.LoggingOpts.LogGenericSent {
		fmt.Printf(">>>>> send [%s-%d] len %d\n", messages.Channel(body.Data()[0]).String(), body.Data()[1], len(body.Buffer))
	} else {
		//fmt.Printf(">>>>> send [nochannel] len %d\n", len(body.Buffer))
	}

	response.WriteBuffer(b)

	err := conn.Send(response)

	if err != nil {
		fmt.Println(err)
		return
	}

	if logging.LoggingOpts.LogGenericSent {
		pc, file, line, ok := runtime.Caller(1)
		callerInfo := "unk"

		if ok {
			details := runtime.FuncForPC(pc)
			callerInfo = fmt.Sprintf("%s() %s:%d", details.Name(), file, line)
		}

		logrus.Info(fmt.Sprintf("Sent: %s\n%s", callerInfo, hex.Dump(body.Data())))
	}
}
