package components

import (
	byter "RainbowRunner/pkg/byter"
)

type BaseComponent struct {
	ID uint16
}

type Component interface {
	AddUpdate(body *byter.Byter)
}
