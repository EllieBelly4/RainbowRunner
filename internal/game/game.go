package game

import (
	"RainbowRunner/internal/byter"
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

type Channel byte

const (
	NoChannel Channel = iota
	Unk1
	Unk2
	Unk3
	CharacterChannel
)

type CharacterMessage byte

const (
	CharacterConnected CharacterMessage = iota
	CharacterDisconnected
	CharacterCreate
	CharacterGetList
	CharacterDelete
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
	var clientID uint32 = 0

	for {
		read, err := conn.Read(buf)

		if err != nil {
			log.Info(fmt.Sprintf("failed to read from connection: %e\n", err))
			break
		}

		log.Info(fmt.Sprintf("(GameServer)Received: \n%s\n", hex.Dump(buf[0:read])))

		reader := byter.NewLEByter(buf[0:read])

		msgType := reader.UInt8() // Message Type?

		switch msgType {
		case 0x02:
			clientID = reader.UInt24() // Unk
			size := reader.UInt24()    // Size

			if size == 9 {
				reader.UInt8()      // Channel
				_ = reader.UInt32() // Unk Message Static
				_ = reader.UInt32() // One Time Key
				reader.Bytes(1)     // Null

				body := byter.NewLEByter(make([]byte, 0, 1024))

				body.WriteByte(0x03)
				WriteMessage(0x10, clientID, 0x0a, conn, body)

				body = byter.NewLEByter(make([]byte, 0, 1024))
				// Unk
				body.WriteUInt24(0xB2B3B4)
				// Unk
				body.WriteByte(0x00)
				WriteCompressedA(clientID, 0x00, 0x03, body, conn)
			} else {
				log.Info(fmt.Sprintf("Ignoring short message 0x02 of length %d\n", size))
			}
		case 0x06:
			reader.UInt24() // Unk
			reader.UInt24() // Size
			reader.UInt8()
			reader.UInt24()              // Client ID
			reader.UInt8()               // Channel?
			reader.UInt8()               // Sub type?
			reader.UInt24()              // Unk
			msgChan := reader.UInt8()    // Channel
			msgSubType := reader.UInt8() // Message Type

			switch Channel(msgChan) {
			case CharacterChannel:
				switch CharacterMessage(msgSubType) {
				case CharacterConnected:
					body := byter.NewLEByter(make([]byte, 0, 1024))
					body.WriteByte(byte(CharacterChannel))   // Character channel
					body.WriteByte(byte(CharacterConnected)) // Connected
					WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
				case CharacterGetList:
					body := byter.NewLEByter(make([]byte, 0, 1024))
					body.WriteByte(byte(CharacterChannel)) // Character channel
					body.WriteByte(byte(CharacterGetList)) // Get character list (GotCharacter)
					body.WriteByte(0x00)                   // Number of characters
					WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
				default:
					log.Warnf("Unhandled msgSubType %x", msgSubType)
				}
			default:
				log.Warnf("Unhandled channel message %x", msgChan)
			}

			//body := byter.NewLEByter(make([]byte, 0, 1024))
			//body.WriteUInt16(0xB3B4)
			//body.WriteUInt16(0xACDC)
			//body.WriteUInt16(0xF00D)
			//body.WriteByte(0xB0)
			//Write6(0x0a, 0x01, body, conn)
		default:
			log.Info(fmt.Sprintf("Unhandled message type %x\n", msgType))
		}
	}
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
