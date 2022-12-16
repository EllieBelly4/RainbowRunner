// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionAmbush interface {
	GetActionAmbush() *ActionAmbush
}

func (a *ActionAmbush) GetActionAmbush() *ActionAmbush {
	return a
}

func registerLuaActionAmbush(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionAmbush")
	state.SetGlobal("ActionAmbush", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionAmbush))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionAmbush(),
	))
}

func luaMethodsActionAmbush() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionAmbush](l, 1)
			obj := objInterface.GetActionAmbush()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},
		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionAmbush](l, 1)
			obj := objInterface.GetActionAmbush()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
	})
}
func newLuaActionAmbush(l *lua2.LState) int {
	obj := NewActionAmbush()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionAmbush"))
	l.Push(ud)
	return 1
}

func (a *ActionAmbush) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionAmbush"))
	return ud
}
