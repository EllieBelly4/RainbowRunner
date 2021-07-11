package objects

var Zones = NewZoneManager()

type Zone struct {
	Name     string
	Entities []DRObject
	Players  []*RRPlayer
}

func (z *Zone) AddEntity(entity DRObject) {
	z.Entities = append(z.Entities, entity)
}

func (z *Zone) AddPlayer(player *RRPlayer) {
	z.Players = append(z.Players, player)
}

func (z *Zone) Spawn(entity DRObject) {
	Entities.RegisterAll(nil, entity)
	z.Entities = append(z.Entities, entity)
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
		Entities: make([]DRObject, 0, 65535),
		Players:  make([]*RRPlayer, 0, 1024),
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
