package game

import (
	"RainbowRunner/internal/byter"
	"RainbowRunner/internal/objects"
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
					sendCharacterList(conn, clientID)
				case CharacterCreate:
					name := reader.String()
					class := reader.String()
					reader.UInt8() // Unk
					reader.UInt8() // Face
					reader.UInt8() // Hair
					reader.UInt8() // Hair Colour

					log.Infof("New character created %s (%s)", name, class)

					body := byter.NewLEByter(make([]byte, 0, 1024))
					body.WriteByte(byte(CharacterChannel)) // Character channel
					body.WriteByte(byte(CharacterCreate))

					// GCSerialisation version (0x13 == latest)
					body.WriteByte(0x13)

					//body.WriteByte(0x01) // Unk

					//body.WriteCString("Player") // Native type?
					body.WriteCString("Avatar") // Native type?
					//body.WriteUInt32(0x11121314)
					//body.WriteUInt32(0xDDEEFFEE)
					body.WriteUInt32(0x01)

					body.WriteCString("ID")
					// 6100
					//body.WriteUInt16(0xEEEE) // Unk
					//body.WriteByte(0x00) // Unk

					body.WriteUInt32(0x00000000) // GCObject count

					//body.WriteCString("Player")
					//body.WriteCString("Name")
					//body.WriteCString("Ellie")

					// GCObject Name, must be same(extend?) as base type
					body.WriteCString("Avatar")
					// GCObject property name
					// Properties are instances of GCNativeProperty
					body.WriteCString("Hair")
					body.WriteUInt32(0x00)
					body.WriteCString("HairColor")
					body.WriteUInt32(0x00)
					body.WriteCString("Name")
					body.WriteCString("Ellie")
					body.WriteCString("Face")
					body.WriteUInt32(0x01)
					body.WriteCString("FaceFeature")
					body.WriteUInt32(0x01)
					body.WriteCString("Skin")
					body.WriteUInt32(0x01)

					//
					// When this is 0x02 it does not check for type strings use section A otherwise use section B
					//
					//body.WriteByte(0x00) // Unk, potentially version again? or type, 0x02 = properties

					////////////////////////////////////////////////////////////////////////////////////////////
					///////////////////////////////// Section A          ///////////////////////////////////////
					////////////////////////////////////////////////////////////////////////////////////////////

					// This seems to be getting an existing entity by ID?

					// Using instead of 2 bytes below
					//body.WriteUInt32(0x00000001)

					////////////////////////////////////////////////////////////////////////////////////////////
					///////////////////////////////// Section B          ///////////////////////////////////////
					////////////////////////////////////////////////////////////////////////////////////////////

					// Entity
					//body.WriteCString("Player") // GCObject base type?
					//body.WriteCString("Hair")   // Property name

					// This works:
					//body.WriteCString("Avatar") // GCObject base type?
					//body.WriteCString("Hair")   // Property name

					//Property Names?:
					// Visible

					//body.WriteUInt32(0x86878889) // Unk

					////////////////////////////////////////////////////////////////////////////////////////////
					///////////////////////////////// Sections END       ///////////////////////////////////////
					////////////////////////////////////////////////////////////////////////////////////////////

					//charBytes := make([]byte, 1024)
					//length, err := hex.Decode(charBytes, []byte("61006176617461722e636c61737365732e466967687465724d616c65000000000000"))
					//
					//if err != nil {
					//	panic(err)
					//}

					//body.WriteBytes(charBytes[:length])

					WriteCompressedA(clientID, 0x01, 0x0f, body, conn)

					sendCharacterList(conn, clientID)
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

func sendCharacterList(conn net.Conn, clientID uint32) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(CharacterChannel)) // Character channel
	body.WriteByte(byte(CharacterGetList)) // Get character list (GotCharacter)

	player := objects.NewGCObject("Player")
	player.ID = 0xBABAFAAB
	player.Name = "Ellie"

	avatar := objects.NewGCObject("Avatar")
	avatar.ID = 0xBABAFAAC
	avatar.Name = "Avatar Name"
	avatar.Properties = []objects.GCObjectProperty{
		objects.Uint32Prop("Hair", 0x00),
		objects.Uint32Prop("HairColor", 0x00),
		objects.Uint32Prop("Face", 0x01),
		objects.Uint32Prop("FaceFeature", 0x01),
		objects.Uint32Prop("Skin", 0x01),
	}

	modifiers := objects.NewGCObject("Modifiers")
	modifiers.ID = 0xBABAFAAC
	modifiers.Name = "Mod Name"
	modifiers.Properties = []objects.GCObjectProperty{
		objects.Uint32Prop("IDGenerator", 0x01),
	}

	//avatar.AddChild(modifiers)
	player.AddChild(avatar)
	player.Serialise(body)

	body.WriteCString("Unk")  // Specific to player read
	body.WriteCString("Unk2") // Specific to player read

	avatar.AddChild(modifiers)
	avatar.Serialise(body)

	WriteCompressedA(clientID, 0x01, 0x0f, body, conn)

	//// GCSerialisation version (0x13 == latest)
	//body.WriteByte(0x13)
	//
	//////////////////////////////////////////////////////////////
	///////////////////// GCObject 1 header ////////////////////////
	//////////////////////////////////////////////////////////////
	//
	//body.WriteCString("Player")  // Native type?
	//body.WriteUInt32(0xBABAFAAB) // Potentially an ID?
	//body.WriteCString("Ellie")   // Name
	//body.WriteUInt32(0x00000001) // Child GCObject count?
	//
	//////////////////////////////////////////////////////////////
	///////////////////// GCObject 2 header ////////////////////////
	//////////////////////////////////////////////////////////////
	//body.WriteByte(0x13)
	//body.WriteCString("Avatar")
	//body.WriteUInt32(0xBABAFAAC) // Potentially an ID?
	//body.WriteCString("Ellie")   // Name
	//body.WriteUInt32(0x00000000) // Child GCObject count?
	//
	//////////////////// GCObject type ///////////////////////////
	//body.WriteCString("avatar") // GCObject name
	//
	//////////////////////////////////////////////////////////////
	//////////////////// GCObject 2 properties /////////////////////
	//////////////////////////////////////////////////////////////
	//body.WriteCString("Hair")
	//body.WriteUInt32(0x00)
	//body.WriteCString("HairColor")
	//body.WriteUInt32(0x00)
	//body.WriteCString("Face")
	//body.WriteUInt32(0x01)
	//body.WriteCString("FaceFeature")
	//body.WriteUInt32(0x01)
	//body.WriteCString("Skin")
	//body.WriteUInt32(0x01)
	//body.WriteNull()
	//
	//////////////////// GCObject 1 type ///////////////////////////
	//body.WriteCString("player")
	//
	//////////////////////////////////////////////////////////////
	//////////////////// GCObject 1 properties ///////////////////
	//////////////////////////////////////////////////////////////
	//body.WriteNull()
	//
	//////////////////////////////////////////////////////////////
	///////////////// Player specific properties//////////////////
	//////////////////////////////////////////////////////////////
	//body.WriteCString("Unk")
	//body.WriteCString("Unk2")
	//
	//////////////////////////////////////////////////////////////
	///////////////////// GCObject 3 header? /////////////////////
	//////////////////////////////////////////////////////////////
	//body.WriteByte(0x13)
	//body.WriteCString("Avatar")
	//body.WriteUInt32(0xBABAFAAC) // Potentially an ID?
	//body.WriteCString("Ellie")   // Name
	//body.WriteUInt32(0x00000001) // Child GCObject count?
	//
	//////////////////////////////////////////////////////////////
	///////////////////// GCObject 4 header /////////////////////
	//////////////////////////////////////////////////////////////
	//
	//body.WriteByte(0x13)
	//body.WriteCString("Modifiers")
	//body.WriteUInt32(0xBABAFAAC)  // Potentially an ID?
	//body.WriteCString("Mod Name") // Name
	//body.WriteUInt32(0x00000000)  // Child GCObject count?
	//
	//////////////////// GCObject type ///////////////////////////
	//
	//body.WriteCString("modifiers") // GCObject name
	//
	//////////////////////////////////////////////////////////////
	//////////////////// GCObject 4 properties /////////////////////
	//////////////////////////////////////////////////////////////
	//body.WriteCString("IDGenerator")
	//body.WriteUInt32(0x01)
	//body.WriteNull()
	//
	//////////////////// GCObject type ///////////////////////////
	//body.WriteCString("avatar") // GCObject name
	//
	//////////////////////////////////////////////////////////////
	//////////////////// GCObject 3 properties /////////////////////
	//////////////////////////////////////////////////////////////
	//body.WriteCString("Hair")
	//body.WriteUInt32(0x00)
	//body.WriteCString("HairColor")
	//body.WriteUInt32(0x00)
	//body.WriteCString("Face")
	//body.WriteUInt32(0x01)
	//body.WriteCString("FaceFeature")
	//body.WriteUInt32(0x01)
	//body.WriteCString("Skin")
	//body.WriteUInt32(0x01)
	//body.WriteNull()
	//
	//WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
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
