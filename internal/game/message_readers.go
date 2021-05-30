package game

import (
	"RainbowRunner/internal/byter"
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
)

func ReadCompressedE(reader *byter.Byter) *byter.Byter {
	reader.UInt24()            // Client ID
	pLength := reader.UInt32() // Packet length

	reader.UInt24()            // Unk previously sent by server I think
	reader.UInt16()            // Unk
	reader.UInt24()            // Unk
	bLength := reader.UInt32() // Body Length

	compressed := reader.Bytes(int(pLength - 12))
	cReader := bytes.NewReader(compressed)
	r, err := zlib.NewReader(cReader)

	if err != nil {
		panic(err)
	}

	uncompressed := make([]byte, bLength)

	_, err = r.Read(uncompressed)

	if err != nil && !errors.Is(io.EOF, err) {
		panic(err)
	}

	return byter.NewLEByter(uncompressed)
}

func ReadCompressedA(reader *byter.Byter, packetLength uint32) *byter.Byter {
	bLength := reader.UInt32() // Body Length
	compressed := reader.Bytes(int(packetLength - 7))
	cReader := bytes.NewReader(compressed)
	r, err := zlib.NewReader(cReader)

	if err != nil {
		panic(err)
	}

	uncompressed := make([]byte, bLength)

	_, err = r.Read(uncompressed)

	if err != nil && !errors.Is(io.EOF, err) {
		panic(err)
	}

	return byter.NewLEByter(uncompressed)
}

func Write6(clientID uint32, channel byte, afterChannel uint32, body *byter.Byter, conn net.Conn) {
	response := byter.NewLEByter(make([]byte, 0, 8))

	response.WriteByte(0x06)                     // Packet Type
	response.WriteUInt24(uint(clientID))         // Unk
	response.WriteUInt24(uint(len(body.Data()))) // Packet Size (max 100000h)
	response.WriteByte(channel)                  // Channel?
	response.WriteUInt24(uint(afterChannel))     // SubType?
	response.Write(body)

	written, err := conn.Write(response.Data())

	if err != nil || written == 0 {
		panic(err)
	}

	log.Info(fmt.Sprintf("Sent: \n%s", hex.Dump(response.Data())))
}

func WriteCompressed8(body *byter.Byter, conn net.Conn) {
	response := byter.NewLEByter(make([]byte, 0, 7))

	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(body.Data())
	w.Close()

	//log.Info(fmt.Sprintf("Compressed raw:\n%sas:\n%s", hex.Dump(body.Data()), hex.Dump(b.Bytes())))

	response.WriteByte(0x08)       // Packet Type
	response.WriteUInt24(0x313233) // Unk
	response.WriteUInt32(uint32(len(b.Bytes())) + 7)
	response.WriteUInt32(uint32(len(body.Data())))

	response.WriteBuffer(b)

	written, err := conn.Write(response.Data())

	if err != nil || written == 0 {
		panic(err)
	}

	log.Info(fmt.Sprintf("Sent: \n%s", hex.Dump(response.Data())))
}

func WriteCompressedA(clientID uint32, dest uint8, messageType uint8, body *byter.Byter, conn net.Conn) {
	response := byter.NewLEByter(make([]byte, 0, 7))

	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(body.Data())
	w.Close()

	response.WriteByte(0x0a)             // Packet Type
	response.WriteUInt24(uint(clientID)) // Unk
	response.WriteUInt32(uint32(len(b.Bytes())) + 7)
	// Source
	response.WriteByte(dest)
	response.WriteByte(messageType)
	response.WriteByte(0x00)
	response.WriteUInt32(uint32(len(body.Data())))

	response.WriteBuffer(b)

	written, err := conn.Write(response.Data())

	if err != nil || written == 0 {
		panic(err)
	}

	log.Info(fmt.Sprintf("Sent Compressed: \n%sCompressed raw body:\n%s", hex.Dump(response.Data()), hex.Dump(body.Data())))
}

func WriteMessage(msgType uint8, unk uint32, channel uint8, conn net.Conn, body *byter.Byter) {
	response := byter.NewLEByter(make([]byte, 0, 8))

	response.WriteByte(msgType)                  // Packet Type
	response.WriteUInt24(uint(unk))              // Unk
	response.WriteUInt24(uint(len(body.Data()))) // Packet Size (max 100000h)
	response.WriteByte(channel)                  // Unk, Channel?
	response.Write(body)

	written, err := conn.Write(response.Data())

	if err != nil || written == 0 {
		panic(err)
	}

	log.Info(fmt.Sprintf("Sent: \n%s", hex.Dump(response.Data())))
}
