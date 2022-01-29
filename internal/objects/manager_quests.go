package objects

import (
	"RainbowRunner/pkg/byter"
)

type QuestManager struct {
	*GCObject
}

func (q QuestManager) Type() DRObjectType {
	return DRObjectManager
}

func NewQuestManager() *QuestManager {
	q := &QuestManager{
		GCObject: NewGCObject("QuestManager"),
	}

	q.GCType = "QuestManager"

	return q
}

func (q QuestManager) WriteInit(b *byter.Byter) {
	// QuestManager::readInit()
	b.WriteUInt32(0x01)
	b.WriteByte(0x01)
	b.WriteCString("Hello")
	b.WriteCString("HelloAgain")
	b.WriteUInt32(0x01)
	b.WriteByte(0x01)
	b.WriteCString("HelloAgainAgain")
	b.WriteCString("HelloAgainAgainAgain")
	b.WriteUInt32(0x01)
	b.WriteCString("Hi")
	b.WriteCString("HiAgain")
	b.WriteCString("HiAgainAgain")

	// QuestManager::ReadAvailableQuests()
	b.WriteByte(0x00) // Probably quest count

	// QuestManager::readInit()
	b.WriteUInt16(0x00) // Objectives count?
	b.WriteUInt16(0x00) // Some count
}

func (q QuestManager) WriteUpdate(b *byter.Byter) {
	panic("implement me")
}
