package game

import (
	"RainbowRunner/internal/byter"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

const msgBar = "=======================================================================\n"

type Channel byte

const (
	NoChannel Channel = iota
	Unk1
	Unk2
	UserChannel
	CharacterChannel
	Unk5
	ChatChannel
	ClientEntityChannel
	Unk8
	GroupChannel
	TradeChannel
	UnkB
	UnkC
	ZoneChannel
	UnkE
	PosseChannel
)

func handleChannelMessage(conn net.Conn, reader *byter.Byter, clientID uint32) {
	msgChan := reader.UInt8()    // Channel
	msgSubType := reader.UInt8() // Message Type

	switch Channel(msgChan) {
	case Unk2:
		body := byter.NewLEByter(make([]byte, 0, 1024))
		body.WriteByte(byte(Unk2)) // Character channel
		body.WriteByte(0x00)
		WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
	case CharacterChannel:
		switch CharacterMessage(msgSubType) {
		case CharacterConnected:
			handleCharacterConnected(conn, clientID)
		case CharacterPlay:
			handleCharacterPlay(conn, clientID)
		case CharacterGetList:
			sendCharacterList(conn, clientID)
		case CharacterCreate:
			handleCharacterCreate(conn, reader, clientID)
		default:
			noticeMessage("Unhandled chan %x msgSubType %x", msgChan, msgSubType)
		}
	case ClientEntityChannel:
		switch ClientEntityMessage(msgSubType) {
		//case ClientEntityUnk7:
		//	fmt.Printf()
		default:
			noticeMessage("Unhandled chan %x msgSubType %x", msgChan, msgSubType)
		}
	case GroupChannel:
		switch GroupChannelMessage(msgSubType) {
		case GroupConnected:
			handleGroupConnected(conn, clientID)
		}
	case ZoneChannel:
		switch ZoneChannelMessage(msgSubType) {
		case ZoneUnk6:
			handleZoneUnk6(conn, clientID)
		}
	case UserChannel:
		switch msgSubType {
		case 0x00: // Request rosters
			handleUserUnk0(conn, clientID)
		case 0x01: // Rosters response
			handleUserUnk1(conn, clientID)
		default:
			noticeMessage("Unhandled chan %x msgSubType %x", msgChan, msgSubType)
		}
	default:
		noticeMessage("Unhandled channel message %x", msgChan)
	}
}

func noticeMessage(s string, a ...interface{}) {
	msg := fmt.Sprintf(s, a...)
	log.Infof("\n%s%s\n%s", msgBar, msg, msgBar)
}

func addCreateComponent(body *byter.Byter, parentID uint16, componentID uint16, typeString string) {
	body.WriteByte(0x32)          // Create Component
	body.WriteUInt16(parentID)    // Parent Entity ID
	body.WriteUInt16(componentID) // Component ID
	body.WriteByte(0xFF)          // Unk
	body.WriteCString(typeString) // Component Type
}
