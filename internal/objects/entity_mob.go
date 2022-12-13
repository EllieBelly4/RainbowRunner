package objects

import (
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/game/components/behavior"
	"RainbowRunner/internal/global"
	"RainbowRunner/internal/message"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
)

type MOB struct {
	*StockUnit

	Name  string
	Level int32
}

func (n *MOB) WriteInit(b *byter.Byter) {
	n.StockUnit.WriteInit(b)
}

func NewMOBSimple(gcType string) *MOB {
	return NewMOB(gcType, "", datatypes.Vector3Float32{}, 0)
}

func NewMOB(
	gcType,
	behaviourType string,
	position datatypes.Vector3Float32,
	rotation float32,
) *MOB {
	unit := NewStockUnit(gcType)
	unit.GCType = gcType

	unit.UnitFlags = 0
	// Adding 0x01 makes it super speedy and disables mouse movement, client selected entity?
	unit.WorldEntityFlags = 0x04
	unit.WorldEntityInitFlags = 0

	unit.WorldPosition = position
	unit.Rotation = rotation

	mob := &MOB{
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

	return mob
}

func (n *MOB) addBehaviour(behaviourType string) {
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

func CreateMOB(player *RRPlayer, zone *Zone, transform datatypes.Transform, mobType, behaviourType string) {
	mob := NewMOB(mobType, behaviourType, transform.Position, transform.Rotation)

	mob.WorldPosition = transform.Position
	mob.Rotation = transform.Rotation

	unitBehavior := NewMonsterBehavior2(behaviourType)

	unitBehavior.Action1 = &behavior.MoveTo{
		PosX: uint32(mob.WorldPosition.X),
		PosY: uint32(mob.WorldPosition.Y),
	}

	//unitBehavior.Action2 = &behavior.WarpTo{
	//	PosX: uint32(npc.WorldPosition.X),
	//	PosY: uint32(npc.WorldPosition.Y),
	//}

	mob.AddChild(unitBehavior)

	skills := NewSkills("skills")
	mob.AddChild(skills)

	manipulators := NewManipulators("manipulators")
	mob.AddChild(manipulators)

	modifiers := NewModifiers("modifiers")
	mob.AddChild(modifiers)

	zone.Spawn(mob)

	//clientEntityWriter := NewClientEntityWriterWithByter()
	//clientEntityWriter.BeginStream()

	global.JobQueue.Enqueue(func() {
		CEWriter := NewClientEntityWriterWithByter()
		CEWriter.Create(mob)
		CEWriter.CreateComponentAndInit(skills, mob)
		CEWriter.CreateComponentAndInit(manipulators, mob)
		CEWriter.CreateComponentAndInit(modifiers, mob)
		// Adding unit behavior makes the NPC move in a random direction, missing something here
		//player.ClientEntityWriter.CreateComponent(unitBehavior, npc)
		CEWriter.Init(mob)
		player.MessageQueue.Enqueue(message.QueueTypeClientEntity, CEWriter.Body, message.OpTypeCreateMOB)
	})

	//player.ClientEntityWriter.EndStream()

	//helpers.WriteCompressedASimple(player.Conn, player.ClientEntityWriter.Body)

	//unitBehavior.Warp(106342, -46263, 12778)
	//unitBehavior.Warp(0, 0, 0)
	//unitBehavior.SendPosition()
}

func NewMOBFromConfig(config *database.MOBConfig) *MOB {
	mob := NewMOBSimple(config.FullGCType)

	mob.Name = config.Name
	mob.Level = config.Level

	if config.Behaviour != nil {
		//if config.Behaviour.Type == "monsterbehavior2" {
		mob.AddChild(NewMonsterBehavior2(config.Behaviour.Type))
		//} else {
		//	npc.AddChild(NewUnitBehavior(config.Behaviour.Type))
		//}
	}

	skills := NewSkills("skills")
	mob.AddChild(skills)

	manipulators := NewManipulators("manipulators")
	mob.AddChild(manipulators)

	modifiers := NewModifiers("modifiers")
	mob.AddChild(modifiers)

	return mob
}
