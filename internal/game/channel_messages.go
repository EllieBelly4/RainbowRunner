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

var UnhandledChannelMessageError = errors.New("unhandled channel message")

type ChannelMessageHandler func(conn *RRConn, msgType byte, reader *byter.Byter) error

var channelMessageHandlers = map[Channel]ChannelMessageHandler{
	CharacterChannel:    handleCharacterChannelMessages,
	Unk2:                handleUnk2ChannelMessages,
	ClientEntityChannel: handleClientEntityChannelMessages,
	GroupChannel:        handleGroupChannelMessages,
	ZoneChannel:         handleZoneChannelMessages,
	UserChannel:         handleUserChannelMessages,
}

func handleUnk2ChannelMessages(conn *RRConn, msgType byte, reader *byter.Byter) error {
	log.Info("sending unknown response for Unk2 channel")
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(Unk2)) // Character channel
	body.WriteByte(0x00)
	WriteCompressedA(conn, 0x01, 0x0f, body)
	return nil
}

func handleChannelMessage(conn *RRConn, reader *byter.Byter) {
	msgChan := reader.UInt8()   // Channel
	msgSubType := reader.Byte() // Message Type

	handler, ok := channelMessageHandlers[Channel(msgChan)]

	if !ok {
		noticeMessage("unhandled channel %x", msgChan)
		return
	}

	err := handler(conn, msgSubType, reader)

	if err != nil {
		if errors.Is(err, UnhandledChannelMessageError) {
			noticeMessage("unhandled channel message chan: %x type: %x", msgChan, msgSubType)
		} else {
			log.Error(err)
		}
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
