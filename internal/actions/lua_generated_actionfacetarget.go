// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionFaceTarget interface {
	GetActionFaceTarget() *ActionFaceTarget
}

func (a *ActionFaceTarget) GetActionFaceTarget() *ActionFaceTarget {
	return a
}

func registerLuaActionFaceTarget(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionFaceTarget")
	state.SetGlobal("ActionFaceTarget", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionFaceTarget))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionFaceTarget(),
	))
}

func luaMethodsActionFaceTarget() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{

		"targetID": lua.LuaGenericGetSetValueAny[IActionFaceTarget](func(v IActionFaceTarget) *uint16 { return &v.GetActionFaceTarget().TargetID }),

		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionFaceTarget](l, 1)
			obj := objInterface.GetActionFaceTarget()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},

		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionFaceTarget](l, 1)
			obj := objInterface.GetActionFaceTarget()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"getActionFaceTarget": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionFaceTarget](l, 1)
			obj := objInterface.GetActionFaceTarget()
			res0 := obj.GetActionFaceTarget()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaActionFaceTarget(l *lua2.LState) int {
	obj := NewActionFaceTarget(uint16(l.CheckNumber(1)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionFaceTarget"))
	l.Push(ud)
	return 1
}

func (a *ActionFaceTarget) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionFaceTarget"))
	return ud
}
