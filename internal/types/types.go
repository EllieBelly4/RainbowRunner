package types

//go:generate stringer -type=EquipmentSlot
type EquipmentSlot uint32

const (
	EquipmentSlotNone EquipmentSlot = iota
	EquipmentSlotAmulet
	EquipmentSlotHand
	EquipmentSlotLRing
	EquipmentSlotRRing
	EquipmentSlotHead
	EquipmentSlotTorso
	EquipmentSlotFoot
	EquipmentSlotShoulder
	EquipmentSlotNone2
	EquipmentSlotWeapon
	EquipmentSlotOffhand
)
