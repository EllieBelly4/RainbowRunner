// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionUse interface {
	GetActionUse() *ActionUse
}

func (a *ActionUse) GetActionUse() *ActionUse {
	return a
}

func registerLuaActionUse(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionUse")
	state.SetGlobal("ActionUse", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionUse))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionUse(),
	))
}

func luaMethodsActionUse() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionUse](l, 1)
			obj := objInterface.GetActionUse()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},
		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionUse](l, 1)
			obj := objInterface.GetActionUse()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"getActionUse": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionUse](l, 1)
			obj := objInterface.GetActionUse()
			res0 := obj.GetActionUse()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaActionUse(l *lua2.LState) int {
	obj := NewActionUse()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionUse"))
	l.Push(ud)
	return 1
}

func (a *ActionUse) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionUse"))
	return ud
}
