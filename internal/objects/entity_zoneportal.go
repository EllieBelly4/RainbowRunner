package objects

import (
	"RainbowRunner/pkg/byter"
	log "github.com/sirupsen/logrus"
)

//go:generate go run ../../scripts/generatelua -type=ZonePortal -extends=WorldEntity
type ZonePortal struct {
	*WorldEntity
	Unk0   string
	Unk1   string
	Width  uint16
	Height uint16
	Unk4   uint32

	Target string
}

func (z ZonePortal) Activate(player *RRPlayer, u *UnitBehavior, id byte, seqID byte) {
	z.WorldEntity.Activate(player, u, id, seqID)

	tZone := Zones.GetOrCreateZone(z.Target)

	if tZone == nil {
		log.Errorf("could not find zone %s", z.Target)
		return
	}

	player.CurrentCharacter.JoinZone(tZone)
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
	worldEntity.WorldEntityFlags = 0x04

	return &ZonePortal{
		WorldEntity: worldEntity,
		Unk0:        unk0,
		Unk1:        unk1,
	}
}
