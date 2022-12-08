package objects

import (
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/types"
	"crypto/md5"
	"encoding/binary"
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
		zone.AddEntity(types.UInt16(uint16(player.Conn.GetID())), child)
	}

	zone.AddPlayer(player)
}

func (m *ZoneManager) getOrCreateZone(zoneName string) *Zone {
	if _, ok := m.Zones[zoneName]; !ok {
		m.CreateZone(zoneName)
		m.Zones[zoneName].Init()

		return m.Zones[zoneName]
	}

	z := m.Zones[zoneName]

	if config.Config.ReinitialiseZonesOnEnter {
		z.ClearEntities()
		z.Init()
	}

	return z
}

func (m *ZoneManager) CreateZone(name string) *Zone {
	nameHashBytes := md5.Sum([]byte(name))
	nameHash := binary.LittleEndian.Uint32(nameHashBytes[:])

	z := &Zone{
		Name:     name,
		ID:       nameHash,
		entities: make(map[uint16]DRObject),
		players:  make(map[uint16]*RRPlayer),
	}

	m.Zones[name] = z

	return z
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

func (m *ZoneManager) Tick() {
	m.RLock()
	defer m.RUnlock()

	for _, zone := range m.Zones {
		zone.Tick()
	}
}

func NewZoneManager() *ZoneManager {
	return &ZoneManager{
		Zones: make(map[string]*Zone),
	}
}
