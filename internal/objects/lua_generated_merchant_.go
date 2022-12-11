package objects

/**
 * This file is generated by scripts/generatelua/generatelua.go
 * DO NOT EDIT
 */

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

func registerLuaMerchant(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("Merchant")
	state.SetGlobal("Merchant", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaMerchant))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsMerchant(),
	))
}

func luaMethodsMerchant() map[string]lua2.LGFunction {
	return luaMethodsExtend(map[string]lua2.LGFunction{
		"writeInit": func(l *lua2.LState) int {
			obj := lua.CheckReferenceValue[Merchant](l, 1)
			obj.WriteInit(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"getInventoryByID": func(l *lua2.LState) int {
			obj := lua.CheckReferenceValue[Merchant](l, 1)
			res0 := obj.GetInventoryByID(
				lua.CheckValue[byte](l, 2),
			)
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("Inventory"))
			l.Push(ud)

			return 1
		},
		"toLua": func(l *lua2.LState) int {
			obj := lua.CheckReferenceValue[Merchant](l, 1)
			res0 := obj.ToLua(
				lua.CheckReferenceValue[lua2.LState](l, 2),
			)
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("lua2.LValue"))
			l.Push(ud)

			return 1
		},
		"Container": func(l *lua2.LState) int {
			obj := lua.CheckReferenceValue[Merchant](l, 1)
			l.Push(obj.Container.ToLua(l))
			return 1
		},
	})
}
func newLuaMerchant(l *lua2.LState) int {
	obj := NewMerchant(string(l.CheckString(1)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("Merchant"))
	l.Push(ud)
	return 1
}

func (m *Merchant) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = m

	l.SetMetatable(ud, l.GetTypeMetatable("Merchant"))
	return ud
}
