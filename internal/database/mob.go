package database

import (
	"RainbowRunner/internal/types/configtypes"
	"strconv"
	"strings"
)

type BehaviourMOBConfig struct {
	Type string
	Desc *BehaviourDescMOBConfig
}

type BehaviourDescMOBConfig struct {
	Type string
}

type MOBConfig struct {
	Name       string
	Level      int32
	FullGCType string
	Behaviour  *BehaviourMOBConfig
}

//type NPCDescriptionConfig struct {}

func NewMOBConfig(config *configtypes.DRClassChildGroup, gcRoot []string) *MOBConfig {
	mobConfig := &MOBConfig{}
	entity := config.Entities[0]

	if name, ok := entity.Properties["Name"]; ok {
		mobConfig.Name = name
	}

	if level, ok := entity.Properties["Level"]; ok {
		intVal, err := strconv.ParseInt(level, 10, 32)

		if err != nil {
			panic(err)
		}

		mobConfig.Level = int32(intVal)
	}

	if behaviour, ok := entity.Children["behavior"]; ok {
		behaviourConf := handleMOBBehavior(behaviour, gcRoot)
		mobConfig.Behaviour = behaviourConf
	}

	return mobConfig
}

func handleMOBBehavior(behaviour *configtypes.DRClassChildGroup, root []string) *BehaviourMOBConfig {
	res := &BehaviourMOBConfig{}

	entity := behaviour.Entities[0]

	res.Type = entity.Extends

	behaviorRoot := append(root, "behavior")

	if res.Type == "" {
		res.Type = strings.Join(behaviorRoot, ".")
	} else if res.Type == "monsterbehavior2" {
		res.Type = "mob.Base.Behavior"
	}

	if entityDesc, ok := entity.Children["description"]; ok {
		descEntity := entityDesc.Entities[0]

		behaviourDescType := descEntity.Extends

		if behaviourDescType == "" {
			behaviourDescType = strings.Join(append(behaviorRoot, "description"), ".")
		} else if behaviourDescType == "monsterbehavior2desc" {
			behaviourDescType = "mob.Base.Behavior.Description"
		}

		res.Desc = &BehaviourDescMOBConfig{
			Type: behaviourDescType,
		}
	}

	return res
}
