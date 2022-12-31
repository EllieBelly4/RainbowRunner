package objects

import (
	"RainbowRunner/internal/serverconfig"
	"RainbowRunner/pkg/events"
	"crypto/md5"
	"encoding/binary"
	log "github.com/sirupsen/logrus"
	"sync"
)

var Zones = NewZoneManager()

type ZoneManager struct {
	sync.RWMutex
	Zones map[string]*Zone
}

func (m *ZoneManager) GetOrCreateZone(zoneName string) *Zone {
	if _, ok := m.Zones[zoneName]; !ok {
		m.CreateZone(zoneName)
		m.Zones[zoneName].Init()

		return m.Zones[zoneName]
	}

	z := m.Zones[zoneName]

	if serverconfig.Config.ReinitialiseZonesOnEnter {
		z.ClearEntities()
		z.Init()
	}

	return z
}

func (m *ZoneManager) CreateZone(name string) *Zone {
	nameHashBytes := md5.Sum([]byte(name))
	nameHash := binary.LittleEndian.Uint32(nameHashBytes[:])

	z := NewZone(name, nameHash)

	m.Zones[name] = z

	return z
}

func (m *ZoneManager) Zone(s string) *Zone {
	return m.GetOrCreateZone(s)
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
		if !zone.Initialised() {
			continue
		}

		err := zone.Tick()

		if err != nil {
			log.Error(err)
		}
	}
}

func NewZoneManager() *ZoneManager {
	zm := &ZoneManager{
		Zones: make(map[string]*Zone),
	}

	events.RegisterHandler[PlayerEnteredZoneEvent](func(event PlayerEnteredZoneEvent) {
		event.Zone.OnPlayerEnter(event.Player)
	})

	return zm
}
