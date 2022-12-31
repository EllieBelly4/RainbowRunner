package objects

import "RainbowRunner/internal/types/configtypes"

//go:generate go run ../../scripts/generatelua -type=Waypoint -extends=WorldEntity
type Waypoint struct {
	*WorldEntity
	BaseConfig *configtypes.WaypointConfig
}

func NewWaypoint(gctype string) *Waypoint {
	worldEntity := NewWorldEntity(gctype)

	return &Waypoint{
		WorldEntity: worldEntity,
	}
}

func NewWaypointFromConfig(config *configtypes.WaypointConfig) *Waypoint {
	waypoint := NewWaypoint(config.FullGCType)

	waypoint.BaseConfig = config

	return waypoint
}
