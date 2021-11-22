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
	entities map[uint16]DRObject
	players  map[uint16]*RRPlayer
}

func (z *Zone) Entities() []DRObject {
	z.RLock()
	defer z.RUnlock()

	l := make([]DRObject, 0)

	for _, drObject := range z.entities {
		l = append(l, drObject)
	}

	return l
}

func (z *Zone) Players() []*RRPlayer {
	z.RLock()
	defer z.RUnlock()

	l := make([]*RRPlayer, 0)

	for _, player := range z.players {
		l = append(l, player)
	}

	return l
}

func (z *Zone) RemovePlayer(id int) {
	z.Lock()
	defer z.Unlock()

	delete(z.players, uint16(id))

	toDelete := make([]uint16, 0, 1024)

	for index, entity := range z.entities {
		if entity == nil || entity.RREntityProperties().OwnerID == id {
			toDelete = append(toDelete, index)
		}
	}

	for _, index := range toDelete {
		z.entities[index].RREntityProperties().Zone = nil
		delete(z.entities, index)
	}
}

func (z *Zone) AddEntity(entity DRObject) {
	z.setZone(entity)

	z.entities[entity.RREntityProperties().ID] = entity
}

func (z *Zone) AddPlayer(player *RRPlayer) {
	z.players[uint16(player.Conn.GetID())] = player
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
	for _, player := range z.players {
		connections.WriteCompressedASimple(player.Conn, body)
	}
}

type ZoneManager struct {
	sync.RWMutex
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
		entities: make(map[uint16]DRObject),
		players:  make(map[uint16]*RRPlayer),
	}
}

func (m *ZoneManager) Spawn(npc *GCObject) {
	Entities.RegisterAll(nil, npc)

}

func (m *ZoneManager) Zone(s string) *Zone {
	return m.getOrCreateZone(s)
}

func (m *ZoneManager) GetZones() []*Zone {
	m.RLock()
	defer m.RUnlock()

	list := make([]*Zone, 0)

	for _, zone := range m.Zones {
		list = append(list, zone)
	}

	return list
}

func NewZoneManager() *ZoneManager {
	return &ZoneManager{
		Zones: make(map[string]*Zone),
	}
}
