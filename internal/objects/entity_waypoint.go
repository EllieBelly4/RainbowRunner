package objects

//go:generate go run ../../scripts/generatelua -type=Waypoint -extends=WorldEntity
type Waypoint struct {
	*WorldEntity
}

func NewWaypoint(gctype string) *Waypoint {
	worldEntity := NewWorldEntity(gctype)

	return &Waypoint{
		WorldEntity: worldEntity,
	}
}
