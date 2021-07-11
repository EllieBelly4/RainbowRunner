package objects

import (
	"RainbowRunner/pkg/byter"
)

type MonsterBehavior2 struct {
	*UnitBehavior
}

func (n *MonsterBehavior2) WriteInit(b *byter.Byter) {
	n.UnitBehavior.WriteInit(b)

	//StateMachine::ReadMessage
	// Flags
	stateMachineFlags := byte(0x00)
	b.WriteByte(stateMachineFlags)

	if stateMachineFlags&0x02 > 0 {
		b.WriteUInt16(0x00)
	}

	if stateMachineFlags&0x04 > 0 {
		b.WriteUInt16(0x00)
	}

	if stateMachineFlags&0x08 > 0 {
		b.WriteUInt16(0x00)
	}

	if stateMachineFlags&0x10 > 0 {
		b.WriteUInt16(0x00)
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

	b.WriteUInt32(0x00)
	b.WriteUInt32(0x00)

	if monsterBehaviorFlags&0x04 > 0 {
		b.WriteUInt16(0x00)
	}

	if monsterBehaviorFlags&0x08 > 0 {
		b.WriteUInt16(0x00)
	}

	if monsterBehaviorFlags&0x10 > 0 {
		b.WriteUInt16(0x00)
	}
}

func NewMonsterBehavior2(gctype string) *MonsterBehavior2 {
	return &MonsterBehavior2{
		UnitBehavior: NewUnitBehavior(gctype),
	}
}
