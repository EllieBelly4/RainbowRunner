package objects

import (
	"RainbowRunner/internal/types/configtypes"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	"strings"
)

//go:generate go run ../../scripts/generateLua/ -type=NPC -extends=Unit
type NPC struct {
	*StockUnit

	Level int32
}

func (n *NPC) WriteInit(b *byter.Byter) {
	n.StockUnit.WriteInit(b)
}

func NewNPCSimple(gcType string) *NPC {
	return NewNPC(gcType, datatypes.Vector3Float32{}, 0)
}

func NewNPC(gcType string, position datatypes.Vector3Float32, rotation float32) *NPC {
	unit := NewStockUnit(gcType)
	unit.GCType = gcType

	unit.WorldEntity.CanBeActivated = true

	unit.UnitFlags = 0
	// Adding 0x01 makes it super speedy and disables mouse movement, client selected entity?
	unit.WorldEntityFlags = 0x04
	unit.WorldEntityInitFlags = 0

	unit.WorldPosition = position
	unit.Heading = rotation

	npc := &NPC{
		StockUnit: unit,
	}

	return npc
}

func NewNPCFromConfig(config *configtypes.NPCConfig) *NPC {
	npc := NewNPC(config.FullGCType, config.Position, float32(config.Heading))

	npc.Name = config.Name
	npc.Level = int32(config.Level)
	npc.CollisionRadius = config.Desc.CollisionRadius

	npc.WorldEntity.CanBeActivated = config.CanBeActivated

	behaviorType := "npc.Base.Behavior"

	if strings.ToLower(config.Behaviour.Type) != "monsterbehavior2" {
		behaviorType = config.Behaviour.Type
	}

	behavior2 := NewMonsterBehavior2(behaviorType)

	behavior2.Speed = int(config.Desc.Speed)
	behavior2.TurnRate = config.Desc.TurnRate

	npc.AddChild(behavior2)

	if config.Animations != nil {
		animationsList := NewAnimationsList()

		for _, animationConf := range config.Animations {
			animationsList.Animations = append(animationsList.Animations, NewAnimationFromConfig(&animationConf))
		}

		npc.WorldEntity.AnimationsList = animationsList
	}

	if config.Merchant != nil {
		npc.AddChild(NewMerchantFromConfig(config.Merchant))
	}

	skills := NewSkills("skills")
	npc.AddChild(skills)

	manipulators := NewManipulators("manipulators")
	npc.AddChild(manipulators)

	modifiers := NewModifiers("modifiers")
	npc.AddChild(modifiers)

	return npc
}
