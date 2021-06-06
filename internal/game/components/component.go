package components

import "RainbowRunner/internal/byter"

type BaseComponent struct {
	ID uint16
}

type Component interface {
	AddUpdate(body *byter.Byter)
}
