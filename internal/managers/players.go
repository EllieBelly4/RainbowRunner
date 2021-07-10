package managers

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/objects"
)

var Players = NewPlayerManager()

type RRPlayer struct {
	Conn             *connections.RRConn
	CurrentCharacter *objects.Player
	Characters       []*objects.Player
}

type PlayerManager struct {
	Players map[int]*RRPlayer
}

func (m *PlayerManager) Register(rrconn *connections.RRConn) *RRPlayer {
	rrPlayer := &RRPlayer{
		Conn: rrconn,
	}

	m.Players[rrconn.Client.ID] = rrPlayer

	return rrPlayer
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		Players: make(map[int]*RRPlayer),
	}
}
