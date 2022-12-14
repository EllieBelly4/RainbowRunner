// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IMeleeWeapon interface {
	GetMeleeWeapon() *MeleeWeapon
}

func (m *MeleeWeapon) GetMeleeWeapon() *MeleeWeapon {
	return m
}

func registerLuaMeleeWeapon(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("MeleeWeapon")
	state.SetGlobal("MeleeWeapon", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaMeleeWeapon))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsMeleeWeapon(),
	))
}

func luaMethodsMeleeWeapon() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{

		"writeInit": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IMeleeWeapon](l, 1)
			obj := objInterface.GetMeleeWeapon()
			obj.WriteInit(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"getMeleeWeapon": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IMeleeWeapon](l, 1)
			obj := objInterface.GetMeleeWeapon()
			res0 := obj.GetMeleeWeapon()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	}, luaMethodsItem)
}
func newLuaMeleeWeapon(l *lua2.LState) int {
	obj := NewMeleeWeapon(string(l.CheckString(1)), string(l.CheckString(2)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("MeleeWeapon"))
	l.Push(ud)
	return 1
}

func (m *MeleeWeapon) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = m

	l.SetMetatable(ud, l.GetTypeMetatable("MeleeWeapon"))
	return ud
}
