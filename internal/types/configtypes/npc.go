package configtypes

//go:generate go run ../../../scripts/generatelua -type=NPCConfig -extends=EntityConfig
type NPCConfig struct {
	*EntityConfig
}

func NewNPCConfig() *NPCConfig {
	return &NPCConfig{EntityConfig: NewEntityConfig(EntityConfigTypeNPC)}
}

func NewNPCConfigFromEntity(entity *EntityConfig) *NPCConfig {
	npcConfig := &NPCConfig{}
	npcConfig.EntityConfig = entity

	return npcConfig
}
