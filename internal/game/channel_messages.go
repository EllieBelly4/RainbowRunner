package game

import (
	"RainbowRunner/internal/byter"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
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

func handleChannelMessage(conn *RRConn, reader *byter.Byter) {
	msgChan := reader.UInt8()    // Channel
	msgSubType := reader.UInt8() // Message Type

	switch Channel(msgChan) {
	case Unk2:
		body := byter.NewLEByter(make([]byte, 0, 1024))
		body.WriteByte(byte(Unk2)) // Character channel
		body.WriteByte(0x00)
		WriteCompressedA(conn, 0x01, 0x0f, body)
	case CharacterChannel:
		err := handleCharacterChannelMessages(conn, reader, CharacterMessage(msgSubType), msgChan)

		if err != nil {
			noticeMessage("Unhandled chan %x msgSubType %x", msgChan, msgSubType)
		}
	case ClientEntityChannel:
		switch ClientEntityMessage(msgSubType) {
		case ClientEntityUnk4:
			handleClientEntityUnk4(conn, reader)
		//case ClientEntityUnk7:
		//	fmt.Printf()
		default:
			noticeMessage("Unhandled chan %x msgSubType %x", msgChan, msgSubType)
		}
	case GroupChannel:
		switch GroupChannelMessage(msgSubType) {
		case GroupConnected:
			handleGroupConnected(conn)
		}
	case ZoneChannel:
		handleZoneMessages(conn, msgSubType, reader)
	case UserChannel:
		switch msgSubType {
		case 0x00: // Request rosters
			handleUserUnk0(conn)
		case 0x01: // Rosters response
			handleUserUnk1(conn)
		default:
			noticeMessage("Unhandled chan %x msgSubType %x", msgChan, msgSubType)
		}
	default:
		noticeMessage("Unhandled channel message %x", msgChan)
	}
}

func handleCharacterChannelMessages(conn *RRConn, reader *byter.Byter, msgSubType CharacterMessage, msgChan uint8) error {
	switch msgSubType {
	case CharacterConnected:
		handleCharacterConnected(conn)
	case CharacterPlay:
		handleCharacterPlay(conn)
	case CharacterGetList:
		sendCharacterList(conn)
	case CharacterCreate:
		handleCharacterCreate(conn, reader)
	default:
		return errors.New("unhandled")
	}

	return nil
}

func handleZoneMessages(conn *RRConn, msgSubType uint8, reader *byter.Byter) {
	switch ZoneChannelMessage(msgSubType) {
	case ZoneUnk6:
		handleZoneUnk6(conn)
	case ZoneUnk8:
		body := byter.NewLEByter(make([]byte, 0, 1024))
		body.WriteByte(byte(ZoneChannel))
		body.WriteByte(0x08)
		WriteCompressedA(conn, 0x01, 0x0f, body)
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
