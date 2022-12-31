// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionActivate interface {
	GetActionActivate() *ActionActivate
}

func (a *ActionActivate) GetActionActivate() *ActionActivate {
	return a
}

func registerLuaActionActivate(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionActivate")
	state.SetGlobal("ActionActivate", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionActivate))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionActivate(),
	))
}

func luaMethodsActionActivate() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{

		"targetEntityID": lua.LuaGenericGetSetValueAny[IActionActivate](func(v IActionActivate) *uint16 { return &v.GetActionActivate().TargetEntityID }),

		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionActivate](l, 1)
			obj := objInterface.GetActionActivate()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},

		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionActivate](l, 1)
			obj := objInterface.GetActionActivate()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"initWithoutOpCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionActivate](l, 1)
			obj := objInterface.GetActionActivate()
			obj.InitWithoutOpCode(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"getActionActivate": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionActivate](l, 1)
			obj := objInterface.GetActionActivate()
			res0 := obj.GetActionActivate()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaActionActivate(l *lua2.LState) int {
	obj := NewActionActivate()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionActivate"))
	l.Push(ud)
	return 1
}

func (a *ActionActivate) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionActivate"))
	return ud
}
