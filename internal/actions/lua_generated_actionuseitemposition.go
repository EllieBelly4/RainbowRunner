// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionUseItemPosition interface {
	GetActionUseItemPosition() *ActionUseItemPosition
}

func (a *ActionUseItemPosition) GetActionUseItemPosition() *ActionUseItemPosition {
	return a
}

func registerLuaActionUseItemPosition(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionUseItemPosition")
	state.SetGlobal("ActionUseItemPosition", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionUseItemPosition))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionUseItemPosition(),
	))
}

func luaMethodsActionUseItemPosition() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionUseItemPosition](l, 1)
			obj := objInterface.GetActionUseItemPosition()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},
		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionUseItemPosition](l, 1)
			obj := objInterface.GetActionUseItemPosition()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
	})
}
func newLuaActionUseItemPosition(l *lua2.LState) int {
	obj := NewActionUseItemPosition()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionUseItemPosition"))
	l.Push(ud)
	return 1
}

func (a *ActionUseItemPosition) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionUseItemPosition"))
	return ud
}
