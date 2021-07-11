package objects

import (
	"RainbowRunner/internal/connections"
	"fmt"
)

var Players = NewPlayerManager()

type RRPlayer struct {
	Conn             *connections.RRConn
	CurrentCharacter *Player
	Characters       []*Player
	Zone             *Zone
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

func (m *PlayerManager) OnDisconnect(id int) {
	fmt.Printf("Player %d Disconnected\n", id)
	if player, ok := Players.Players[id]; ok {
		if player.Zone != nil {
			player.Zone.RemovePlayer(id)
		}
	}

	Entities.RemoveOwnedBy(id)
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		Players: make(map[int]*RRPlayer),
	}
}
