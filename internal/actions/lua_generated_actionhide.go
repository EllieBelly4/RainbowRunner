// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionHide interface {
	GetActionHide() *ActionHide
}

func (a *ActionHide) GetActionHide() *ActionHide {
	return a
}

func registerLuaActionHide(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionHide")
	state.SetGlobal("ActionHide", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionHide))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionHide(),
	))
}

func luaMethodsActionHide() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionHide](l, 1)
			obj := objInterface.GetActionHide()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},
		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionHide](l, 1)
			obj := objInterface.GetActionHide()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
	})
}
func newLuaActionHide(l *lua2.LState) int {
	obj := NewActionHide()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionHide"))
	l.Push(ud)
	return 1
}

func (a *ActionHide) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionHide"))
	return ud
}
