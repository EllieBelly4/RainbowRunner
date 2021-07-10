package connections

import "RainbowRunner/pkg/byter"

type Connection interface {
	Send(b *byter.Byter) error
	GetID() int
}
