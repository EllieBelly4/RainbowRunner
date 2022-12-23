package objects

import (
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/connections"
	"fmt"
	"strings"
	"sync"
)

// TODO consider removing/refactoring this entire thing as entities probably need to be separated by zones
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
		if strings.ToLower(object.(IGCObject).GetGCObject().GCType) == "player" {
			panic("do not try to register player")
		}

		props := object.(IRREntityProperties).GetRREntityProperties()

		m.RegisterAll(owner, object.Children()...)

		if props.ID == 0 {
			if owner != nil {
				props.Conn = owner
				props.OwnerID = uint16(owner.GetID())
			}

			props.ID = uint32(NewID())

			if config.Config.Logging.LogIDs {
				fmt.Printf("%d - %s(%s)\n", props.ID, object.(IGCObject).GetGCObject().GCType, object.(IGCObject).GetGCObject().GCLabel)
			}

			m.Lock()
			m.Entities[uint16(props.ID)] = object
			m.Unlock()
		}
	}
}

func (m *EntityManager) FindByID(id uint16) DRObject {
	m.RLock()
	defer m.RUnlock()
	for _, entity := range m.Entities {
		if entity.(IRREntityProperties).GetRREntityProperties().ID == uint32(id) {
			return entity
		}
	}
	return nil
}

func (m *EntityManager) RemoveOwnedBy(id int) {
	toDelete := make([]uint16, 0, 1024)

	m.RLock()
	for index, entity := range m.Entities {
		if entity == nil || entity.(IRREntityProperties).GetRREntityProperties().OwnerID == uint16(id) {
			toDelete = append(toDelete, index)
		}
	}
	m.RUnlock()

	m.Lock()
	for _, index := range toDelete {
		m.Entities[index].(IRREntityProperties).GetRREntityProperties().Zone = nil
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
