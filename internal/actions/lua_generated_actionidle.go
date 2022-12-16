// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionIdle interface {
	GetActionIdle() *ActionIdle
}

func (a *ActionIdle) GetActionIdle() *ActionIdle {
	return a
}

func registerLuaActionIdle(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionIdle")
	state.SetGlobal("ActionIdle", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionIdle))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionIdle(),
	))
}

func luaMethodsActionIdle() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionIdle](l, 1)
			obj := objInterface.GetActionIdle()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},
		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionIdle](l, 1)
			obj := objInterface.GetActionIdle()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
	})
}
func newLuaActionIdle(l *lua2.LState) int {
	obj := NewActionIdle()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionIdle"))
	l.Push(ud)
	return 1
}

func (a *ActionIdle) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionIdle"))
	return ud
}
