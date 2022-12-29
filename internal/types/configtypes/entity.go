package configtypes

import (
	"RainbowRunner/pkg/datatypes"
	"RainbowRunner/pkg/datatypes/drfloat"
)

//go:generate go run ../../../scripts/generatelua -type=EntityConfig
type EntityConfig struct {
	Name             string
	HitPoints        drfloat.DRFloat
	ManaPoints       drfloat.DRFloat
	Position         datatypes.Vector3Float32
	Heading          int
	Width            int
	Zone             string
	EncounterTable   string
	Height           int
	SpawnPoint       string
	SizeX            int
	SizeY            int
	SizeZ            int
	CanBeActivated   bool
	RespawnWhenClear bool
	Blocking         bool
	TableSelector    int
	Color            uint `parse:"hex"`
	ZoneStart        bool
	Level            int
	AutoRespawn      bool
	WorldEntityTable string
	RespawnRate      int
}

func NewEntityConfig() *EntityConfig {
	return &EntityConfig{
		CanBeActivated: true,
	}
}
