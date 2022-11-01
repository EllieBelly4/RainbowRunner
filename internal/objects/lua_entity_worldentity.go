package objects

import (
	"RainbowRunner/internal/lua"
	lua2 "github.com/yuin/gopher-lua"
)

const luaWorldEntityTypeName = "WorldEntity"

func registerLuaWorldEntity(state *lua2.LState) {
	mt := state.NewTypeMetatable(luaWorldEntityTypeName)
	state.SetGlobal("WorldEntity", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaWorldEntity))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(), entityLuaWorldEntityMethods))
}

var entityLuaWorldEntityMethods = map[string]lua2.LGFunction{
	"label": entityLuaWorldEntityGetSetLabel,
}

func entityLuaWorldEntityGetSetLabel(state *lua2.LState) int {
	worldEntity := lua.CheckReferenceValue[WorldEntity](state, 1)

	if state.GetTop() == 2 {
		worldEntity.Label = state.CheckString(2)
		return 0
	}

	state.Push(lua2.LString(worldEntity.Label))
	return 1
}

func newLuaWorldEntity(l *lua2.LState) int {
	worldEntity := NewWorldEntity(l.CheckString(1))

	ud := l.NewUserData()
	ud.Value = worldEntity

	l.SetMetatable(ud, l.GetTypeMetatable(luaWorldEntityTypeName))
	l.Push(ud)
	return 1
}
