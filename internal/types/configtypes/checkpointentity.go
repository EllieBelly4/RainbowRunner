package configtypes

//go:generate go run ../../../scripts/generatelua -type=CheckpointEntityConfig -extends=EntityConfig
type CheckpointEntityConfig struct {
	*EntityConfig
}

func NewCheckpointEntityConfig() *CheckpointEntityConfig {
	return &CheckpointEntityConfig{
		EntityConfig: NewEntityConfig(EntityConfigTypeCheckpoint),
	}
}
