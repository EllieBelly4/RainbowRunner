package game

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/helpers"
	"RainbowRunner/internal/objects"
	"RainbowRunner/pkg/byter"
)

func handleChatChannelMessages(conn *connections.RRConn, msgType byte, reader *byter.Byter) error {
	if msgType == 0x02 ||
		msgType == 0x03 ||
		msgType == 0x04 ||
		msgType == 0x0B ||
		msgType == 0x0C {
		return handleIndirectChatMessageSent(conn, reader, msgType)
	} else {
		return UnhandledChannelMessageError
	}
}

func handleIndirectChatMessageSent(conn *connections.RRConn, reader *byter.Byter, channel byte) error {
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

	err := sendMessageToTargets(msg, objects.Players.GetPlayers())

	if err != nil {
		return sendUndeliveredMessageNotification(conn)
	}

	return nil
}

func sendUndeliveredMessageNotification(conn *connections.RRConn) error {
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
	body.WriteByte(0x09)

	helpers.WriteCompressedASimple(conn, body)

	return nil
}

func sendMessageToTargets(msg string, players []*objects.RRPlayer) error {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ChatChannel))
	body.WriteByte(0x00) // Chat Message

	// 0x02 - World
	// 0x03 - Local
	// 0x04 - Group
	// 0x05 - Tell
	// 0x06 - Part of Tell? Sends back "To {NAME}" e.g. "To Testy" in pink
	// 0x0B - Market
	// 0x0C - Noob
	// 0x0D - Global Announcement
	messageChatChannelSource := byte(0x0C)

	body.WriteByte(messageChatChannelSource) // Unk

	if messageChatChannelSource != 0x0D {
		body.WriteByte(0x00) // Unk
		// Sender name
		body.WriteCString("Testy")
	}

	body.WriteCString(msg)

	// TODO only send to relevant players
	for _, player := range players {
		helpers.WriteCompressedASimple(player.Conn, body)
	}

	return nil
}
