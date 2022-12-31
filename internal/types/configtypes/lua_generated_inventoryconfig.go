// Code generated by scripts/generatelua DO NOT EDIT.
package configtypes

import (
	lua "RainbowRunner/internal/lua"
	lua2 "github.com/yuin/gopher-lua"
)

type IInventoryConfig interface {
	GetInventoryConfig() *InventoryConfig
}

func (i *InventoryConfig) GetInventoryConfig() *InventoryConfig {
	return i
}

func registerLuaInventoryConfig(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("InventoryConfig")
	state.SetGlobal("InventoryConfig", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaInventoryConfig))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsInventoryConfig(),
	))
}

func luaMethodsInventoryConfig() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{

		"id": lua.LuaGenericGetSetValueAny[IInventoryConfig](func(v IInventoryConfig) *int { return &v.GetInventoryConfig().ID }),

		"description": lua.LuaGenericGetSetValueAny[IInventoryConfig](func(v IInventoryConfig) **InventoryDescConfig { return &v.GetInventoryConfig().Description }),

		"gctype": lua.LuaGenericGetSetValueAny[IInventoryConfig](func(v IInventoryConfig) *string { return &v.GetInventoryConfig().GCType }),

		"getInventoryConfig": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IInventoryConfig](l, 1)
			obj := objInterface.GetInventoryConfig()
			res0 := obj.GetInventoryConfig()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaInventoryConfig(l *lua2.LState) int {
	obj := NewInventoryConfig(string(l.CheckString(1)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("InventoryConfig"))
	l.Push(ud)
	return 1
}

func (i *InventoryConfig) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = i

	l.SetMetatable(ud, l.GetTypeMetatable("InventoryConfig"))
	return ud
}
