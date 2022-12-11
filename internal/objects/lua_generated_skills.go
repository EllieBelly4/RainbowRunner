// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type ISkills interface {
	GetSkills() *Skills
}

func (s *Skills) GetSkills() *Skills {
	return s
}

func registerLuaSkills(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("Skills")
	state.SetGlobal("Skills", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaSkills))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsSkills(),
	))
}

func luaMethodsSkills() map[string]lua2.LGFunction {
	return luaMethodsExtend(map[string]lua2.LGFunction{
		"writeInit": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[ISkills](l, 1)
			obj := objInterface.GetSkills()
			obj.WriteInit(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"getSkills": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[ISkills](l, 1)
			obj := objInterface.GetSkills()
			res0 := obj.GetSkills()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("Skills"))
			l.Push(ud)

			return 1
		},
	}, luaMethodsComponent)
}
func newLuaSkills(l *lua2.LState) int {
	obj := NewSkills(string(l.CheckString(1)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("Skills"))
	l.Push(ud)
	return 1
}

func (s *Skills) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = s

	l.SetMetatable(ud, l.GetTypeMetatable("Skills"))
	return ud
}
