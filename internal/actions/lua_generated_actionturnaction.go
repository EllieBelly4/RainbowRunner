// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionTurnAction interface {
	GetActionTurnAction() *ActionTurnAction
}

func (a *ActionTurnAction) GetActionTurnAction() *ActionTurnAction {
	return a
}

func registerLuaActionTurnAction(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionTurnAction")
	state.SetGlobal("ActionTurnAction", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionTurnAction))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionTurnAction(),
	))
}

func luaMethodsActionTurnAction() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionTurnAction](l, 1)
			obj := objInterface.GetActionTurnAction()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},
		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionTurnAction](l, 1)
			obj := objInterface.GetActionTurnAction()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"getActionTurnAction": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionTurnAction](l, 1)
			obj := objInterface.GetActionTurnAction()
			res0 := obj.GetActionTurnAction()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaActionTurnAction(l *lua2.LState) int {
	obj := NewActionTurnAction()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionTurnAction"))
	l.Push(ud)
	return 1
}

func (a *ActionTurnAction) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionTurnAction"))
	return ud
}