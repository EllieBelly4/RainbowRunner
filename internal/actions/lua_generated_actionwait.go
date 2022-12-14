// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionWait interface {
	GetActionWait() *ActionWait
}

func (a *ActionWait) GetActionWait() *ActionWait {
	return a
}

func registerLuaActionWait(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionWait")
	state.SetGlobal("ActionWait", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionWait))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionWait(),
	))
}

func luaMethodsActionWait() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{

		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionWait](l, 1)
			obj := objInterface.GetActionWait()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},

		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionWait](l, 1)
			obj := objInterface.GetActionWait()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"getActionWait": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionWait](l, 1)
			obj := objInterface.GetActionWait()
			res0 := obj.GetActionWait()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaActionWait(l *lua2.LState) int {
	obj := NewActionWait()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionWait"))
	l.Push(ud)
	return 1
}

func (a *ActionWait) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionWait"))
	return ud
}
