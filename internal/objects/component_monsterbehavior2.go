package objects

import (
	"RainbowRunner/pkg/byter"
)

//go:generate go run ../../scripts/generatelua -type=MonsterBehavior2 -extends=UnitBehavior
type MonsterBehavior2 struct {
	*UnitBehavior

	MonsterBehaviorStateMachineFlags byte
	MonsterBehavior2Unk0             uint16
	MonsterBehavior2Unk1             uint16
	MonsterBehavior2Unk2             uint16
	MonsterBehavior2Unk3             uint16

	MonsterBehaviorFlags byte

	MonsterBehavior2Unk4 uint32
	MonsterBehavior2Unk5 uint32

	MonsterBehavior2Unk6 uint16
	MonsterBehavior2Unk7 uint16
	MonsterBehavior2Unk8 uint16
}

func (n *MonsterBehavior2) WriteInit(b *byter.Byter) {
	n.UnitBehavior.WriteInit(b)

	//StateMachine::ReadMessage
	// Flags
	stateMachineFlags := byte(0x00)
	b.WriteByte(n.MonsterBehaviorStateMachineFlags)

	if stateMachineFlags&0x02 > 0 {
		b.WriteUInt16(n.MonsterBehavior2Unk0)
	}

	if stateMachineFlags&0x04 > 0 {
		b.WriteUInt16(n.MonsterBehavior2Unk1)
	}

	if stateMachineFlags&0x08 > 0 {
		b.WriteUInt16(n.MonsterBehavior2Unk2)
	}

	if stateMachineFlags&0x10 > 0 {
		b.WriteUInt16(n.MonsterBehavior2Unk3)
	}

	if stateMachineFlags&0x20 > 0 {
		// Unk
	}

	// MonsterBehavior2::readInit
	// Flags
	// Part of this flag seems to be stored after init at the end `0x0051B949`
	// bl is flag
	//.text:0051B949 010 shr     bl, 5
	//.text:0051B94C 010 xor     bl, [esi+194h]
	//.text:0051B952 010 pop     edi
	//.text:0051B953 00C and     bl, 3
	//.text:0051B956 00C xor     [esi+194h], bl
	monsterBehaviorFlags := byte(0x00)
	b.WriteByte(monsterBehaviorFlags)

	b.WriteUInt32(n.MonsterBehavior2Unk4)
	b.WriteUInt32(n.MonsterBehavior2Unk5)

	if monsterBehaviorFlags&0x04 > 0 {
		b.WriteUInt16(n.MonsterBehavior2Unk6)
	}

	if monsterBehaviorFlags&0x08 > 0 {
		b.WriteUInt16(n.MonsterBehavior2Unk7)
	}

	if monsterBehaviorFlags&0x10 > 0 {
		b.WriteUInt16(n.MonsterBehavior2Unk8)
	}
}

func NewMonsterBehavior2(gctype string) *MonsterBehavior2 {
	unitBehavior := NewUnitBehavior(gctype)

	unitBehavior.UnitMoverFlags = 0x01 | 0x04 | 0x80
	//unitBehavior.UnitMoverFlags = 0x00

	return &MonsterBehavior2{
		UnitBehavior: unitBehavior,
	}
}
