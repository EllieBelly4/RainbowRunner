package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/message"
)

//go:generate go run ../../scripts/generatelua -type=RRPlayer
type RRPlayer struct {
	Conn               *connections.RRConn
	CurrentCharacter   *Player
	Characters         []*Player
	ClientEntityWriter *ClientEntityWriter
	MessageQueue       *message.Queue
}

// Deprecated: Use p.CurrentCharacter.Zone instead
func (p *RRPlayer) Zone() *Zone {
	return p.CurrentCharacter.Zone
}
