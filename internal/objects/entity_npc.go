package objects

import "RainbowRunner/pkg/byter"

type NPC struct {
	*StockUnit
}

func (n *NPC) WriteInit(b *byter.Byter) {
	n.StockUnit.WriteInit(b)
}

func NewNPC(gcType string) *NPC {
	unit := NewStockUnit(gcType)
	unit.GCType = gcType

	unit.UnitFlags = 0
	// Adding 0x01 makes it super speedy and disables mouse movement, client selected entity?
	unit.WorldEntityFlags = 0x04
	unit.WorldEntityInitFlags = 0

	return &NPC{
		StockUnit: unit,
	}
}
