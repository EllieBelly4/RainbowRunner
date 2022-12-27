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

//go:generate go run ../../scripts/generatelua -type NPCConfig
type NPCConfig struct {
	Name            string
	Level           int32
	FullGCType      string
	Behaviour       *BehaviourConfig
	Animations      map[int]AnimationConfig
	Speed           int
	CollisionRadius int
	TurnRate        int
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

		if speed, ok := description.Properties["Speed"]; ok {
			intVal, err := strconv.ParseInt(speed, 10, 32)

			if err != nil {
				panic(err)
			}

			npcConfig.Speed = int(intVal)
		}

		if collisionRadius, ok := description.Properties["CollisionRadius"]; ok {
			intVal, err := strconv.ParseInt(collisionRadius, 10, 32)

			if err != nil {
				panic(err)
			}

			npcConfig.CollisionRadius = int(intVal)
		}

		if turnRate, ok := description.Properties["TurnRate"]; ok {
			intVal, err := strconv.ParseInt(turnRate, 10, 32)

			if err != nil {
				panic(err)
			}

			npcConfig.TurnRate = int(intVal)
		}
	}

	return npcConfig
}

func loadAnimations(gcType string) map[int]AnimationConfig {
	animationResults := make(map[int]AnimationConfig)

	config, err := config.Get(gcType)

	if err != nil {
		return animationResults
	}

	for _, animations := range config[0].Entities[0].Children {
		for _, animation := range animations.Entities {
			anim := AnimationConfig{}

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
