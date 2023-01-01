package objects

import (
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
)

//go:generate go run ../../scripts/generatelua -type=Hero -extends=Unit
type Hero struct {
	*Unit

	// The actual EXP value you want to add needs to be multiplied by 20
	ExpThisLevel uint32

	Strength            uint16
	Agility             uint16
	Endurance           uint16
	Intellect           uint16
	StatPointsRemaining uint16
	RespecSomething     uint16

	HeroUnk0 uint32
	HeroUnk1 uint32
}

func (h *Hero) WriteInit(body *byter.Byter) {
	h.Unit.WriteInit(body)

	body.WriteUInt32(h.ExpThisLevel)

	body.WriteUInt16(h.Strength)
	body.WriteUInt16(h.Agility)
	body.WriteUInt16(h.Endurance)
	body.WriteUInt16(h.Intellect)
	body.WriteUInt16(h.StatPointsRemaining)
	body.WriteUInt16(h.RespecSomething)

	body.WriteUInt32(h.HeroUnk0)
	body.WriteUInt32(h.HeroUnk1)
}

func (h *Hero) AddChild(child drobjecttypes.DRObject) {
	h.Unit.AddChild(child)
	child.SetParent(h)
}

func NewHero(gcType string) *Hero {
	return &Hero{Unit: NewUnit(gcType)}
}
