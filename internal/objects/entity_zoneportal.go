package objects

import "RainbowRunner/pkg/byter"

type ZonePortal struct {
	*WorldEntity
	Unk0   string
	Unk1   string
	Width  uint16
	Height uint16
	Unk4   uint32
}

func (z ZonePortal) Activate(player *RRPlayer, u *UnitBehavior, id byte) {
	z.WorldEntity.Activate(player, u, id)
}

func (z ZonePortal) WriteInit(b *byter.Byter) {
	z.WorldEntity.WriteInit(b)

	b.WriteCString(z.Unk0)
	b.WriteCString(z.Unk1)

	b.WriteUInt16(z.Width)
	b.WriteUInt16(z.Height)
	b.WriteUInt32(z.Unk4)
}

func NewZonePortal(unk0, unk1 string) *ZonePortal {
	worldEntity := NewWorldEntity("misc.ZonePortal")
	worldEntity.GCNativeType = "zoneportal"

	return &ZonePortal{
		WorldEntity: worldEntity,
		Unk0:        unk0,
		Unk1:        unk1,
	}
}
