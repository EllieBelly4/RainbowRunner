package actions

var (
	sessionIDMap = map[BehaviourAction]bool{
		BehaviourActionMoveTo:                true,
		BehaviourActionSpawn:                 false,
		BehaviourActionActivate:              true,
		BehaviourActionKnockBack:             false,
		BehaviourActionKnockDown:             false,
		BehaviourActionStun:                  false,
		BehaviourActionSearchForAttack:       false,
		BehaviourActionSpawnAnimation:        false,
		BehaviourActionUnSpawn:               false,
		BehaviourActionDodge:                 false,
		BehaviourActionWarpTo:                true,
		BehaviourActionAmbush:                false,
		BehaviourActionMoveInDirectionAction: false,
		BehaviourActionTurnAction:            false,
		BehaviourActionWander:                false,
		BehaviourActionFollow:                false,
		BehaviourActionPlayAnimation:         false,
		BehaviourActionFaceTarget:            false,
		BehaviourActionWait:                  false,
		BehaviourActionImmobilize:            false,
		BehaviourActionIdle:                  false,
		BehaviourActionRessurect:             false,
		BehaviourActionRemove:                false,
		BehaviourActionHide:                  false,
		BehaviourActionSetBlocking:           false,
		BehaviourActionFlee:                  false,
		BehaviourActionUseTarget:             false,
		BehaviourActionUsePosition:           true,
		BehaviourActionUse:                   false,
		BehaviourActionUseItemTarget:         false,
		BehaviourActionUseItemPosition:       false,
		BehaviourActionUseItem:               false,
		BehaviourActionDoEffect:              false,
		BehaviourActionRetrieveItem:          false,
		BehaviourActionConvertItemsToGold:    false,
		BehaviourActionAttackTarget2:         false,
		BehaviourActionKill:                  false,
		BehaviourActionDie:                   false,
	}
)

func UsesSessionID(action BehaviourAction) bool {
	return sessionIDMap[action]
}
