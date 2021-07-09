package game

import (
	byter "RainbowRunner/pkg/byter"
	"encoding/hex"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

const msgBar = "=======================================================================\n"

//go:generate stringer -type=Channel
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
		noticeMessage("unhandled channel %x\n%s", msgChan, hex.Dump(reader.Data()))
		return
	}

	fmt.Printf("<---- recv [%s-0x%x] len %d\n", Channel(msgChan).String(), msgSubType, len(reader.Buffer))

	err := handler(conn, msgSubType, reader)

	if err != nil {
		if errors.Is(err, UnhandledChannelMessageError) {
			noticeMessage("unhandled channel message chan: %x type: %x", msgChan, msgSubType)
		} else {
			fmt.Println(err)
		}
	}
}

func noticeMessage(s string, a ...interface{}) {
	msg := fmt.Sprintf(s, a...)
	fmt.Printf("\n%s%s\n%s", msgBar, msg, msgBar)
}

func addCreateComponent(body *byter.Byter, parentID uint16, componentID uint16, typeString string) {
	body.WriteByte(0x32)          // Create Component
	body.WriteUInt16(parentID)    // Parent Entity ID
	body.WriteUInt16(componentID) // Component ID
	body.WriteByte(0xFF)          // Unk
	body.WriteCString(typeString) // Component Type
	body.WriteByte(0x01)          // Unk
}
