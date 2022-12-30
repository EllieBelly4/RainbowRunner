package objects

import (
	"RainbowRunner/internal/types/configtypes"
)

//go:generate go run ../../scripts/generatelua -type=CheckpointEntity -extends=WorldEntity
type CheckpointEntity struct {
	*WorldEntity
	BaseConfig *configtypes.CheckpointEntityConfig
}

func NewCheckpointEntity(gctype string) *CheckpointEntity {
	worldEntity := NewWorldEntity(gctype)

	worldEntity.WorldEntityFlags = 0x07

	return &CheckpointEntity{
		WorldEntity: worldEntity,
	}
}

func NewCheckpointEntityFromConfig(config *configtypes.CheckpointEntityConfig) *CheckpointEntity {
	entity := NewCheckpointEntity(config.FullGCType)

	entity.BaseConfig = config

	return entity
}
