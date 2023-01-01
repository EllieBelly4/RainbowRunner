package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/message"
	"RainbowRunner/internal/serverconfig"
	"RainbowRunner/pkg/byter"
	"encoding/hex"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
)

var Players = NewPlayerManager()

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

	m.Players[int(rrconn.Client.ID)] = rrPlayer

	return rrPlayer
}

func (m *PlayerManager) OnDisconnect(id int) {
	m.Lock()
	defer m.Unlock()

	fmt.Printf("Player %d Disconnected\n", id)
	if player, ok := Players.Players[id]; ok {
		if player.CurrentCharacter != nil && player.CurrentCharacter.Zone != nil {
			player.CurrentCharacter.Zone.RemovePlayer(id)
		}
	}

	//Entities.RemoveOwnedBy(id)

	delete(Players.Players, id)
}

func (m *PlayerManager) GetPlayerByCharacterName(name string) *RRPlayer {
	m.RLock()
	defer m.RUnlock()
	for _, player := range m.Players {
		if strings.ToLower(player.CurrentCharacter.Name) == strings.ToLower(name) {
			return player
		}
	}

	return nil
}

func (m *PlayerManager) GetPlayer(id uint16) *RRPlayer {
	m.RLock()
	defer m.RUnlock()

	return m.Players[int(id)]
}

func (m *PlayerManager) GetPlayerOrNil(id uint16) *RRPlayer {
	m.RLock()
	defer m.RUnlock()

	player, ok := m.Players[int(id)]

	if !ok {
		return nil
	}

	return player
}

var playerSendBuffer = byter.NewByter(make([]byte, 1024*1024+10))

func (m *PlayerManager) AfterTick() {
	playerSendBuffer.Clear()

	for _, player := range m.Players {
		if player.CurrentCharacter == nil {
			continue
		}

		if !player.CurrentCharacter.Spawned {
			player.MessageQueue.Clear(message.QueueTypeClientEntity)
			continue
		}

		player.ClientEntityWriter.Clear()

		clientEntitySend := false

		player.ClientEntityWriter.BeginStream()

		for !player.MessageQueue.IsEmpty(message.QueueTypeClientEntity) {
			item := player.MessageQueue.Dequeue(message.QueueTypeClientEntity)
			playerSendBuffer.Write(item.Data)

			if serverconfig.Config.Logging.LogFilterMessages {
				if logIt, ok := serverconfig.Config.Logging.LogSentMessageTypes[strings.ToLower(item.OpType.String())]; ok && logIt {
					logrus.Info(fmt.Sprintf("Sent Message:\n%s", hex.Dump(item.Data.Data())))
				}
			}

			clientEntitySend = true
		}

		player.ClientEntityWriter.EndStream()

		if clientEntitySend {
			connections.WriteCompressedASimple(player.Conn, playerSendBuffer)
		}
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
