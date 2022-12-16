package objects

import (
	"RainbowRunner/internal/actions"
)

type ExecuteActionEvent struct {
	Action       actions.Action
	UnitBehavior IUnitBehavior
}
