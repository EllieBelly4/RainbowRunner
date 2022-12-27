// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	lua2 "github.com/yuin/gopher-lua"
)

type IItem interface {
	GetItem() *Item
}

func (i *Item) GetItem() *Item {
	return i
}

func registerLuaItem(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("Item")
	state.SetGlobal("Item", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaItem))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsItem(),
	))
}

func luaMethodsItem() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"modCount": lua.LuaGenericGetSetNumber[IItem](func(v IItem) *int { return &v.GetItem().ModCount }),
		"mod":      lua.LuaGenericGetSetString[IItem](func(v IItem) *string { return &v.GetItem().Mod }),
		// -------------------------------------------------------------------------------------------------------------
		// Unsupported field type ItemType
		// -------------------------------------------------------------------------------------------------------------
		// -------------------------------------------------------------------------------------------------------------
		// Unsupported field type InventoryPosition
		// -------------------------------------------------------------------------------------------------------------
		"index": lua.LuaGenericGetSetNumber[IItem](func(v IItem) *int { return &v.GetItem().Index }),

		"setInventoryPosition": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IItem](l, 1)
			obj := objInterface.GetItem()
			obj.SetInventoryPosition(
				lua.CheckValue[datatypes.Vector2](l, 2),
			)

			return 0
		},

		"writeInit": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IItem](l, 1)
			obj := objInterface.GetItem()
			obj.WriteInit(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"getItem": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IItem](l, 1)
			obj := objInterface.GetItem()
			res0 := obj.GetItem()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	}, luaMethodsManipulator)
}
func newLuaItem(l *lua2.LState) int {
	obj := NewItem(string(l.CheckString(1)),
		lua.CheckValue[ItemType](l, 2),
	)
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("Item"))
	l.Push(ud)
	return 1
}

func (i *Item) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = i

	l.SetMetatable(ud, l.GetTypeMetatable("Item"))
	return ud
}
