package configtypes

import (
	"RainbowRunner/internal/types/drconfigtypes"
)

type BehaviourConfig struct {
	Type string
	Desc *BehaviourDescConfig
}

func (c *BehaviourConfig) Init(entity *drconfigtypes.DRClass, rootGCType string) {
	if desc, ok := entity.Children["description"]; ok {
		descEntity := desc.Entities[0]

		descGCType := "npc.Base.Behavior.Description"

		if descEntity.Extends == "" {
			descGCType = rootGCType + ".behavior.description"
		}

		c.Desc = &BehaviourDescConfig{
			Type: descGCType,
		}
	}
}

type BehaviourDescConfig struct {
	Type string
}
