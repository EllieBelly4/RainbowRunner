package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/pkg/byter"
	"sync"
)

var Zones = NewZoneManager()

type Zone struct {
	sync.RWMutex
	Name     string
	Entities map[uint16]DRObject
	Players  map[uint16]*RRPlayer
}

func (z *Zone) RemovePlayer(id int) {
	z.Lock()
	defer z.Unlock()

	delete(z.Players, uint16(id))

	toDelete := make([]uint16, 0, 1024)

	for index, entity := range z.Entities {
		if entity == nil || entity.RREntityProperties().OwnerID == id {
			toDelete = append(toDelete, index)
		}
	}

	for _, index := range toDelete {
		z.Entities[index].RREntityProperties().Zone = nil
		delete(z.Entities, index)
	}
}

func (z *Zone) AddEntity(entity DRObject) {
	z.setZone(entity)

	z.Entities[entity.RREntityProperties().ID] = entity
}

func (z *Zone) AddPlayer(player *RRPlayer) {
	z.Players[uint16(player.Conn.GetID())] = player
}

func (z *Zone) Spawn(entity DRObject) {
	Entities.RegisterAll(nil, entity)
	z.AddEntity(entity)
}

func (z *Zone) setZone(entities ...DRObject) {
	for _, entity := range entities {
		entity.RREntityProperties().Zone = z
		z.setZone(entity.Children()...)
	}
}

func (z *Zone) SendToAll(body *byter.Byter) {
	for _, player := range z.Players {
		connections.WriteCompressedASimple(player.Conn, body)
	}
}

type ZoneManager struct {
	Zones map[string]*Zone
}

func (m *ZoneManager) PlayerJoin(zoneName string, player *RRPlayer) {
	zone := m.getOrCreateZone(zoneName)

	player.Zone = zone

	for _, child := range player.CurrentCharacter.Children() {
		zone.AddEntity(child)
	}

	zone.AddPlayer(player)
}

func (m *ZoneManager) getOrCreateZone(zoneName string) *Zone {
	if _, ok := m.Zones[zoneName]; !ok {
		m.CreateZone(zoneName)
	}

	return m.Zones[zoneName]
}

func (m *ZoneManager) CreateZone(name string) {
	m.Zones[name] = &Zone{
		Name:     name,
		Entities: make(map[uint16]DRObject),
		Players:  make(map[uint16]*RRPlayer),
	}
}

func (m *ZoneManager) Spawn(npc *GCObject) {
	Entities.RegisterAll(nil, npc)

}

func (m *ZoneManager) Zone(s string) *Zone {
	return m.getOrCreateZone(s)
}

func NewZoneManager() *ZoneManager {
	return &ZoneManager{
		Zones: make(map[string]*Zone),
	}
}
