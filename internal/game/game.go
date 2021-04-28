package game

import (
	"RainbowRunner/internal/byter"
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

		reader.UInt8()      // Message Type?
		reader.UInt8()      // Unk Dynamic
		reader.UInt16()     // Unk
		reader.UInt32()     // Unk
		reader.Bytes(3)     // Unk
		_ = reader.UInt8()  // Unk Message Static
		_ = reader.UInt32() // One Time Key
		reader.Bytes(1)     // Null

		//fmt.Printf("%x %x", unkMessageStatic, key)

		response := byter.NewLEByter(make([]byte, 0, 1024))
		response.WriteUInt32(0xFEEDBEE5)
		response.WriteUInt32(0xFEEDBEE5)
		response.WriteUInt32(0xFEEDBEE5)
		response.WriteUInt32(0xFEEDBEE5)
		response.WriteUInt32(0xFEEDBEE5)
		response.WriteUInt32(0xFEEDBEE5)
		response.WriteUInt32(0xFEEDBEE5)
		response.WriteUInt32(0xFEEDBEE5)
		response.WriteByte(0x00)

		conn.Write(response.Data())
		fmt.Printf("Sent: \n%s", hex.Dump(response.Data()))
		//parser.Parse(buf, read)
	}
}
