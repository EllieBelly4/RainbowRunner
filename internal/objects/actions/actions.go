package actions

import (
	byter "RainbowRunner/pkg/byter"
)

//go:generate stringer -type=BehaviourAction
type BehaviourAction byte

const (
	BehaviourActionMoveTo                BehaviourAction = 1
	BehaviourActionSpawn                 BehaviourAction = 4
	BehaviourActionActivate              BehaviourAction = 6
	BehaviourActionKnockBack             BehaviourAction = 10
	BehaviourActionKnockDown             BehaviourAction = 11
	BehaviourActionStun                  BehaviourAction = 12
	BehaviourActionSearchForAttack       BehaviourAction = 13
	BehaviourActionSpawnAnimation        BehaviourAction = 14
	BehaviourActionUnSpawn               BehaviourAction = 15
	BehaviourActionDodge                 BehaviourAction = 16
	BehaviourActionWarpTo                BehaviourAction = 17
	BehaviourActionAmbush                BehaviourAction = 18
	BehaviourActionMoveInDirectionAction BehaviourAction = 19
	BehaviourActionTurnAction            BehaviourAction = 20
	BehaviourActionWander                BehaviourAction = 21
	BehaviourActionFollow                BehaviourAction = 22
	BehaviourActionPlayAnimation         BehaviourAction = 32
	BehaviourActionFaceTarget            BehaviourAction = 33
	BehaviourActionWait                  BehaviourAction = 34
	BehaviourActionImmobilize            BehaviourAction = 35
	BehaviourActionIdle                  BehaviourAction = 47
	BehaviourActionRessurect             BehaviourAction = 48
	BehaviourActionRemove                BehaviourAction = 64
	BehaviourActionHide                  BehaviourAction = 65
	BehaviourActionSetBlocking           BehaviourAction = 69
	BehaviourActionFlee                  BehaviourAction = 79
	BehaviourActionUseTarget             BehaviourAction = 80
	BehaviourActionUsePosition           BehaviourAction = 81
	BehaviourActionUse                   BehaviourAction = 82
	BehaviourActionUseItemTarget         BehaviourAction = 83
	BehaviourActionUseItemPosition       BehaviourAction = 84
	BehaviourActionUseItem               BehaviourAction = 85
	BehaviourActionDoEffect              BehaviourAction = 131
	BehaviourActionRetrieveItem          BehaviourAction = 160
	BehaviourActionConvertItemsToGold    BehaviourAction = 161
	BehaviourActionAttackTarget2         BehaviourAction = 240
	BehaviourActionKill                  BehaviourAction = 254
	BehaviourActionDie                   BehaviourAction = 255
)

// Action See [resources/Docs/v2/Behavior.md]
type Action interface {
	OpCode() BehaviourAction
	Init(body *byter.Byter, sessionID byte)
}
