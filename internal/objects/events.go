package objects

import "RainbowRunner/internal/objects/actions"

type ExecuteActionEvent struct {
	Action       actions.Action
	UnitBehavior IUnitBehavior
}
