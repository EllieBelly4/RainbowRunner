package database

import "RainbowRunner/internal/types/configtypes"

type NPCConfig struct {
	Name string
}

func NewNPCCOnfig(config *configtypes.DRClassChildGroup) *NPCConfig {
	return &NPCConfig{
		Name: config.Name,
	}
}
