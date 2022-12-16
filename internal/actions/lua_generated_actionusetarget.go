// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionUseTarget interface {
	GetActionUseTarget() *ActionUseTarget
}

func (a *ActionUseTarget) GetActionUseTarget() *ActionUseTarget {
	return a
}

func registerLuaActionUseTarget(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionUseTarget")
	state.SetGlobal("ActionUseTarget", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionUseTarget))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionUseTarget(),
	))
}

func luaMethodsActionUseTarget() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionUseTarget](l, 1)
			obj := objInterface.GetActionUseTarget()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},
		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionUseTarget](l, 1)
			obj := objInterface.GetActionUseTarget()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
	})
}
func newLuaActionUseTarget(l *lua2.LState) int {
	obj := NewActionUseTarget()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionUseTarget"))
	l.Push(ud)
	return 1
}

func (a *ActionUseTarget) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionUseTarget"))
	return ud
}
