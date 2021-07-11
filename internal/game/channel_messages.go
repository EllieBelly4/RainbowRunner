package game

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	byter "RainbowRunner/pkg/byter"
	"encoding/hex"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

const msgBar = "=======================================================================\n"

var UnhandledChannelMessageError = errors.New("unhandled channel message")

type ChannelMessageHandler func(conn *connections.RRConn, msgType byte, reader *byter.Byter) error

var channelMessageHandlers = map[messages.Channel]ChannelMessageHandler{
	messages.CharacterChannel:    handleCharacterChannelMessages,
	messages.Unk2:                handleUnk2ChannelMessages,
	messages.ClientEntityChannel: handleClientEntityChannelMessages,
	messages.GroupChannel:        handleGroupChannelMessages,
	messages.ZoneChannel:         handleZoneChannelMessages,
	messages.UserChannel:         handleUserChannelMessages,
}

func handleUnk2ChannelMessages(conn *connections.RRConn, msgType byte, reader *byter.Byter) error {
	log.Info("sending unknown response for Unk2 channel")
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.Unk2)) // Character channel
	body.WriteByte(0x00)
	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
	return nil
}

func handleChannelMessage(conn *connections.RRConn, reader *byter.Byter) {
	msgChan := reader.UInt8()   // Channel
	msgSubType := reader.Byte() // Message Type

	handler, ok := channelMessageHandlers[messages.Channel(msgChan)]

	if !ok {
		noticeMessage("unhandled channel %x\n%s", msgChan, hex.Dump(reader.Data()))
		return
	}

	fmt.Printf("<---- recv [%s-0x%x] len %d\n", messages.Channel(msgChan).String(), msgSubType, len(reader.Buffer))

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
