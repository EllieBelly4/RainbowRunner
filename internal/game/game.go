package game

import (
	"RainbowRunner/internal/byter"
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

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

	// We are receiving mixed messages, 0x0a + 0x0e
	// Need to support splitting now
	//time=2021-05-30T12:58:21+01:00 level=info msg=(GameServer)Received:
	//00000000  0a dd 7b 00 0f 00 00 00  00 02 00 00 00 00 00 78  |..{............x|
	//00000010  9c 03 00 00 00 00 01 0e  dd 7b 00 19 00 00 00 b4  |.........{......|
	//00000020  b3 b2 01 00 01 0f 00 05  00 00 00 78 9c 63 67 61  |...........x.cga|
	//00000030  62 10 05 00 00 53 00 23                           |b....S.#|

	for {
		read, err := conn.Read(buf)

		if err != nil {
			log.Info(fmt.Sprintf("failed to read from connection: %e\n", err))
			break
		}

		log.Info(fmt.Sprintf("(GameServer)Received: \n%s\n", hex.Dump(buf[0:read])))

		reader := byter.NewLEByter(buf[0:read])

		msgType := reader.UInt8() // Message Type?

		if msgType == 0x0a {
			clientID = reader.UInt24() // Unk
			reader.UInt32()            // Packet Length
			reader.UInt8()
			msgTypeA := reader.UInt8()
			reader.UInt8()
			reader = ReadCompressedA(reader)

			log.Infof("Uncompressed A:\n%s", hex.Dump(reader.Buffer))

			if msgTypeA == 0x00 {
				reader.UInt8()      // Some type?
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
			} else if msgTypeA == 0x02 {
			} else {
				log.Panicf("Unhandled 0x0a message type %x", msgTypeA)
			}
		} else if msgType == 0x06 || msgType == 0x0e {
			if msgType == 0x0e {
				first := true
				// TODO when we can reliably get the length of all packets we can ensure this is done for everything
				for reader.HasRemainingData() {
					if !first {
						reader.UInt8()
					}
					msgReader := ReadCompressedE(reader)

					log.Infof("Uncompressed E:\n%s", hex.Dump(msgReader.Buffer))

					handleChannelMessage(conn, msgReader, clientID)
					first = false
				}
			} else {
				reader.UInt24() // Unk
				reader.UInt24() // Size
				reader.UInt8()
				reader.UInt24() // Client ID
				reader.UInt8()  // Channel?
				reader.UInt8()  // Sub type?
				reader.UInt24() // Unk

				handleChannelMessage(conn, reader, clientID)
			}
		} else {
			log.Info(fmt.Sprintf("Unhandled message type %x\n", msgType))
		}
	}
}

func sendCharacterList(conn net.Conn, clientID uint32) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(CharacterChannel)) // Character channel
	body.WriteByte(byte(CharacterGetList)) // Get character list (GotCharacter)

	count := 0x01

	body.WriteByte(byte(count))

	for i := 0; i < count; i++ {
		body.WriteUInt32(uint32(i + 1)) // ID?
		sendPlayer(body)
	}

	WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
}
