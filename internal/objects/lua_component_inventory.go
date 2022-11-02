package objects

import (
	lua2 "github.com/yuin/gopher-lua"
)

const luaInventoryTypeName = "Inventory"

func registerLuaInventory(state *lua2.LState) {
	mt := state.NewTypeMetatable(luaInventoryTypeName)
	state.SetGlobal("Inventory", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaInventory))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsInventory(),
	))

	state.SetGlobal("MerchantInventory", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaInventory))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsInventory(),
	))
}

func luaMethodsInventory() map[string]lua2.LGFunction {
	return luaDRObjectExtendMethods(entityLuaInventoryMethods)
}

var entityLuaInventoryMethods = map[string]lua2.LGFunction{}

func newLuaMerchantInventory(l *lua2.LState) int {
	inventory := NewMerchantInventory(l.CheckString(1), byte(l.CheckNumber(2)))

	ud := l.NewUserData()
	ud.Value = inventory

	l.SetMetatable(ud, l.GetTypeMetatable(luaInventoryTypeName))
	l.Push(ud)
	return 1
}

func newLuaInventory(l *lua2.LState) int {
	inventory := NewInventory(l.CheckString(1), byte(l.CheckNumber(2)))

	ud := l.NewUserData()
	ud.Value = inventory

	l.SetMetatable(ud, l.GetTypeMetatable(luaInventoryTypeName))
	l.Push(ud)
	return 1
}
