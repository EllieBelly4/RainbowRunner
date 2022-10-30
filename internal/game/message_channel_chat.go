package game

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/helpers"
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

	msgChannel := ClientMessageChannelSource(msgType)

	/**
	0x01 - World
	0x02 - Zone
	0x03 - Group
	0x04 - Tell
	0x05 - Market
	0x06 - Noob
	0x07 - PVP
	*/
	if msgChannel == ClientMessageChannelSourceZone ||
		msgChannel == ClientMessageChannelSourceGroup ||
		msgChannel == ClientMessageChannelSourceMarket ||
		msgChannel == ClientMessageChannelSourceWorld ||
		msgChannel == ClientMessageChannelSourceNoob ||
		msgChannel == ClientMessageChannelSourcePVP {
		return handleIndirectChatMessageSent(sendingPlayer, conn, reader, msgChannel)
	} else if msgChannel == ClientMessageChannelSourceTell {
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
		return sendUndeliveredMessageNotification(conn, UndeliveredMessageNotificationReasonTargetNotFound)
	}

	err := sendTell(player, message, target)

	if err != nil {
		return sendUndeliveredMessageNotification(conn, UndeliveredMessageNotificationReasonNoReason)
	}

	return nil
}

func sendTell(player *objects.RRPlayer, msg string, target *objects.RRPlayer) error {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ChatChannel))
	body.WriteByte(0x00) // Chat Message
	body.WriteByte(byte(MessageChannelSourceTell))
	body.WriteByte(0x02) // Unk - must not be 0
	// Sender name
	body.WriteCString(player.CurrentCharacter.Name)
	body.WriteCString(msg)
	helpers.WriteCompressedASimple(target.Conn, body)

	body.Clear()
	body.WriteByte(byte(messages.ChatChannel))
	body.WriteByte(0x00) // Chat Message
	body.WriteByte(byte(MessageChannelSourceTell2))
	body.WriteByte(0x01) // Unk - must not be 0
	// Sender name
	body.WriteCString(player.CurrentCharacter.Name)
	body.WriteCString(msg)
	helpers.WriteCompressedASimple(player.Conn, body)

	return nil
}

func handleIndirectChatMessageSent(player *objects.RRPlayer, conn *connections.RRConn, reader *byter.Byter, channel ClientMessageChannelSource) error {
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
		sendUndeliveredMessageNotification(conn, UndeliveredMessageNotificationReasonNoReason)
		return err
	}

	err = sendMessageToTargets(player, msg, severChannelSource, objects.Players.GetPlayers())

	if err != nil {
		return sendUndeliveredMessageNotification(conn, UndeliveredMessageNotificationReasonNoReason)
	}

	return nil
}

func sendUndeliveredMessageNotification(conn *connections.RRConn, reason UndeliveredMessageNotificationReasonString) error {
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

	helpers.WriteCompressedASimple(conn, body)

	return nil
}

func sendMessageToTargets(sendingPlayer *objects.RRPlayer, msg string, channel MessageChannelSource, players []*objects.RRPlayer) error {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ChatChannel))
	body.WriteByte(0x00) // Chat Message

	body.WriteByte(byte(channel))

	if channel != MessageChannelSourceGlobalAnnouncement {
		body.WriteByte(0x00) // Unk, if not 0 then text colour is white
		// Sender name
		body.WriteCString(sendingPlayer.CurrentCharacter.Name)
	}

	body.WriteCString(msg)

	// TODO only send to relevant players
	for _, player := range players {
		helpers.WriteCompressedASimple(player.Conn, body)
	}

	return nil
}
