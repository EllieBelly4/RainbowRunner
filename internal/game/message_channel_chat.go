package game

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/objects"
	"RainbowRunner/pkg/byter"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func handleChatChannelMessages(conn *connections.RRConn, msgType byte, reader *byter.Byter) error {
	sendingPlayer := objects.Players.GetPlayer(uint16(conn.GetID()))

	if sendingPlayer == nil {
		return errors.New(fmt.Sprintf("could not find player sending chat message with ID: %d", conn.GetID()))
	}

	msgChannel := messages.ClientMessageChannelSource(msgType)

	/**
	0x01 - World
	0x02 - Zone
	0x03 - Group
	0x04 - Tell
	0x05 - Market
	0x06 - Noob
	0x07 - PVP
	*/
	if msgChannel == messages.ClientMessageChannelSourceZone ||
		msgChannel == messages.ClientMessageChannelSourceGroup ||
		msgChannel == messages.ClientMessageChannelSourceMarket ||
		msgChannel == messages.ClientMessageChannelSourceWorld ||
		msgChannel == messages.ClientMessageChannelSourceNoob ||
		msgChannel == messages.ClientMessageChannelSourcePVP {
		return handleIndirectChatMessageSent(sendingPlayer, conn, reader, msgChannel)
	} else if msgChannel == messages.ClientMessageChannelSourceTell {
		return handleDirectChatMessageSent(sendingPlayer, conn, reader)
	} else {
		return UnhandledChannelMessageError
	}
}

func handleDirectChatMessageSent(player *objects.RRPlayer, conn *connections.RRConn, reader *byter.Byter) error {
	msg := reader.CString()
	splitMsg := strings.Split(msg, " ")

	if len(splitMsg) < 2 {
		return nil
	}

	targetName := splitMsg[0]
	message := splitMsg[1]

	quoteRegex := regexp.MustCompile("['\"]")

	targetName = quoteRegex.ReplaceAllString(targetName, "")

	target := objects.Players.GetPlayerByCharacterName(targetName)

	if target == nil {
		return sendUndeliveredMessageNotification(conn, messages.UndeliveredMessageNotificationReasonTargetNotFound)
	}

	err := sendTell(player, message, target)

	if err != nil {
		return sendUndeliveredMessageNotification(conn, messages.UndeliveredMessageNotificationReasonNoReason)
	}

	return nil
}

func sendTell(player *objects.RRPlayer, msg string, target *objects.RRPlayer) error {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ChatChannel))
	body.WriteByte(0x00) // Chat Message
	body.WriteByte(byte(messages.MessageChannelSourceTell))
	body.WriteByte(0x02) // Unk - must not be 0
	// Sender name
	body.WriteCString(player.CurrentCharacter.Name)
	body.WriteCString(msg)
	connections.WriteCompressedASimple(target.Conn, body)

	body.Clear()
	body.WriteByte(byte(messages.ChatChannel))
	body.WriteByte(0x00) // Chat Message
	body.WriteByte(byte(messages.MessageChannelSourceTell2))
	body.WriteByte(0x01) // Unk - must not be 0
	// Sender name
	body.WriteCString(player.CurrentCharacter.Name)
	body.WriteCString(msg)
	connections.WriteCompressedASimple(player.Conn, body)

	return nil
}

func handleIndirectChatMessageSent(player *objects.RRPlayer, conn *connections.RRConn, reader *byter.Byter, channel messages.ClientMessageChannelSource) error {
	msg := reader.CString()

	// 0x00 Looks like message reading
	//
	// read 1 byte (A)
	// if A is 2 or 3 or 4 or 5 or 0x0B or 0x0C or 0x10:
	//   read 1 byte (B)
	// if A is not 0x0D // This is a not a global announcement
	//   read string and generate GCObject
	// read string // probably message
	//
	// 0x01
	// 0x02 Undelivered message notification
	// 0x03

	severChannelSource, err := channel.ToMessageChannelSource()

	if err != nil {
		sendUndeliveredMessageNotification(conn, messages.UndeliveredMessageNotificationReasonNoReason)
		return err
	}

	err = sendMessageToTargets(player, msg, severChannelSource, objects.Players.GetPlayers())

	if err != nil {
		return sendUndeliveredMessageNotification(conn, messages.UndeliveredMessageNotificationReasonNoReason)
	}

	return nil
}

func sendUndeliveredMessageNotification(conn *connections.RRConn, reason messages.UndeliveredMessageNotificationReasonString) error {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ChatChannel))
	body.WriteByte(0x02) // Undelivered message notification

	// Undelivered message notification string
	// 0x00 - "No Reason"
	/**
	Guessing:
	0x00 'No Reason'
	0x01 'Unknown Target Domain'
	0x02 'Unable To Find Sender'
	0x03 'Unauthorized Broadcast'
	0x04 'No Targets'
	0x05 'Unable To Find Target Session Id'
	0x06 'Target Not Logged In'
	0x07 'Target Is Ignoring Sender'
	0x08 'Target Not Found'
	0x09 'Chat System Unavailable'
	*/
	body.WriteByte(byte(reason))

	connections.WriteCompressedASimple(conn, body)

	return nil
}

func sendMessageToTargets(sendingPlayer *objects.RRPlayer, msg string, channel messages.MessageChannelSource, players []*objects.RRPlayer) error {
	chatMessage := messages.ChatMessage{
		Channel: channel,
		Unk0:    0x00,
		Message: msg,
		Sender:  sendingPlayer.CurrentCharacter.Name,
	}

	// TODO only send to relevant players
	for _, player := range players {
		player.Conn.SendMessage(chatMessage)
	}

	return nil
}
