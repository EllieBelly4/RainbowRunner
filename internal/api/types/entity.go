package types

import (
	"RainbowRunner/internal/objects"
	"RainbowRunner/internal/types/drobjecttypes"
)

type EntityCollection struct {
	entities []*Entity
}

func (e *EntityCollection) Entities() *[]*Entity {
	return &e.entities
}

func NewEntityCollection(entities []*Entity) *EntityCollection {
	return &EntityCollection{entities: entities}
}

type Entity struct {
	obj drobjecttypes.DRObject
}

func (e *Entity) Zone() *Zone {
	if e.obj.(objects.IRREntityPropertiesHaver).GetRREntityProperties().Zone == nil {
		return nil
	}

	return NewZone(e.obj.(objects.IRREntityPropertiesHaver).GetRREntityProperties().Zone)
}

func (e *Entity) TypeName() *string {
	return &e.obj.(objects.IGCObject).GetGCObject().GCType
}

func (e *Entity) Id() *int32 {
	id := int32(e.obj.(objects.IRREntityPropertiesHaver).GetRREntityProperties().ID)
	return &id
}

func (e *Entity) OwnerId() *int32 {
	id := int32(e.obj.(objects.IRREntityPropertiesHaver).GetRREntityProperties().OwnerID)
	return &id
}

func (e *Entity) Children() *[]*Entity {
	//list := make([]*EntityChildResolver, 0)
	//
	//for _, child := range e.obj.Children() {
	//	list = append(list, &EntityChildResolver{result: NewEntity(child)})
	//}

	list := make([]*Entity, 0)

	for _, child := range e.obj.Children() {
		list = append(list, NewEntity(child))
	}

	return &list
}

func NewEntity(e drobjecttypes.DRObject) *Entity {
	return &Entity{
		obj: e,
	}
}
