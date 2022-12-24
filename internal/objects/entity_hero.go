package objects

import "RainbowRunner/internal/types/drobjecttypes"

//go:generate go run ../../scripts/generatelua -type=Hero -extends=Unit
type Hero struct {
	*Unit
}

func (h *Hero) AddChild(child drobjecttypes.DRObject) {
	h.Unit.AddChild(child)
	child.SetParent(h)
}

func NewHero(gcType string) *Hero {
	return &Hero{Unit: NewUnit(gcType)}
}
