// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/internal/types/configtypes"
	lua2 "github.com/yuin/gopher-lua"
)

type IMerchantInventory interface {
	GetMerchantInventory() *MerchantInventory
}

func (m *MerchantInventory) GetMerchantInventory() *MerchantInventory {
	return m
}

func registerLuaMerchantInventory(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("MerchantInventory")
	state.SetGlobal("MerchantInventory", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaMerchantInventory))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsMerchantInventory(),
	))
}

func luaMethodsMerchantInventory() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"baseConfig": lua.LuaGenericGetSetValueAny[IMerchantInventory](func(v IMerchantInventory) **configtypes.MerchantInventoryConfig {
			return &v.GetMerchantInventory().BaseConfig
		}),

		"getMerchantInventory": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IMerchantInventory](l, 1)
			obj := objInterface.GetMerchantInventory()
			res0 := obj.GetMerchantInventory()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	}, luaMethodsInventory)
}
func newLuaMerchantInventory(l *lua2.LState) int {
	obj := NewMerchantInventory(string(l.CheckString(1)), byte(l.CheckNumber(2)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("MerchantInventory"))
	l.Push(ud)
	return 1
}

func (m *MerchantInventory) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = m

	l.SetMetatable(ud, l.GetTypeMetatable("MerchantInventory"))
	return ud
}
