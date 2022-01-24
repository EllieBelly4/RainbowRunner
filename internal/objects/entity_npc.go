package objects

import (
	"RainbowRunner/internal/game/components/behavior"
	"RainbowRunner/pkg"
	"RainbowRunner/pkg/byter"
)

type NPC struct {
	*StockUnit
}

func (n *NPC) WriteInit(b *byter.Byter) {
	n.StockUnit.WriteInit(b)
}

func NewNPC(
	gcType,
	behaviourType string,
	position pkg.Vector3,
	rotation int,
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
