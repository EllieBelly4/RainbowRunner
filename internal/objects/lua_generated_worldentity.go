// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	lua2 "github.com/yuin/gopher-lua"
)

type IWorldEntity interface {
	GetWorldEntity() *WorldEntity
}

func (w *WorldEntity) GetWorldEntity() *WorldEntity {
	return w
}

func registerLuaWorldEntity(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("WorldEntity")
	state.SetGlobal("WorldEntity", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaWorldEntity))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsWorldEntity(),
	))
}

func luaMethodsWorldEntity() map[string]lua2.LGFunction {
	return luaMethodsExtend(map[string]lua2.LGFunction{
		"worldPosition":           luaGenericGetSetValue[IWorldEntity, datatypes.Vector3Float32](func(v IWorldEntity) *datatypes.Vector3Float32 { return &v.GetWorldEntity().WorldPosition }),
		"rotation":                luaGenericGetSetNumber[IWorldEntity](func(v IWorldEntity) *float32 { return &v.GetWorldEntity().Rotation }),
		"worldEntityFlags":        luaGenericGetSetNumber[IWorldEntity](func(v IWorldEntity) *uint32 { return &v.GetWorldEntity().WorldEntityFlags }),
		"worldEntityInitFlags":    luaGenericGetSetNumber[IWorldEntity](func(v IWorldEntity) *byte { return &v.GetWorldEntity().WorldEntityInitFlags }),
		"label":                   luaGenericGetSetString[IWorldEntity](func(v IWorldEntity) *string { return &v.GetWorldEntity().Label }),
		"unk1Case":                luaGenericGetSetNumber[IWorldEntity](func(v IWorldEntity) *uint16 { return &v.GetWorldEntity().Unk1Case }),
		"unk2Case":                luaGenericGetSetNumber[IWorldEntity](func(v IWorldEntity) *byte { return &v.GetWorldEntity().Unk2Case }),
		"unk4Case":                luaGenericGetSetNumber[IWorldEntity](func(v IWorldEntity) *uint32 { return &v.GetWorldEntity().Unk4Case }),
		"useCustomAnimationSpeed": luaGenericGetSetBool[IWorldEntity](func(v IWorldEntity) *bool { return &v.GetWorldEntity().UseCustomAnimationSpeed }),
		"animationSpeed":          luaGenericGetSetNumber[IWorldEntity](func(v IWorldEntity) *float32 { return &v.GetWorldEntity().AnimationSpeed }),
		"activate": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IWorldEntity](l, 1)
			obj := objInterface.GetWorldEntity()
			obj.Activate(
				lua.CheckReferenceValue[RRPlayer](l, 2),
				lua.CheckReferenceValue[UnitBehavior](l, 3), byte(l.CheckNumber(4)), byte(l.CheckNumber(5)),
			)

			return 0
		},
		"setPosition": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IWorldEntity](l, 1)
			obj := objInterface.GetWorldEntity()
			obj.SetPosition(
				lua.CheckValue[datatypes.Vector3Float32](l, 2),
			)

			return 0
		},
		"setRotation": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IWorldEntity](l, 1)
			obj := objInterface.GetWorldEntity()
			obj.SetRotation(float32(l.CheckNumber(2)))

			return 0
		},
		"type": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IWorldEntity](l, 1)
			obj := objInterface.GetWorldEntity()
			res0 := obj.Type()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("DRObjectType"))
			l.Push(ud)

			return 1
		},
		"writeInit": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IWorldEntity](l, 1)
			obj := objInterface.GetWorldEntity()
			obj.WriteInit(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
		"getWorldEntity": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IWorldEntity](l, 1)
			obj := objInterface.GetWorldEntity()
			res0 := obj.GetWorldEntity()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	}, luaMethodsEntity)
}
func newLuaWorldEntity(l *lua2.LState) int {
	obj := NewWorldEntity(string(l.CheckString(1)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("WorldEntity"))
	l.Push(ud)
	return 1
}

func (w *WorldEntity) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = w

	l.SetMetatable(ud, l.GetTypeMetatable("WorldEntity"))
	return ud
}
