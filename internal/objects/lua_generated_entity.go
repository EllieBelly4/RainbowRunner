// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/internal/types/drobjecttypes"
	lua2 "github.com/yuin/gopher-lua"
)

type IEntity interface {
	GetEntity() *Entity
}

func (e *Entity) GetEntity() *Entity {
	return e
}

func registerLuaEntity(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("Entity")
	state.SetGlobal("Entity", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaEntity))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsEntity(),
	))
}

func luaMethodsEntity() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"name": lua.LuaGenericGetSetString[IEntity](func(v IEntity) *string { return &v.GetEntity().Name }),

		"getName": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IEntity](l, 1)
			obj := objInterface.GetEntity()
			res0 := obj.GetName()
			l.Push(lua2.LString(res0))

			return 1
		},

		"addChild": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IEntity](l, 1)
			obj := objInterface.GetEntity()
			obj.AddChild(
				lua.CheckValue[drobjecttypes.DRObject](l, 2),
			)

			return 0
		},

		"getEntity": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IEntity](l, 1)
			obj := objInterface.GetEntity()
			res0 := obj.GetEntity()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	}, luaMethodsGCObject)
}
func newLuaEntity(l *lua2.LState) int {
	obj := NewEntity(string(l.CheckString(1)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("Entity"))
	l.Push(ud)
	return 1
}

func (e *Entity) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = e

	l.SetMetatable(ud, l.GetTypeMetatable("Entity"))
	return ud
}
