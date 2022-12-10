package objects

import (
	lua2 "github.com/yuin/gopher-lua"
)

const luaZonePortalTypeName = "ZonePortal"

func registerLuaZonePortal(state *lua2.LState) {
	mt := state.NewTypeMetatable(luaZonePortalTypeName)
	state.SetGlobal("ZonePortal", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaZonePortal))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsZonePortal(),
	))
}

func luaMethodsZonePortal() map[string]lua2.LGFunction {
	return luaMethodsExtend(map[string]lua2.LGFunction{
		"width":  luaGenericGetSetNumber[*ZonePortal, uint16](func(v *ZonePortal) *uint16 { return &v.Width }),
		"height": luaGenericGetSetNumber[*ZonePortal, uint16](func(v *ZonePortal) *uint16 { return &v.Height }),
		"unk4":   luaGenericGetSetNumber[*ZonePortal, uint32](func(v *ZonePortal) *uint32 { return &v.Unk4 }),
	}, luaMethodsWorldEntity)
}

func newLuaZonePortal(l *lua2.LState) int {
	zonePortal := NewZonePortal(l.CheckString(1), l.CheckString(2))

	ud := l.NewUserData()
	ud.Value = zonePortal

	l.SetMetatable(ud, l.GetTypeMetatable(luaZonePortalTypeName))
	l.Push(ud)
	return 1
}
