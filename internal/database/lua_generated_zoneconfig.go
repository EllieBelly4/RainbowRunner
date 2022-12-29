// Code generated by scripts/generatelua DO NOT EDIT.
package database

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/internal/types/configtypes"
	lua2 "github.com/yuin/gopher-lua"
)

type IZoneConfig interface {
	GetZoneConfig() *ZoneConfig
}

func (z *ZoneConfig) GetZoneConfig() *ZoneConfig {
	return z
}

func registerLuaZoneConfig(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ZoneConfig")
	state.SetGlobal("ZoneConfig", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaZoneConfig))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsZoneConfig(),
	))
}

func luaMethodsZoneConfig() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"name": lua.LuaGenericGetSetString[IZoneConfig](func(v IZoneConfig) *string { return &v.GetZoneConfig().Name }),
		// -------------------------------------------------------------------------------------------------------------
		// Unsupported field type NPCs
		// -------------------------------------------------------------------------------------------------------------
		// -------------------------------------------------------------------------------------------------------------
		// Unsupported field type Checkpoints
		// -------------------------------------------------------------------------------------------------------------
		"world":  lua.LuaGenericGetSetValue[IZoneConfig, *configtypes.WorldConfig](func(v IZoneConfig) **configtypes.WorldConfig { return &v.GetZoneConfig().World }),
		"gctype": lua.LuaGenericGetSetString[IZoneConfig](func(v IZoneConfig) *string { return &v.GetZoneConfig().GCType }),

		"getZoneConfig": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IZoneConfig](l, 1)
			obj := objInterface.GetZoneConfig()
			res0 := obj.GetZoneConfig()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},

		"getAllNPCs": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IZoneConfig](l, 1)
			obj := objInterface.GetZoneConfig()
			res0 := obj.GetAllNPCs()
			res0Array := l.NewTable()

			for _, res0 := range res0 {
				if res0 != nil {
					res0Array.Append(res0.ToLua(l))
				} else {
					res0Array.Append(lua2.LNil)
				}
			}

			l.Push(res0Array)

			return 1
		},
	})
}
func newLuaZoneConfig(l *lua2.LState) int {
	obj := NewZoneConfig(string(l.CheckString(1)), string(l.CheckString(2)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ZoneConfig"))
	l.Push(ud)
	return 1
}

func (z *ZoneConfig) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = z

	l.SetMetatable(ud, l.GetTypeMetatable("ZoneConfig"))
	return ud
}
