package configtypes

//go:generate go run ../../../scripts/generatelua -type=WaypointConfig -extends=EntityConfig
type WaypointConfig struct {
	*EntityConfig
}

func NewWaypointConfig() *WaypointConfig {
	return &WaypointConfig{
		EntityConfig: NewEntityConfig(EntityConfigTypeWaypoint),
	}
}
