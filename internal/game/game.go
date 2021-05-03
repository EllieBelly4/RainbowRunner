package game

import (
	"RainbowRunner/internal/byter"
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"fmt"
	"net"
)

var blowfishKey = "[;',27h,'.]94-31==-%&@!^+]"

func StartGameServer() {
	listen, err := net.Listen("tcp", "0.0.0.0:2603")

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
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	//parser := message.NewParser(conn)
	buf := make([]byte, 1024*10)

	fmt.Println("Client connected to gameserver")

	for {
		read, err := conn.Read(buf)

		if err != nil {
			fmt.Printf("failed to read from connection: %e\n", err)
			break
		}

		fmt.Printf("(GameServer)Received: \n%s\n", hex.Dump(buf[0:read]))

		reader := byter.NewLEByter(buf[0:read])

		msgType := reader.UInt8() // Message Type?

		switch msgType {
		case 0x02:
			reader.UInt8()          // Unk Dynamic
			reader.UInt16()         // Unk
			size := reader.UInt24() // Size

			if size == 9 {
				reader.UInt8()      // Channel
				_ = reader.UInt32() // Unk Message Static
				_ = reader.UInt32() // One Time Key
				reader.Bytes(1)     // Null

				body := byter.NewLEByter(make([]byte, 0, 1024))

				body.WriteByte(0x03)
				WriteMessage(0x10, 0x262728, 0x0a, conn, body)

				body = byter.NewLEByter(make([]byte, 0, 1024))
				body.WriteUInt32(0x01020304)
				WriteCompressedA(0x00, 0x03, body, conn)
			} else {
				fmt.Printf("Ignoring short message 0x02 of length %d\n", size)
			}
		default:
			fmt.Printf("Unhandled message type %x\n", msgType)
		}
	}
}

func WriteCompressed8(body *byter.Byter, conn net.Conn) {
	response := byter.NewLEByter(make([]byte, 0, 7))

	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(body.Data())
	w.Close()

	//fmt.Printf("Compressed raw:\n%sas:\n%s", hex.Dump(body.Data()), hex.Dump(b.Bytes()))

	response.WriteByte(0x08)       // Packet Type
	response.WriteUInt24(0x313233) // Unk
	response.WriteUInt32(uint32(len(b.Bytes())) + 7)
	response.WriteUInt32(uint32(len(body.Data())))

	response.WriteBuffer(b)

	written, err := conn.Write(response.Data())

	if err != nil || written == 0 {
		panic(err)
	}

	fmt.Printf("Sent: \n%s", hex.Dump(response.Data()))
}

func WriteCompressedA(dest uint8, messageType uint8, body *byter.Byter, conn net.Conn) {
	response := byter.NewLEByter(make([]byte, 0, 7))

	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(body.Data())
	w.Close()

	//fmt.Printf("Compressed raw:\n%sas:\n%s", hex.Dump(body.Data()), hex.Dump(b.Bytes()))

	response.WriteByte(0x0a)       // Packet Type
	response.WriteUInt24(0x313233) // Unk
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

	fmt.Printf("Sent: \n%s", hex.Dump(response.Data()))
}

func WriteMessage(msgType uint8, unk uint, channel uint8, conn net.Conn, body *byter.Byter) {
	response := byter.NewLEByter(make([]byte, 0, 8))

	//Enum
	// 0x01 = disconnect me please for one reason or another
	// 0x02 = message
	//        [Unk] [Unk ]
	//        00    00 00\
	// 0x10 = works better?
	// 0x08 = compressed
	response.WriteByte(msgType)                  // Packet Type
	response.WriteUInt24(unk)                    // Unk
	response.WriteUInt24(uint(len(body.Data()))) // Packet Size (max 100000h)
	response.WriteByte(channel)                  // Unk, Channel?
	response.Write(body)

	written, err := conn.Write(response.Data())

	if err != nil || written == 0 {
		panic(err)
	}

	fmt.Printf("Sent: \n%s", hex.Dump(response.Data()))
}
