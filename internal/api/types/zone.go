package types

import "RainbowRunner/internal/objects"

type Zone struct {
	name string
}

func (z *Zone) Name() *string {
	return &z.name
}

func NewZone(z *objects.Zone) *Zone {
	return &Zone{
		name: z.Name,
	}
}
