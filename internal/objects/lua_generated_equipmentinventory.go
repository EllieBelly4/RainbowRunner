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

type IEquipmentInventory interface {
	GetEquipmentInventory() *EquipmentInventory
}

func (e *EquipmentInventory) GetEquipmentInventory() *EquipmentInventory {
	return e
}

func registerLuaEquipmentInventory(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("EquipmentInventory")
	state.SetGlobal("EquipmentInventory", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaEquipmentInventory))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsEquipmentInventory(),
	))
}

func luaMethodsEquipmentInventory() map[string]lua2.LGFunction {
	return luaMethodsExtend(map[string]lua2.LGFunction{
		"addChild": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IEquipmentInventory](l, 1)
			obj := objInterface.GetEquipmentInventory()
			obj.AddChild(
				lua.CheckValue[DRObject](l, 2),
			)

			return 0
		},
		"readUpdate": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IEquipmentInventory](l, 1)
			obj := objInterface.GetEquipmentInventory()
			res0 := obj.ReadUpdate(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("error"))
			l.Push(ud)

			return 1
		},
		"removeEquipmentBySlot": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IEquipmentInventory](l, 1)
			obj := objInterface.GetEquipmentInventory()
			res0 := obj.RemoveEquipmentBySlot(
				lua.CheckValue[types.EquipmentSlot](l, 2),
			)
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("Equipment"))
			l.Push(ud)

			return 1
		},
		"getEquipment": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IEquipmentInventory](l, 1)
			obj := objInterface.GetEquipmentInventory()
			res0 := obj.GetEquipment()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("[]*Equipment"))
			l.Push(ud)

			return 1
		},
		"getEquipmentInventory": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IEquipmentInventory](l, 1)
			obj := objInterface.GetEquipmentInventory()
			res0 := obj.GetEquipmentInventory()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("EquipmentInventory"))
			l.Push(ud)

			return 1
		},
	}, luaMethodsComponent)
}
func newLuaEquipmentInventory(l *lua2.LState) int {
	obj := NewEquipmentInventory(string(l.CheckString(1)),
		lua.CheckReferenceValue[Avatar](l, 2),
	)
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("EquipmentInventory"))
	l.Push(ud)
	return 1
}

func (e *EquipmentInventory) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = e

	l.SetMetatable(ud, l.GetTypeMetatable("EquipmentInventory"))
	return ud
}