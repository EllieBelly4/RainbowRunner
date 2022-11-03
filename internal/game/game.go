package game

import (
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/helpers"
	"RainbowRunner/internal/objects"
	"RainbowRunner/pkg/byter"
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

var Connections = make(map[int]*connections.RRConn)

func StartGameServer() {
	objects.Entities = objects.NewEntityManager()

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

	go StartGameLoop()

	for {
		conn, err := listen.Accept()

		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	//parser := message.NewParser(conn)
	buf := make([]byte, 1024*10)

	fmt.Println("Client connected to gameserver")

	rrconn := &connections.RRConn{
		NetConn:     conn,
		IsConnected: true,
	}

	rrconn.Client = connections.NewRRConnClient(
		1,
		rrconn,
	)
	objects.Players.Register(rrconn)

	Connections[int(rrconn.Client.ID)] = rrconn

	defer func(conn net.Conn) {
		rrconn.IsConnected = false
		err := conn.Close()
		objects.Players.OnDisconnect(int(rrconn.Client.ID))
		if err != nil {
			panic(err)
		}
	}(conn)

	for {
		read, err := conn.Read(buf)

		if err != nil {
			log.Info(fmt.Sprintf("failed to read from connection: %e\n", err))
			break
		}

		//log.Info(fmt.Sprintf("(GameServer)Received: \n%s\n", hex.Dump(buf[0:read])))

		reader := byter.NewLEByter(buf[0:read])

		for reader.HasRemainingData() {
			readPacket(rrconn, reader)
		}
	}
}

func readPacket(conn *connections.RRConn, reader *byter.Byter) {
	msgType := reader.UInt8() // Message Type?

	if msgType == 0x0a {
		reader.UInt24()                 // Unk
		packetLength := reader.UInt32() // Packet Length
		reader.UInt8()
		msgTypeA := reader.UInt8()
		reader.UInt8()

		if config.Config.Logging.LogSmallAs {
			fmt.Printf("Received compressed A:\n%s\n", hex.Dump(reader.Data()))
		}

		reader = ReadCompressedA(reader, packetLength)

		if msgTypeA == 0x00 {
			reader.UInt8()      // Some type?
			_ = reader.UInt32() // One Time Key
			reader.Bytes(1)     // Null

			body := byter.NewLEByter(make([]byte, 0, 1024))

			body.WriteByte(0x03)
			WriteMessage(conn, 0x10, 0x0a, body)

			body = byter.NewLEByter(make([]byte, 0, 1024))
			// Unk
			body.WriteUInt24(0xB2B3B4)
			// Unk
			body.WriteByte(0x00)
			helpers.WriteCompressedA(conn, 0x00, 0x03, body)
		} else if msgTypeA == 0x02 {
			body := byter.NewLEByter(make([]byte, 0, 1024))
			helpers.WriteCompressedA(conn, 0x00, 0x02, body)
		} else {
			panic(fmt.Sprintf("Unhandled 0x0a message type %x", msgTypeA))
		}
	} else if msgType == 0x0e {
		msgReader := ReadCompressedE(reader)

		if config.Config.Logging.LogEMessages {
			log.Infof("Received E:\n%s", hex.Dump(msgReader.Buffer))
		}

		handleChannelMessage(conn, msgReader)
	} else if msgType == 0x06 {
		reader.UInt24() // Unk
		reader.UInt24() // Size
		reader.UInt8()
		reader.UInt24() // Client ID
		reader.UInt8()  // Channel?
		reader.UInt8()  // Sub type?
		reader.UInt24() // Unk

		handleChannelMessage(conn, reader)
	} else {
		fmt.Errorf("Unhandled message type %x\n", msgType)
	}
}
