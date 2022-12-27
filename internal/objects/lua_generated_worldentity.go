// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/internal/script"
	"RainbowRunner/internal/types/drobjecttypes"
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
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"collisionRadius":         lua.LuaGenericGetSetNumber[IWorldEntity](func(v IWorldEntity) *int { return &v.GetWorldEntity().CollisionRadius }),
		"animationsList":          lua.LuaGenericGetSetValue[IWorldEntity, *AnimationsList](func(v IWorldEntity) **AnimationsList { return &v.GetWorldEntity().AnimationsList }),
		"worldPosition":           lua.LuaGenericGetSetValue[IWorldEntity, datatypes.Vector3Float32](func(v IWorldEntity) *datatypes.Vector3Float32 { return &v.GetWorldEntity().WorldPosition }),
		"rotation":                lua.LuaGenericGetSetNumber[IWorldEntity](func(v IWorldEntity) *float32 { return &v.GetWorldEntity().Rotation }),
		"worldEntityFlags":        lua.LuaGenericGetSetNumber[IWorldEntity](func(v IWorldEntity) *uint32 { return &v.GetWorldEntity().WorldEntityFlags }),
		"worldEntityInitFlags":    lua.LuaGenericGetSetNumber[IWorldEntity](func(v IWorldEntity) *byte { return &v.GetWorldEntity().WorldEntityInitFlags }),
		"label":                   lua.LuaGenericGetSetString[IWorldEntity](func(v IWorldEntity) *string { return &v.GetWorldEntity().Label }),
		"unk1Case":                lua.LuaGenericGetSetNumber[IWorldEntity](func(v IWorldEntity) *uint16 { return &v.GetWorldEntity().Unk1Case }),
		"unk2Case":                lua.LuaGenericGetSetNumber[IWorldEntity](func(v IWorldEntity) *byte { return &v.GetWorldEntity().Unk2Case }),
		"unk4Case":                lua.LuaGenericGetSetNumber[IWorldEntity](func(v IWorldEntity) *uint32 { return &v.GetWorldEntity().Unk4Case }),
		"useCustomAnimationSpeed": lua.LuaGenericGetSetBool[IWorldEntity](func(v IWorldEntity) *bool { return &v.GetWorldEntity().UseCustomAnimationSpeed }),
		"animationSpeed":          lua.LuaGenericGetSetNumber[IWorldEntity](func(v IWorldEntity) *float32 { return &v.GetWorldEntity().AnimationSpeed }),
		"addChild": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IWorldEntity](l, 1)
			obj := objInterface.GetWorldEntity()
			obj.AddChild(
				lua.CheckValue[drobjecttypes.DRObject](l, 2),
			)

			return 0
		},
		"tick": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IWorldEntity](l, 1)
			obj := objInterface.GetWorldEntity()
			obj.Tick()

			return 0
		},
		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IWorldEntity](l, 1)
			obj := objInterface.GetWorldEntity()
			obj.Init()

			return 0
		},
		"setScript": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IWorldEntity](l, 1)
			obj := objInterface.GetWorldEntity()
			obj.SetScript(
				lua.CheckValue[script.IEntityScript](l, 2),
			)

			return 0
		},
		"animations": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IWorldEntity](l, 1)
			obj := objInterface.GetWorldEntity()
			res0 := obj.Animations()
			res0Array := l.NewTable()

			for _, res0 := range res0 {
				if res0 != nil {
					res0Array.Append(res0.ToLua(l))
				} else {
					res0Array.Append(lua2.LNil)
				}
			}

			l.Push(res0Array)

			return 1
		},
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
			l.SetMetatable(ud, l.GetTypeMetatable("drobjecttypes.DRObjectType"))
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
