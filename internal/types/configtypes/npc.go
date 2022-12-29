package configtypes

//go:generate go run ../../../scripts/generatelua -type=NPCConfig -extends=EntityConfig
type NPCConfig struct {
	*EntityConfig
}

func NewNPCConfigFromEntity(entity *EntityConfig) *NPCConfig {
	npcConfig := &NPCConfig{}
	npcConfig.EntityConfig = entity

	return npcConfig
}
