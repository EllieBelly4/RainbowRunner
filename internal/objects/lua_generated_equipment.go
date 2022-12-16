// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IEquipment interface {
	GetEquipment() *Equipment
}

func (e *Equipment) GetEquipment() *Equipment {
	return e
}

func registerLuaEquipment(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("Equipment")
	state.SetGlobal("Equipment", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaEquipment))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsEquipment(),
	))
}

func luaMethodsEquipment() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		// -------------------------------------------------------------------------------------------------------------
		// Unsupported field type Slot
		// -------------------------------------------------------------------------------------------------------------
		"writeInit": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IEquipment](l, 1)
			obj := objInterface.GetEquipment()
			obj.WriteInit(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"writeManipulatorInit": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IEquipment](l, 1)
			obj := objInterface.GetEquipment()
			obj.WriteManipulatorInit(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"getEquipment": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IEquipment](l, 1)
			obj := objInterface.GetEquipment()
			res0 := obj.GetEquipment()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	}, luaMethodsItem)
}
func newLuaEquipment(l *lua2.LState) int {
	obj := NewEquipment(string(l.CheckString(1)), string(l.CheckString(2)),
		lua.CheckValue[ItemType](l, 3),
		lua.CheckValue[types.EquipmentSlot](l, 4),
	)
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("Equipment"))
	l.Push(ud)
	return 1
}

func (e *Equipment) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = e

	l.SetMetatable(ud, l.GetTypeMetatable("Equipment"))
	return ud
}
