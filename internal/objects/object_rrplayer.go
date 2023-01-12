package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/message"
	"RainbowRunner/internal/serverconfig"
)

//go:generate go run ../../scripts/generatelua -type=RRPlayer
type RRPlayer struct {
	Conn               *connections.RRConn
	CurrentCharacter   *Player
	Characters         []*Player
	ClientEntityWriter *ClientEntityWriter
	MessageQueue       *message.Queue

	debugOptions *RRPlayerDebugOptions
}

//go:generate go run ../../scripts/generatelua -type=RRPlayerDebugOptions
type RRPlayerDebugOptions struct {
	SendMovementMessages bool
}

func (p *RRPlayer) SetDebugSendMovementMessages(b bool) {
	p.debugOptions.SendMovementMessages = b
}

func (p *RRPlayer) GetDebugSendMovementMessages() bool {
	return p.debugOptions.SendMovementMessages
}

func (p *RRPlayer) DebugOptions() *RRPlayerDebugOptions {
	return p.debugOptions
}

// Deprecated: Use p.CurrentCharacter.Zone instead
func (p *RRPlayer) Zone() *Zone {
	return p.CurrentCharacter.Zone
}

func NewRRPlayer(rrconn *connections.RRConn, cewriter *ClientEntityWriter, queue *message.Queue) *RRPlayer {
	defaultSendMovement := serverconfig.Config.SendMovementMessages

	return &RRPlayer{
		Conn:               rrconn,
		ClientEntityWriter: cewriter,
		MessageQueue:       queue,
		debugOptions: &RRPlayerDebugOptions{
			SendMovementMessages: defaultSendMovement,
		},
	}
}
