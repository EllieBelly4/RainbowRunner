package objects

import "RainbowRunner/pkg/byter"

type Container struct {
	*Component
}

func (c *Container) WriteInit(b *byter.Byter) {
	// Container::readInit
	b.WriteUInt32(0xAAAABBBB)
	b.WriteUInt32(0xCCCCDDDD)

	// GCObject::readChildData<Inventory>
	for _, child := range c.Children() {
		if inventory, ok := child.(IInventory); ok {
			inventory.GetInventory().WriteInitData(b)
		}
	}

	b.WriteByte(0x00)

	// Potentially if above is something
	//b.WriteUInt16(0x00)
}

func NewContainer(gcType string, gcNativeType string) *Container {
	return &Container{
		Component: NewComponent(gcType, gcNativeType),
	}
}
