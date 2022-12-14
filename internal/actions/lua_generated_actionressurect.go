// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionRessurect interface {
	GetActionRessurect() *ActionRessurect
}

func (a *ActionRessurect) GetActionRessurect() *ActionRessurect {
	return a
}

func registerLuaActionRessurect(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionRessurect")
	state.SetGlobal("ActionRessurect", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionRessurect))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionRessurect(),
	))
}

func luaMethodsActionRessurect() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{

		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionRessurect](l, 1)
			obj := objInterface.GetActionRessurect()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},

		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionRessurect](l, 1)
			obj := objInterface.GetActionRessurect()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"getActionRessurect": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionRessurect](l, 1)
			obj := objInterface.GetActionRessurect()
			res0 := obj.GetActionRessurect()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaActionRessurect(l *lua2.LState) int {
	obj := NewActionRessurect()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionRessurect"))
	l.Push(ud)
	return 1
}

func (a *ActionRessurect) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionRessurect"))
	return ud
}
