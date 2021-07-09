package objects

import (
	byter "RainbowRunner/pkg/byter"
)

type DRObject interface {
	Serialise(b *byter.Byter)
	WriteInit(b *byter.Byter)
	WriteUpdate(b *byter.Byter)
}
