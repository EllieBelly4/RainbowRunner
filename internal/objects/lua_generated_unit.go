// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	lua2 "github.com/yuin/gopher-lua"
)

type IUnit interface {
	GetUnit() *Unit
}

func (u *Unit) GetUnit() *Unit {
	return u
}

func registerLuaUnit(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("Unit")
	state.SetGlobal("Unit", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaUnit))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsUnit(),
	))
}

func luaMethodsUnit() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"hp":                lua.LuaGenericGetSetNumber[IUnit](func(v IUnit) *int { return &v.GetUnit().HP }),
		"mp":                lua.LuaGenericGetSetNumber[IUnit](func(v IUnit) *int { return &v.GetUnit().MP }),
		"unitFlags":         lua.LuaGenericGetSetNumber[IUnit](func(v IUnit) *byte { return &v.GetUnit().UnitFlags }),
		"level":             lua.LuaGenericGetSetNumber[IUnit](func(v IUnit) *byte { return &v.GetUnit().Level }),
		"unk10Case":         lua.LuaGenericGetSetNumber[IUnit](func(v IUnit) *byte { return &v.GetUnit().Unk10Case }),
		"unk20CaseEntityID": lua.LuaGenericGetSetNumber[IUnit](func(v IUnit) *uint16 { return &v.GetUnit().Unk20CaseEntityID }),
		"unk40Case0":        lua.LuaGenericGetSetNumber[IUnit](func(v IUnit) *uint16 { return &v.GetUnit().Unk40Case0 }),
		"unk40Case1":        lua.LuaGenericGetSetNumber[IUnit](func(v IUnit) *uint16 { return &v.GetUnit().Unk40Case1 }),
		"unk40Case2":        lua.LuaGenericGetSetNumber[IUnit](func(v IUnit) *uint16 { return &v.GetUnit().Unk40Case2 }),
		"unk40Case3":        lua.LuaGenericGetSetNumber[IUnit](func(v IUnit) *byte { return &v.GetUnit().Unk40Case3 }),
		"unk80Case":         lua.LuaGenericGetSetNumber[IUnit](func(v IUnit) *byte { return &v.GetUnit().Unk80Case }),
		"addChild": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnit](l, 1)
			obj := objInterface.GetUnit()
			obj.AddChild(
				lua.CheckValue[drobjecttypes.DRObject](l, 2),
			)

			return 0
		},
		"writeInit": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnit](l, 1)
			obj := objInterface.GetUnit()
			obj.WriteInit(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"warp": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnit](l, 1)
			obj := objInterface.GetUnit()
			obj.Warp(
				lua.CheckValue[datatypes.Vector3Float32](l, 2),
			)

			return 0
		},
		"getUnit": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IUnit](l, 1)
			obj := objInterface.GetUnit()
			res0 := obj.GetUnit()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	}, luaMethodsWorldEntity)
}
func newLuaUnit(l *lua2.LState) int {
	obj := NewUnit(string(l.CheckString(1)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("Unit"))
	l.Push(ud)
	return 1
}

func (u *Unit) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = u

	l.SetMetatable(ud, l.GetTypeMetatable("Unit"))
	return ud
}
