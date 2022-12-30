package database

import (
	"RainbowRunner/internal/types/configtypes"
)

func loadAnimations(gcType string) map[int]configtypes.AnimationConfig {
	animationResults := make(map[int]configtypes.AnimationConfig)

	config, err := config.Get(gcType)

	if err != nil {
		return animationResults
	}

	for _, animations := range config[0].Entities[0].Children {
		for _, animation := range animations.Entities {
			anim := &configtypes.AnimationConfig{}

			configtypes.SetPropertiesOnStruct(anim, animation.Properties)

			animationResults[anim.ID] = *anim
		}
	}

	return animationResults
}
