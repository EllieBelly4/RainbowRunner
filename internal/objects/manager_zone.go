package objects

import (
	"RainbowRunner/internal/lua"
	"sync"
)

var Zones = NewZoneManager()

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
		m.Zones[zoneName].Init()
	}

	return m.Zones[zoneName]
}

func (m *ZoneManager) CreateZone(name string) *Zone {
	z := &Zone{
		Name:     name,
		entities: make(map[uint16]DRObject),
		players:  make(map[uint16]*RRPlayer),
	}

	z.Scripts = lua.GetScriptGroup("zones." + name)

	m.Zones[name] = z

	return z
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
