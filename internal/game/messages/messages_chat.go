package messages

import (
	"RainbowRunner/pkg/byter"
)

type ChatMessage struct {
	Channel MessageChannelSource
	Unk0    int
	Message string
	Sender  string
}

func (c ChatMessage) Write(b *byter.Byter) {
	b.WriteByte(byte(ChatChannel))
	b.WriteByte(0x00) // Chat Message

	b.WriteByte(byte(c.Channel))

	if c.Channel != MessageChannelSourceGlobalAnnouncement {
		b.WriteByte(0x00) // Unk, if not 0 then text colour is white
		// Sender name
		b.WriteCString(c.Sender)
	}

	b.WriteCString(c.Message)
}
