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
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsWorldEntity(),
	))
}

func luaMethodsWorldEntity() map[string]lua2.LGFunction {
	return luaMethodsExtend(map[string]lua2.LGFunction{
		"label":                entityLuaWorldEntityGetSetLabel,
		"worldEntityFlags":     luaGenericGetSetNumber[IWorldEntity, uint32](func(v IWorldEntity) *uint32 { return &v.GetWorldEntity().WorldEntityFlags }),
		"worldEntityInitFlags": luaGenericGetSetNumber[IWorldEntity, byte](func(v IWorldEntity) *byte { return &v.GetWorldEntity().WorldEntityInitFlags }),
		"worldEntityUnk1Case":  luaGenericGetSetNumber[IWorldEntity, uint16](func(v IWorldEntity) *uint16 { return &v.GetWorldEntity().Unk1Case }),
		"worldEntityUnk2Case":  luaGenericGetSetNumber[IWorldEntity, byte](func(v IWorldEntity) *byte { return &v.GetWorldEntity().Unk2Case }),
		"worldEntityUnk4Case":  luaGenericGetSetNumber[IWorldEntity, uint32](func(v IWorldEntity) *uint32 { return &v.GetWorldEntity().Unk4Case }),
		"worldEntityUnk8Case":  luaGenericGetSetNumber[IWorldEntity, uint32](func(v IWorldEntity) *uint32 { return &v.GetWorldEntity().Unk8Case }),
	}, luaMethodsDRObject)
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
