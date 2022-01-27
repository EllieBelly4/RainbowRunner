package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/helpers"
	"RainbowRunner/internal/logging"
	"RainbowRunner/internal/message"
	"RainbowRunner/pkg/byter"
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
	"sync"
)

var Players = NewPlayerManager()

type RRPlayer struct {
	Conn               *connections.RRConn
	CurrentCharacter   *Player
	Characters         []*Player
	Zone               *Zone
	ClientEntityWriter *ClientEntityWriter
	MessageQueue       *message.Queue
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
	m.Lock()
	defer m.Unlock()

	rrPlayer := &RRPlayer{
		Conn:               rrconn,
		ClientEntityWriter: NewClientEntityWriterWithByter(),
		MessageQueue:       message.NewQueue(),
	}

	m.Players[rrconn.Client.ID] = rrPlayer

	return rrPlayer
}

func (m *PlayerManager) OnDisconnect(id int) {
	m.Lock()
	defer m.Unlock()

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
	body := byter.NewByter(make([]byte, 0, 1024*1024))
	clientEntityWriter := NewClientEntityWriter(body)

	for _, player := range m.Players {
		clientEntitySend := false

		clientEntityWriter.BeginStream()

		for !player.MessageQueue.IsEmpty(message.QueueTypeClientEntity) {
			item := player.MessageQueue.Dequeue(message.QueueTypeClientEntity)
			body.Write(item.Data)

			if logging.LoggingOpts.LogSentMessages {
				if logIt, ok := logging.LoggingOpts.LogSentMessageTypes[item.OpType]; !ok || logIt {
					logrus.Info(fmt.Sprintf("Sent Message:\n%s", hex.Dump(item.Data.Data())))
				}
			}

			clientEntitySend = true
		}

		clientEntityWriter.EndStream()

		if clientEntitySend {
			helpers.WriteCompressedASimple(player.Conn, body)
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
