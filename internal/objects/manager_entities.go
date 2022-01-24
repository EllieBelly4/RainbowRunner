package objects

import (
	"RainbowRunner/internal/connections"
	"strings"
	"sync"
)

var Entities *EntityManager

// Reserving 10 IDs for player characters
var currentID = uint16(10)

type EntityManager struct {
	sync.RWMutex
	// Readonly
	Entities map[uint16]DRObject
}

func (m *EntityManager) RegisterAll(owner connections.Connection, objects ...DRObject) {
	for _, object := range objects {
		if strings.ToLower(object.GetGCObject().GCType) == "player" {
			panic("do not try to register player")
		}

		props := object.RREntityProperties()

		m.RegisterAll(owner, object.Children()...)

		if props.ID == 0 {
			if owner != nil {
				props.Conn = owner
				props.OwnerID = owner.GetID()
			}

			props.ID = NewID()

			m.Lock()
			m.Entities[props.ID] = object
			m.Unlock()
		}
	}
}

func (m *EntityManager) Tick() {
	m.RLock()
	defer m.RUnlock()

	for _, entity := range m.Entities {
		entity.Tick()
	}
}

func (m *EntityManager) FindByID(id uint16) DRObject {
	m.RLock()
	defer m.RUnlock()
	for _, entity := range m.Entities {
		if entity.RREntityProperties().ID == id {
			return entity
		}
	}
	return nil
}

func (m *EntityManager) RemoveOwnedBy(id int) {
	toDelete := make([]uint16, 0, 1024)

	m.RLock()
	for index, entity := range m.Entities {
		if entity == nil || entity.RREntityProperties().OwnerID == id {
			toDelete = append(toDelete, index)
		}
	}
	m.RUnlock()

	m.Lock()
	for _, index := range toDelete {
		m.Entities[index].RREntityProperties().Zone = nil
		delete(m.Entities, index)
	}
	m.Unlock()
}

func (m *EntityManager) GetEntities() []DRObject {
	m.RLock()
	defer m.RUnlock()

	list := make([]DRObject, 0)

	for _, entity := range m.Entities {
		list = append(list, entity)
	}

	return list
}

func NewID() (ID uint16) {
	ID = currentID
	currentID++
	return ID
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		Entities: make(map[uint16]DRObject, 1024*10),
	}
}
