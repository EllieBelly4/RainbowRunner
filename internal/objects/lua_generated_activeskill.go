// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua "RainbowRunner/internal/lua"
	lua2 "github.com/yuin/gopher-lua"
)

type IActiveSkill interface {
	GetActiveSkill() *ActiveSkill
}

func (a *ActiveSkill) GetActiveSkill() *ActiveSkill {
	return a
}

func registerLuaActiveSkill(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActiveSkill")
	state.SetGlobal("ActiveSkill", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActiveSkill))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActiveSkill(),
	))
}

func luaMethodsActiveSkill() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{

		"getActiveSkill": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActiveSkill](l, 1)
			obj := objInterface.GetActiveSkill()
			res0 := obj.GetActiveSkill()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	}, luaMethodsSkill)
}
func newLuaActiveSkill(l *lua2.LState) int {
	obj := NewActiveSkill(string(l.CheckString(1)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActiveSkill"))
	l.Push(ud)
	return 1
}

func (a *ActiveSkill) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActiveSkill"))
	return ud
}
