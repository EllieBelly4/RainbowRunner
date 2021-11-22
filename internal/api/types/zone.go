package types

import "RainbowRunner/internal/objects"

type ZoneCollection struct {
	zones []*Zone
}

func (e *ZoneCollection) Zones() *[]*Zone {
	return &e.zones
}

func NewZoneCollection(zones []*Zone) *ZoneCollection {
	return &ZoneCollection{zones: zones}
}

type Zone struct {
	zone *objects.Zone
}

func (z *Zone) Entities() *[]*Entity {
	list := make([]*Entity, 0)

	for _, entity := range z.zone.Entities() {
		list = append(list, NewEntity(entity))
	}

	return &list
}

func (z *Zone) Players() *[]*Player {
	list := make([]*Player, 0)

	for _, entity := range z.zone.Players() {
		list = append(list, NewPlayer(entity))
	}

	return &list
}

func (z *Zone) Name() *string {
	return &z.zone.Name
}

func NewZone(z *objects.Zone) *Zone {
	return &Zone{
		zone: z,
	}
}
