package objects

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/pkg/byter"
)

//go:generate go run ../../scripts/generatelua -type=RREntityProperties
type RREntityProperties struct {
	OwnerID uint16
	ID      uint32
	Conn    connections.Connection
	Zone    *Zone
}

func (p *RREntityProperties) SetOwner(id uint16) {
	p.Conn = Players.GetPlayer(id).Conn
	p.OwnerID = id
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

func NewRREntityProperties() *RREntityProperties {
	return &RREntityProperties{}
}
