// Code generated by scripts/generatelua DO NOT EDIT.
package configtypes

import (
	lua "RainbowRunner/internal/lua"
	lua2 "github.com/yuin/gopher-lua"
)

type IZoneDefConfig interface {
	GetZoneDefConfig() *ZoneDefConfig
}

func (z *ZoneDefConfig) GetZoneDefConfig() *ZoneDefConfig {
	return z
}

func registerLuaZoneDefConfig(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ZoneDefConfig")
	state.SetGlobal("ZoneDefConfig", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaZoneDefConfig))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsZoneDefConfig(),
	))
}

func luaMethodsZoneDefConfig() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"label":                 lua.LuaGenericGetSetString[IZoneDefConfig](func(v IZoneDefConfig) *string { return &v.GetZoneDefConfig().Label }),
		"name":                  lua.LuaGenericGetSetString[IZoneDefConfig](func(v IZoneDefConfig) *string { return &v.GetZoneDefConfig().Name }),
		"updateFrequency":       lua.LuaGenericGetSetNumber[IZoneDefConfig](func(v IZoneDefConfig) *int { return &v.GetZoneDefConfig().UpdateFrequency }),
		"private":               lua.LuaGenericGetSetBool[IZoneDefConfig](func(v IZoneDefConfig) *bool { return &v.GetZoneDefConfig().Private }),
		"isLegendary":           lua.LuaGenericGetSetBool[IZoneDefConfig](func(v IZoneDefConfig) *bool { return &v.GetZoneDefConfig().IsLegendary }),
		"isTown":                lua.LuaGenericGetSetBool[IZoneDefConfig](func(v IZoneDefConfig) *bool { return &v.GetZoneDefConfig().IsTown }),
		"deathPenalty":          lua.LuaGenericGetSetBool[IZoneDefConfig](func(v IZoneDefConfig) *bool { return &v.GetZoneDefConfig().DeathPenalty }),
		"useEliteGenerators":    lua.LuaGenericGetSetBool[IZoneDefConfig](func(v IZoneDefConfig) *bool { return &v.GetZoneDefConfig().UseEliteGenerators }),
		"sendBankContents":      lua.LuaGenericGetSetBool[IZoneDefConfig](func(v IZoneDefConfig) *bool { return &v.GetZoneDefConfig().SendBankContents }),
		"maxOccupancy":          lua.LuaGenericGetSetNumber[IZoneDefConfig](func(v IZoneDefConfig) *int { return &v.GetZoneDefConfig().MaxOccupancy }),
		"maxLevel":              lua.LuaGenericGetSetNumber[IZoneDefConfig](func(v IZoneDefConfig) *int { return &v.GetZoneDefConfig().MaxLevel }),
		"minLevel":              lua.LuaGenericGetSetNumber[IZoneDefConfig](func(v IZoneDefConfig) *int { return &v.GetZoneDefConfig().MinLevel }),
		"respawnZone":           lua.LuaGenericGetSetString[IZoneDefConfig](func(v IZoneDefConfig) *string { return &v.GetZoneDefConfig().RespawnZone }),
		"respawnSpawnPoint":     lua.LuaGenericGetSetString[IZoneDefConfig](func(v IZoneDefConfig) *string { return &v.GetZoneDefConfig().RespawnSpawnPoint }),
		"allowPvPAnnouncements": lua.LuaGenericGetSetBool[IZoneDefConfig](func(v IZoneDefConfig) *bool { return &v.GetZoneDefConfig().AllowPvPAnnouncements }),
		"pvptype":               lua.LuaGenericGetSetNumber[IZoneDefConfig](func(v IZoneDefConfig) *int { return &v.GetZoneDefConfig().PVPType }),
		"pvpmatchType":          lua.LuaGenericGetSetString[IZoneDefConfig](func(v IZoneDefConfig) *string { return &v.GetZoneDefConfig().PVPMatchType }),
		"allowDuelRequest":      lua.LuaGenericGetSetBool[IZoneDefConfig](func(v IZoneDefConfig) *bool { return &v.GetZoneDefConfig().AllowDuelRequest }),
		"entryModifier":         lua.LuaGenericGetSetString[IZoneDefConfig](func(v IZoneDefConfig) *string { return &v.GetZoneDefConfig().EntryModifier }),

		"getZoneDefConfig": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IZoneDefConfig](l, 1)
			obj := objInterface.GetZoneDefConfig()
			res0 := obj.GetZoneDefConfig()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaZoneDefConfig(l *lua2.LState) int {
	obj := NewZoneDefConfig()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ZoneDefConfig"))
	l.Push(ud)
	return 1
}

func (z *ZoneDefConfig) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = z

	l.SetMetatable(ud, l.GetTypeMetatable("ZoneDefConfig"))
	return ud
}
