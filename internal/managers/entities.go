package managers

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/objects"
)

var Entities = NewEntityManager()

var currentID = uint16(1)

type EntityManager struct {
	Entities map[uint16]objects.DRObject
}

func (m *EntityManager) RegisterAll(owner connections.Connection, object objects.DRObject) {
	props := object.RREntityProperties()

	for _, c := range object.Children() {
		m.RegisterAll(owner, c)
	}

	if props.ID == 0 {
		props.Conn = owner
		props.OwnerID = owner.GetID()
		props.ID = NewID()

		m.Entities[props.ID] = object
	}
}

func (m *EntityManager) Tick() {
	for _, entity := range m.Entities {
		entity.Tick()
	}
}

func NewID() (ID uint16) {
	ID = currentID
	currentID++
	return ID
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		Entities: make(map[uint16]objects.DRObject, 1024*10),
	}
}
