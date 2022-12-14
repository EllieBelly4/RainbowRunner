// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua "RainbowRunner/internal/lua"
	lua2 "github.com/yuin/gopher-lua"
)

type IAnimationsList interface {
	GetAnimationsList() *AnimationsList
}

func (a *AnimationsList) GetAnimationsList() *AnimationsList {
	return a
}

func registerLuaAnimationsList(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("AnimationsList")
	state.SetGlobal("AnimationsList", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaAnimationsList))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsAnimationsList(),
	))
}

func luaMethodsAnimationsList() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"animations": lua.LuaGenericGetSetValueAny[IAnimationsList](func(v IAnimationsList) *[]*Animation { return &v.GetAnimationsList().Animations }),

		"getAnimationsList": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAnimationsList](l, 1)
			obj := objInterface.GetAnimationsList()
			res0 := obj.GetAnimationsList()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaAnimationsList(l *lua2.LState) int {
	obj := NewAnimationsList()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("AnimationsList"))
	l.Push(ud)
	return 1
}

func (a *AnimationsList) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("AnimationsList"))
	return ud
}
