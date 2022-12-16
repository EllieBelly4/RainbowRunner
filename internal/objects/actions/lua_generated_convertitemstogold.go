// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IConvertItemsToGold interface {
	GetConvertItemsToGold() *ConvertItemsToGold
}

func (c *ConvertItemsToGold) GetConvertItemsToGold() *ConvertItemsToGold {
	return c
}

func registerLuaConvertItemsToGold(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ConvertItemsToGold")
	state.SetGlobal("ConvertItemsToGold", mt)
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsConvertItemsToGold(),
	))
}

func luaMethodsConvertItemsToGold() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IConvertItemsToGold](l, 1)
			obj := objInterface.GetConvertItemsToGold()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},
		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IConvertItemsToGold](l, 1)
			obj := objInterface.GetConvertItemsToGold()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
	})
}

func (c *ConvertItemsToGold) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = c

	l.SetMetatable(ud, l.GetTypeMetatable("ConvertItemsToGold"))
	return ud
}