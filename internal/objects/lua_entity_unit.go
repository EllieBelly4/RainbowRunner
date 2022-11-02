package objects

import (
	lua2 "github.com/yuin/gopher-lua"
)

const luaUnitTypeName = "Unit"

func registerLuaUnit(state *lua2.LState) {
	mt := state.NewTypeMetatable(luaUnitTypeName)
	state.SetGlobal("Unit", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaUnit))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsUnit(),
	))
}

func luaMethodsUnit() map[string]lua2.LGFunction {
	return luaMethodsExtend(map[string]lua2.LGFunction{
		"flags":                 luaGenericGetSetNumber[IUnit, byte](func(v IUnit) *byte { return &v.GetUnit().UnitFlags }),
		"level":                 luaGenericGetSetNumber[IUnit, byte](func(v IUnit) *byte { return &v.GetUnit().Level }),
		"hp":                    luaGenericGetSetNumber[IUnit, int](func(v IUnit) *int { return &v.GetUnit().HP }),
		"mp":                    luaGenericGetSetNumber[IUnit, int](func(v IUnit) *int { return &v.GetUnit().MP }),
		"unitUnk10Case":         luaGenericGetSetNumber[IUnit, byte](func(v IUnit) *byte { return &v.GetUnit().Unk10Case }),
		"unitUnk20CaseEntityID": luaGenericGetSetNumber[IUnit, uint16](func(v IUnit) *uint16 { return &v.GetUnit().Unk20CaseEntityID }),
		"unitUnk40Case0":        luaGenericGetSetNumber[IUnit, uint16](func(v IUnit) *uint16 { return &v.GetUnit().Unk40Case0 }),
		"unitUnk40Case1":        luaGenericGetSetNumber[IUnit, uint16](func(v IUnit) *uint16 { return &v.GetUnit().Unk40Case1 }),
		"unitUnk40Case2":        luaGenericGetSetNumber[IUnit, uint16](func(v IUnit) *uint16 { return &v.GetUnit().Unk40Case2 }),
		"unitUnk40Case3":        luaGenericGetSetNumber[IUnit, byte](func(v IUnit) *byte { return &v.GetUnit().Unk40Case3 }),
		"unitUnk80Case":         luaGenericGetSetNumber[IUnit, byte](func(v IUnit) *byte { return &v.GetUnit().Unk80Case }),
	}, luaMethodsWorldEntity)
}

func newLuaUnit(l *lua2.LState) int {
	unit := NewUnit(l.CheckString(1))

	ud := l.NewUserData()
	ud.Value = unit

	l.SetMetatable(ud, l.GetTypeMetatable(luaUnitTypeName))
	l.Push(ud)
	return 1
}
