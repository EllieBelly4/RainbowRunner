package types

import "RainbowRunner/internal/objects"

type PlayerCollection struct {
	players []*Player
}

func (e *PlayerCollection) Players() *[]*Player {
	return &e.players
}

func NewPlayerCollection(players []*Player) *PlayerCollection {
	return &PlayerCollection{players: players}
}

type Player struct {
	obj *objects.RRPlayer
}

func (e *Player) Zone() *Zone {
	if e.obj.Zone == nil {
		return nil
	}

	return NewZone(e.obj.Zone)
}

func (e *Player) Name() *string {
	return &e.obj.CurrentCharacter.Name
}

func (e *Player) Id() *int32 {
	id := int32(e.obj.Conn.GetID())
	return &id
}

func (e *Player) CurrentCharacter() *Entity {
	if e.obj.CurrentCharacter == nil {
		return nil
	}

	return NewEntity(e.obj.CurrentCharacter)
}

func NewPlayer(p *objects.RRPlayer) *Player {
	return &Player{
		obj: p,
	}
}
