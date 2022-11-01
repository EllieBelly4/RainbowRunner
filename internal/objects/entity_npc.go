package objects

import (
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/game/components/behavior"
	"RainbowRunner/internal/global"
	"RainbowRunner/internal/message"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
)

type NPC struct {
	*StockUnit

	Name  string
	Level int32
}

func (n *NPC) WriteInit(b *byter.Byter) {
	n.StockUnit.WriteInit(b)
}

func NewNPCSimple(gcType string) *NPC {
	return NewNPC(gcType, "", datatypes.Vector3Float32{}, 0)
}

func NewNPC(
	gcType,
	behaviourType string,
	position datatypes.Vector3Float32,
	rotation float32,
) *NPC {
	unit := NewStockUnit(gcType)
	unit.GCType = gcType

	unit.UnitFlags = 0
	// Adding 0x01 makes it super speedy and disables mouse movement, client selected entity?
	unit.WorldEntityFlags = 0x04
	unit.WorldEntityInitFlags = 0

	unit.WorldPosition = position
	unit.Rotation = rotation

	npc := &NPC{
		StockUnit: unit,
	}

	//npc.addBehaviour(behaviourType)
	//
	//skills := NewSkills("skills")
	//npc.AddChild(skills)
	//
	//manipulators := NewManipulators("manipulators")
	//npc.AddChild(manipulators)
	//
	//modifiers := NewModifiers("modifiers")
	//npc.AddChild(modifiers)

	return npc
}

func (n *NPC) addBehaviour(behaviourType string) {
	unitBehavior := NewMonsterBehavior2(behaviourType)

	unitBehavior.Action1 = &behavior.MoveTo{
		PosX: uint32(n.WorldPosition.X),
		PosY: uint32(n.WorldPosition.Y),
	}

	//unitBehavior.Action2 = &behavior.WarpTo{
	//	PosX: uint32(npc.WorldPosition.X),
	//	PosY: uint32(npc.WorldPosition.Y),
	//}

	n.AddChild(unitBehavior)
}

func CreateNPC(player *RRPlayer, zone *Zone, transform datatypes.Transform, npcType, behaviourType string) {
	npc := NewNPC(npcType, behaviourType, transform.Position, transform.Rotation)

	npc.WorldPosition = transform.Position
	npc.Rotation = transform.Rotation

	unitBehavior := NewMonsterBehavior2(behaviourType)

	unitBehavior.Action1 = &behavior.MoveTo{
		PosX: uint32(npc.WorldPosition.X),
		PosY: uint32(npc.WorldPosition.Y),
	}

	//unitBehavior.Action2 = &behavior.WarpTo{
	//	PosX: uint32(npc.WorldPosition.X),
	//	PosY: uint32(npc.WorldPosition.Y),
	//}

	npc.AddChild(unitBehavior)

	skills := NewSkills("skills")
	npc.AddChild(skills)

	manipulators := NewManipulators("manipulators")
	npc.AddChild(manipulators)

	modifiers := NewModifiers("modifiers")
	npc.AddChild(modifiers)

	zone.Spawn(npc)

	//clientEntityWriter := NewClientEntityWriterWithByter()
	//clientEntityWriter.BeginStream()

	global.JobQueue.Enqueue(func() {
		CEWriter := NewClientEntityWriterWithByter()
		CEWriter.Create(npc)
		CEWriter.CreateComponentAndInit(skills, npc)
		CEWriter.CreateComponentAndInit(manipulators, npc)
		CEWriter.CreateComponentAndInit(modifiers, npc)
		// Adding unit behavior makes the NPC move in a random direction, missing something here
		//player.ClientEntityWriter.CreateComponent(unitBehavior, npc)
		CEWriter.Init(npc)
		player.MessageQueue.Enqueue(message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeCreateNPC)
	})

	//player.ClientEntityWriter.EndStream()

	//helpers.WriteCompressedASimple(player.Conn, player.ClientEntityWriter.Body)

	//unitBehavior.Warp(106342, -46263, 12778)
	//unitBehavior.Warp(0, 0, 0)
	//unitBehavior.SendPosition()
}

func NewNPCFromConfig(config *database.NPCConfig) *NPC {
	npc := NewNPCSimple(config.FullGCType)

	npc.Name = config.Name
	npc.Level = config.Level

	if config.Behaviour != nil {
		//if config.Behaviour.Type == "monsterbehavior2" {
		npc.AddChild(NewMonsterBehavior2(config.Behaviour.Type))
		//} else {
		//	npc.AddChild(NewUnitBehavior(config.Behaviour.Type))
		//}
	}

	skills := NewSkills("skills")
	npc.AddChild(skills)

	manipulators := NewManipulators("manipulators")
	npc.AddChild(manipulators)

	modifiers := NewModifiers("modifiers")
	npc.AddChild(modifiers)

	return npc
}
