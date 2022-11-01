package database

import (
	"RainbowRunner/internal/types/configtypes"
	"strconv"
	"strings"
)

type BehaviourConfig struct {
	Type string
	Desc *BehaviourDescConfig
}

type BehaviourDescConfig struct {
	Type string
}

type NPCConfig struct {
	Name       string
	Level      int32
	FullGCType string
	Behaviour  *BehaviourConfig
}

//type NPCDescriptionConfig struct {}

func NewNPCConfig(config *configtypes.DRClassChildGroup, gcRoot []string) *NPCConfig {
	npcConfig := &NPCConfig{}
	entity := config.Entities[0]

	if name, ok := entity.Properties["Name"]; ok {
		npcConfig.Name = name
	}

	if level, ok := entity.Properties["Level"]; ok {
		intVal, err := strconv.ParseInt(level, 10, 32)

		if err != nil {
			panic(err)
		}

		npcConfig.Level = int32(intVal)
	}

	if behaviour, ok := entity.Children["behavior"]; ok {
		behaviourConf := handleNPCBehavior(behaviour, gcRoot)
		npcConfig.Behaviour = behaviourConf
	}

	return npcConfig
}

func handleNPCBehavior(behaviour *configtypes.DRClassChildGroup, root []string) *BehaviourConfig {
	res := &BehaviourConfig{}

	entity := behaviour.Entities[0]

	res.Type = entity.Extends

	behaviorRoot := append(root, "behavior")

	if res.Type == "" {
		res.Type = strings.Join(behaviorRoot, ".")
	} else if res.Type == "monsterbehavior2" {
		res.Type = "npc.Base.Behavior"
	}

	if entityDesc, ok := entity.Children["description"]; ok {
		descEntity := entityDesc.Entities[0]

		behaviourDescType := descEntity.Extends

		if behaviourDescType == "" {
			behaviourDescType = strings.Join(append(behaviorRoot, "description"), ".")
		} else if behaviourDescType == "monsterbehavior2desc" {
			behaviourDescType = "npc.Base.Behavior.Description"
		}

		res.Desc = &BehaviourDescConfig{
			Type: behaviourDescType,
		}
	}

	return res
}
