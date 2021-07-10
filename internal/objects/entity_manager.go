package objects

var currentID = uint16(1)

type EntityManager struct {
	Entities map[uint16]DRObject
}

func (m *EntityManager) RegisterAll(ownerID int, object DRObject) {
	props := object.RREntityProperties()

	for _, c := range object.Children() {
		m.RegisterAll(ownerID, c)
	}

	if props.ID == 0 {
		props.OwnerID = ownerID
		props.ID = NewID()

		m.Entities[props.ID] = object
	}
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
