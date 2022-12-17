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

//go:generate go run ../../scripts/generatelua -type Animation
type Animation struct {
	ID int
	// This is used for remapping IDs to another animation
	// e.g. I think melee attacks by default have 3 animations but not all units have 3 attack animations,
	// so they map all 3 to the first animationID
	AnimationID int
	NumFrames   int

	// Unk
	TriggerTime int

	// Unk if we can actually use this
	SoundTriggerTime int

	// TODO extract this data from the config, animations do not have names but they usually have comments explaining what they are
	//Comment string
}

type NPCConfig struct {
	Name       string
	Level      int32
	FullGCType string
	Behaviour  *BehaviourConfig
	Animations map[int]Animation
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

	if descriptions, ok := entity.Children["description"]; ok {
		description := descriptions.Entities[0]
		if animationsGCType, ok := description.Properties["Animations"]; ok {
			npcConfig.Animations = loadAnimations(animationsGCType)
		}
	}

	return npcConfig
}

func loadAnimations(gcType string) map[int]Animation {
	animationResults := make(map[int]Animation)

	config, err := config.Get(gcType)

	if err != nil {
		return animationResults
	}

	for _, animations := range config[0].Entities[0].Children {
		animation := animations.Entities[0]
		anim := Animation{}

		if id, ok := animation.Properties["ID"]; ok {
			intVal, err := strconv.ParseInt(id, 10, 32)

			if err != nil {
				panic(err)
			}

			anim.ID = int(intVal)
		}

		if animID, ok := animation.Properties["AnimationID"]; ok {
			intVal, err := strconv.ParseInt(animID, 10, 32)

			if err != nil {
				panic(err)
			}

			anim.AnimationID = int(intVal)
		}

		if numFrames, ok := animation.Properties["NumFrames"]; ok {
			intVal, err := strconv.ParseInt(numFrames, 10, 32)

			if err != nil {
				panic(err)
			}

			anim.NumFrames = int(intVal)
		}

		if triggerTime, ok := animation.Properties["TriggerTime"]; ok {
			intVal, err := strconv.ParseInt(triggerTime, 10, 32)

			if err != nil {
				panic(err)
			}

			anim.TriggerTime = int(intVal)
		}

		if soundTriggerTime, ok := animation.Properties["SoundTriggerTime"]; ok {
			intVal, err := strconv.ParseInt(soundTriggerTime, 10, 32)

			if err != nil {
				panic(err)
			}

			anim.SoundTriggerTime = int(intVal)
		}

		animationResults[anim.ID] = anim
	}

	return animationResults
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
