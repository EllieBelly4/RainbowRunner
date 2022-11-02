package objects

import (
	lua2 "github.com/yuin/gopher-lua"
)

const luaMerchantTypeName = "Merchant"

func registerLuaMerchant(state *lua2.LState) {
	mt := state.NewTypeMetatable(luaMerchantTypeName)
	state.SetGlobal("Merchant", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaMerchant))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsMerchant(),
	))
}

func luaMethodsMerchant() map[string]lua2.LGFunction {
	return luaDRObjectExtendMethods(entityLuaMerchantMethods)
}

var entityLuaMerchantMethods = map[string]lua2.LGFunction{}

func newLuaMerchant(l *lua2.LState) int {
	Merchant := NewMerchant(l.CheckString(1))

	ud := l.NewUserData()
	ud.Value = Merchant

	l.SetMetatable(ud, l.GetTypeMetatable(luaMerchantTypeName))
	l.Push(ud)
	return 1
}
