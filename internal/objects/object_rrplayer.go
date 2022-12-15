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

func (p *RRPlayer) Zone() *Zone {
	return p.CurrentCharacter.Zone
}

func (p *RRPlayer) OnZoneJoin() {
	p.CurrentCharacter.OnZoneJoin()
}

func (p *RRPlayer) LeaveZone() {
	p.CurrentCharacter.LeaveZone()
}

func (p *RRPlayer) JoinZone(zone *Zone) {
	p.CurrentCharacter.ChangeZone(zone.Name)
}
