// Code generated by scripts/generateluaregistrations DO NOT EDIT.
package actions

import lua2 "github.com/yuin/gopher-lua"

func RegisterAllLuaFunctions(state *lua2.LState) {
	registerLuaActivate(state)
	registerLuaAmbush(state)
	registerLuaAttackTarget2(state)
	registerLuaConvertItemsToGold(state)
	registerLuaDie(state)
	registerLuaDodge(state)
	registerLuaDoEffect(state)
	registerLuaFaceTarget(state)
	registerLuaFlee(state)
	registerLuaFollow(state)
	registerLuaHide(state)
	registerLuaIdle(state)
	registerLuaImmobilize(state)
	registerLuaKill(state)
	registerLuaKnockBack(state)
	registerLuaKnockDown(state)
	registerLuaMoveInDirectionAction(state)
	registerLuaMoveTo(state)
	registerLuaPlayAnimation(state)
	registerLuaRemove(state)
	registerLuaRessurect(state)
	registerLuaRetrieveItem(state)
	registerLuaSearchForAttack(state)
	registerLuaSetBlocking(state)
	registerLuaSpawn(state)
	registerLuaSpawnAnimation(state)
	registerLuaStun(state)
	registerLuaTurnAction(state)
	registerLuaUnSpawn(state)
	registerLuaUse(state)
	registerLuaUseItem(state)
	registerLuaUseItemPosition(state)
	registerLuaUseItemTarget(state)
	registerLuaUsePosition(state)
	registerLuaUseTarget(state)
	registerLuaWait(state)
	registerLuaWander(state)
	registerLuaWarpTo(state)
}