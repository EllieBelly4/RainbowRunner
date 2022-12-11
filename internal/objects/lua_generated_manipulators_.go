package objects

/**
 * This file is generated by scripts/generatelua/generatelua.go
 * DO NOT EDIT
 */

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

func registerLuaManipulators(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("Manipulators")
	state.SetGlobal("Manipulators", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaManipulators))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsManipulators(),
	))
}

func luaMethodsManipulators() map[string]lua2.LGFunction {
	return luaMethodsExtend(map[string]lua2.LGFunction{
		"writeInit": func(l *lua2.LState) int {
			obj := lua.CheckReferenceValue[Manipulators](l, 1)
			obj.WriteInit(
				lua.CheckReferenceValue[byter.Byter](l, 1),
			)

			return 0
		},
		"removeEquipmentByID": func(l *lua2.LState) int {
			obj := lua.CheckReferenceValue[Manipulators](l, 1)
			obj.RemoveEquipmentByID(uint32(l.CheckNumber(1)))

			return 0
		},
		"writeRemoveItem": func(l *lua2.LState) int {
			obj := lua.CheckReferenceValue[Manipulators](l, 1)
			obj.WriteRemoveItem(
				lua.CheckReferenceValue[byter.Byter](l, 1),
				lua.CheckValue[types.EquipmentSlot](l, 2),
			)

			return 0
		},
		"writeAddItem": func(l *lua2.LState) int {
			obj := lua.CheckReferenceValue[Manipulators](l, 1)
			obj.WriteAddItem(
				lua.CheckReferenceValue[byter.Byter](l, 1),
				lua.CheckReferenceValue[Equipment](l, 2),
			)

			return 0
		},
	}, luaMethodsComponent)
}
func newLuaManipulators(l *lua2.LState) int {
	obj := NewManipulators(string(l.CheckString(1)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("Manipulators"))
	l.Push(ud)
	return 1
}
