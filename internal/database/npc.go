package database

import (
	"RainbowRunner/internal/types/configtypes"
	"strconv"
)

type NPCConfig struct {
	Name       string
	Level      int32
	FullGCType string
}

func NewNPCConfig(config *configtypes.DRClassChildGroup) *NPCConfig {
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

	return npcConfig
}
