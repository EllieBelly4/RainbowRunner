package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/helpers"
	"fmt"
	"sync"
)

var Players = NewPlayerManager()

type RRPlayer struct {
	Conn               *connections.RRConn
	CurrentCharacter   *Player
	Characters         []*Player
	Zone               *Zone
	ClientEntityWriter *ClientEntityWriter
}

type PlayerManager struct {
	sync.RWMutex
	Players map[int]*RRPlayer
}

func (m *PlayerManager) GetPlayers() []*RRPlayer {
	m.RLock()
	defer m.RUnlock()

	list := make([]*RRPlayer, 0)

	for _, entity := range m.Players {
		list = append(list, entity)
	}

	return list
}

func (m *PlayerManager) Register(rrconn *connections.RRConn) *RRPlayer {
	rrPlayer := &RRPlayer{
		Conn:               rrconn,
		ClientEntityWriter: NewClientEntityWriterWithByter(),
	}

	m.Players[rrconn.Client.ID] = rrPlayer

	return rrPlayer
}

func (m *PlayerManager) OnDisconnect(id int) {
	m.RLock()
	defer m.RUnlock()

	fmt.Printf("Player %d Disconnected\n", id)
	if player, ok := Players.Players[id]; ok {
		if player.Zone != nil {
			player.Zone.RemovePlayer(id)
		}
	}

	Entities.RemoveOwnedBy(id)

	delete(Players.Players, id)
}

func (m *PlayerManager) GetPlayer(id uint16) *RRPlayer {
	return m.Players[int(id)]
}

func (m *PlayerManager) AfterTick() {
	for _, player := range m.Players {
		if player.ClientEntityWriter.IsDirty() {
			player.ClientEntityWriter.EndStream()
			helpers.WriteCompressedASimple(player.Conn, player.ClientEntityWriter.GetBody())
		}

		player.ClientEntityWriter.Clear()
	}
}

func (m *PlayerManager) BeforeTick() {
	for _, player := range m.Players {
		player.ClientEntityWriter.BeginStream()
	}
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		Players: make(map[int]*RRPlayer),
	}
}
