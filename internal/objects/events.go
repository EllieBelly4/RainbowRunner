package objects

import (
	"RainbowRunner/internal/actions"
	"RainbowRunner/pkg/datatypes"
	"RainbowRunner/pkg/datatypes/drfloat"
)

type ExecuteActionEvent struct {
	Action       actions.Action
	UnitBehavior IUnitBehavior
}

type PlayerEnteredZoneEvent struct {
	Player *Player
	Zone   *Zone
}

type PlayerMoveEvent struct {
	UnitBehavior *UnitBehavior
	UpdateType   byte
	PrevPosition datatypes.Vector3Float32
	PrevHeading  drfloat.DRFloat
	NewPosition  datatypes.Vector3Float32
	NewHeading   drfloat.DRFloat
}
