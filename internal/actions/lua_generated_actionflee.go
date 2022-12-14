// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionFlee interface {
	GetActionFlee() *ActionFlee
}

func (a *ActionFlee) GetActionFlee() *ActionFlee {
	return a
}

func registerLuaActionFlee(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionFlee")
	state.SetGlobal("ActionFlee", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionFlee))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionFlee(),
	))
}

func luaMethodsActionFlee() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{

		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionFlee](l, 1)
			obj := objInterface.GetActionFlee()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},

		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionFlee](l, 1)
			obj := objInterface.GetActionFlee()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"getActionFlee": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionFlee](l, 1)
			obj := objInterface.GetActionFlee()
			res0 := obj.GetActionFlee()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaActionFlee(l *lua2.LState) int {
	obj := NewActionFlee()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionFlee"))
	l.Push(ud)
	return 1
}

func (a *ActionFlee) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionFlee"))
	return ud
}
