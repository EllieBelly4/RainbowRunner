package objects

import (
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
)

//go:generate go run ../../scripts/generatelua -type=QuestManager -extends=GCObject
type QuestManager struct {
	*GCObject
}

func (q QuestManager) Type() drobjecttypes.DRObjectType {
	return drobjecttypes.DRObjectManager
}

func NewQuestManager() *QuestManager {
	q := &QuestManager{
		GCObject: NewGCObject("QuestManager"),
	}

	q.GCType = "QuestManager"

	return q
}

func (q *QuestManager) WriteFullGCObject(byter *byter.Byter) {
	q.GCObject.WriteFullGCObject(byter)

	byter.WriteCString("SomethingUnknown")
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

	// Testing Quests:
	// Snowman1 world.solo.dungeon_snowman.quest.Q01_a1

	// QuestManager::ReadAvailableQuests()
	quests := []string{
		"world.town.quest.class.fi.Q01_a1",
		"world.solo.dungeon_snowman.quest.Q01_a1",
		"world.dungeon00.quest.Q02_a4",
	}

	// Setting this to more than 0 makes the player a questgiver with an exclamation mark and everything
	questgiverQuestCount0 := 0
	b.WriteByte(byte(questgiverQuestCount0)) // Probably quest count

	for i := 0; i < questgiverQuestCount0; i++ {
		// Cannot resolve ArchetypeRef<class Entity> - Reference with name 'something' cannot be found
		// "Entity" works
		b.WriteCString("Entity") // Unk

		questgiverQuestCount1 := 1

		b.WriteByte(byte(questgiverQuestCount1))
		writeGCType(b, quests[i])
	}

	// QuestManager::readInit()
	// Currently assuming this is active
	activeQuestCount := 1
	b.WriteUInt16(uint16(activeQuestCount))

	for i := 0; i < activeQuestCount; i++ {
		writeGCType(b, quests[i])

		b.WriteUInt32(0x00) // Unk
		b.WriteByte(0x00)   // Completed flag - 0x01 = completed

		// Quest::readObjectives
		objectiveCount := 0x01
		b.WriteByte(byte(objectiveCount)) // Unknown Count

		for j := 0; j < objectiveCount; j++ {
			someFlag := 0x00
			b.WriteByte(byte(someFlag))
			b.WriteCString("Make something great happen 0/100") // Current objective text

			if someFlag == 0x02 {
				b.WriteUInt16(0x00) // Unk
			}
		}
	}

	checkpointCount := 0x01

	b.WriteUInt16(uint16(checkpointCount)) // Checkpoints count, no idea what this actually does

	for j := 0; j < checkpointCount; j++ {
		writeGCType(b, "world.checkpoints.TownCheckpoint")
	}
}

func (q QuestManager) WriteUpdate(b *byter.Byter) {
	panic("implement me")
}
