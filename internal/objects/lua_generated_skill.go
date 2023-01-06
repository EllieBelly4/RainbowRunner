// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type ISkill interface {
	GetSkill() *Skill
}

func (s *Skill) GetSkill() *Skill {
	return s
}

func registerLuaSkill(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("Skill")
	state.SetGlobal("Skill", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaSkill))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsSkill(),
	))
}

func luaMethodsSkill() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"level":        lua.LuaGenericGetSetNumber[ISkill](func(v ISkill) *byte { return &v.GetSkill().Level }),
		"originalSlot": lua.LuaGenericGetSetNumber[ISkill](func(v ISkill) *int { return &v.GetSkill().OriginalSlot }),
		"skillFlags":   lua.LuaGenericGetSetNumber[ISkill](func(v ISkill) *byte { return &v.GetSkill().SkillFlags }),

		"writeData": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[ISkill](l, 1)
			obj := objInterface.GetSkill()
			obj.WriteData(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"getSkill": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[ISkill](l, 1)
			obj := objInterface.GetSkill()
			res0 := obj.GetSkill()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	}, luaMethodsManipulator)
}
func newLuaSkill(l *lua2.LState) int {
	obj := NewSkill(string(l.CheckString(1)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("Skill"))
	l.Push(ud)
	return 1
}

func (s *Skill) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = s

	l.SetMetatable(ud, l.GetTypeMetatable("Skill"))
	return ud
}
