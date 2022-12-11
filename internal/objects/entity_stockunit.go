package objects

import "RainbowRunner/pkg/byter"

//go:generate go run ../../scripts/generatelua -type=StockUnit -extends=Unit
type StockUnit struct {
	*Unit
}

func (n *StockUnit) WriteInit(b *byter.Byter) {
	n.Unit.WriteInit(b)

	// StockUnit::readInit
	b.WriteByte(0x00)
	b.WriteUInt16(0x00)
	b.WriteUInt16(0x00)
	b.WriteByte(0x00)

	b.WriteUInt16(0x00)
	b.WriteUInt32(0x00)
	b.WriteByte(0x00)

	b.WriteInt32(0x00)
	b.WriteInt32(0x00)
	b.WriteInt32(0x00)
}

func NewStockUnit(gcType string) *StockUnit {
	unit := NewUnit(gcType)
	unit.GCType = gcType

	return &StockUnit{
		Unit: unit,
	}
}
