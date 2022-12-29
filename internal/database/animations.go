package database

import (
	"RainbowRunner/internal/types/configtypes"
	"strconv"
)

func loadAnimations(gcType string) map[int]configtypes.AnimationConfig {
	animationResults := make(map[int]configtypes.AnimationConfig)

	config, err := config.Get(gcType)

	if err != nil {
		return animationResults
	}

	for _, animations := range config[0].Entities[0].Children {
		for _, animation := range animations.Entities {
			anim := configtypes.AnimationConfig{}

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
