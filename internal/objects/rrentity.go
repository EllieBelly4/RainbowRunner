package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/pkg/byter"
)

type RREntityProperties struct {
	OwnerID int
	ID      uint32
	Conn    connections.Connection
	Zone    *Zone
}

type RREntity struct {
	Object  DRObject
	OwnerID int
	ID      uint32
}

func (R *RREntity) WriteFullGCObject(b *byter.Byter) {
	panic("implement me")
}

func (R *RREntity) WriteInit(b *byter.Byter) {
	panic("implement me")
}

func (R *RREntity) WriteUpdate(b *byter.Byter) {
	panic("implement me")
}

func (R *RREntity) AddChild(avatar *RREntity) {

}
