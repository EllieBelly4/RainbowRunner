// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua "RainbowRunner/internal/lua"
	lua2 "github.com/yuin/gopher-lua"
)

type ISkillSlot interface {
	GetSkillSlot() *SkillSlot
}

func (s *SkillSlot) GetSkillSlot() *SkillSlot {
	return s
}

func registerLuaSkillSlot(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("SkillSlot")
	state.SetGlobal("SkillSlot", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaSkillSlot))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsSkillSlot(),
	))
}

func luaMethodsSkillSlot() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{

		"slotID": lua.LuaGenericGetSetValueAny[ISkillSlot](func(v ISkillSlot) *int { return &v.GetSkillSlot().SlotID }),

		"slotType": lua.LuaGenericGetSetValueAny[ISkillSlot](func(v ISkillSlot) *uint32 { return &v.GetSkillSlot().SlotType }),

		"getSkillSlot": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[ISkillSlot](l, 1)
			obj := objInterface.GetSkillSlot()
			res0 := obj.GetSkillSlot()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	}, luaMethodsComponent)
}
func newLuaSkillSlot(l *lua2.LState) int {
	obj := NewSkillSlot(string(l.CheckString(1)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("SkillSlot"))
	l.Push(ud)
	return 1
}

func (s *SkillSlot) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = s

	l.SetMetatable(ud, l.GetTypeMetatable("SkillSlot"))
	return ud
}
