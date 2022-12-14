// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionSetBlocking interface {
	GetActionSetBlocking() *ActionSetBlocking
}

func (a *ActionSetBlocking) GetActionSetBlocking() *ActionSetBlocking {
	return a
}

func registerLuaActionSetBlocking(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionSetBlocking")
	state.SetGlobal("ActionSetBlocking", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionSetBlocking))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionSetBlocking(),
	))
}

func luaMethodsActionSetBlocking() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{

		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionSetBlocking](l, 1)
			obj := objInterface.GetActionSetBlocking()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},

		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionSetBlocking](l, 1)
			obj := objInterface.GetActionSetBlocking()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"getActionSetBlocking": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionSetBlocking](l, 1)
			obj := objInterface.GetActionSetBlocking()
			res0 := obj.GetActionSetBlocking()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaActionSetBlocking(l *lua2.LState) int {
	obj := NewActionSetBlocking()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionSetBlocking"))
	l.Push(ud)
	return 1
}

func (a *ActionSetBlocking) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionSetBlocking"))
	return ud
}
