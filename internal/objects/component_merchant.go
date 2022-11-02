package objects

import (
	"RainbowRunner/pkg/byter"
)

type Merchant struct {
	*Container
}

func (m *Merchant) WriteInit(b *byter.Byter) {
	m.Container.WriteInit(b)

	someFlag := 0x00

	b.WriteByte(byte(someFlag))

	if someFlag == 0x02 {
		b.WriteUInt16(0x00) // Unk
	}

	b.WriteUInt16(0xFF)
}

func (m *Merchant) GetInventoryByID(index byte) *Inventory {
	for _, child := range m.GCChildren {
		if inventory, ok := child.(*Inventory); ok {
			if inventory.InventoryID == index {
				return inventory
			}
		}
	}

	return nil
}

func NewMerchant(gcType string) *Merchant {
	container := NewContainer(gcType, "Merchant")

	return &Merchant{
		Container: container,
	}
}
