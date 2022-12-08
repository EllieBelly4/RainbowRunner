package messages

import "RainbowRunner/pkg/byter"

type DRMessage interface {
	Write(b *byter.Byter)
}
