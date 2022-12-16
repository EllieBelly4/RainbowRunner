// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionWander interface {
	GetActionWander() *ActionWander
}

func (a *ActionWander) GetActionWander() *ActionWander {
	return a
}

func registerLuaActionWander(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionWander")
	state.SetGlobal("ActionWander", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionWander))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionWander(),
	))
}

func luaMethodsActionWander() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionWander](l, 1)
			obj := objInterface.GetActionWander()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},
		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionWander](l, 1)
			obj := objInterface.GetActionWander()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"getActionWander": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionWander](l, 1)
			obj := objInterface.GetActionWander()
			res0 := obj.GetActionWander()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaActionWander(l *lua2.LState) int {
	obj := NewActionWander()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionWander"))
	l.Push(ud)
	return 1
}

func (a *ActionWander) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionWander"))
	return ud
}
