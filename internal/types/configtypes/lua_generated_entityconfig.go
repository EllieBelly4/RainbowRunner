// Code generated by scripts/generatelua DO NOT EDIT.
package configtypes

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/internal/types/drconfigtypes"
	"RainbowRunner/pkg/datatypes"
	lua2 "github.com/yuin/gopher-lua"
)

type IEntityConfig interface {
	GetEntityConfig() *EntityConfig
}

func (e *EntityConfig) GetEntityConfig() *EntityConfig {
	return e
}

func registerLuaEntityConfig(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("EntityConfig")
	state.SetGlobal("EntityConfig", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaEntityConfig))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsEntityConfig(),
	))
}

func luaMethodsEntityConfig() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"name": lua.LuaGenericGetSetString[IEntityConfig](func(v IEntityConfig) *string { return &v.GetEntityConfig().Name }),
		// -------------------------------------------------------------------------------------------------------------
		// Unsupported field type HitPoints
		// -------------------------------------------------------------------------------------------------------------
		// -------------------------------------------------------------------------------------------------------------
		// Unsupported field type ManaPoints
		// -------------------------------------------------------------------------------------------------------------
		"position":         lua.LuaGenericGetSetValue[IEntityConfig, datatypes.Vector3Float32](func(v IEntityConfig) *datatypes.Vector3Float32 { return &v.GetEntityConfig().Position }),
		"heading":          lua.LuaGenericGetSetNumber[IEntityConfig](func(v IEntityConfig) *int { return &v.GetEntityConfig().Heading }),
		"width":            lua.LuaGenericGetSetNumber[IEntityConfig](func(v IEntityConfig) *int { return &v.GetEntityConfig().Width }),
		"zone":             lua.LuaGenericGetSetString[IEntityConfig](func(v IEntityConfig) *string { return &v.GetEntityConfig().Zone }),
		"encounterTable":   lua.LuaGenericGetSetString[IEntityConfig](func(v IEntityConfig) *string { return &v.GetEntityConfig().EncounterTable }),
		"height":           lua.LuaGenericGetSetNumber[IEntityConfig](func(v IEntityConfig) *int { return &v.GetEntityConfig().Height }),
		"spawnPoint":       lua.LuaGenericGetSetString[IEntityConfig](func(v IEntityConfig) *string { return &v.GetEntityConfig().SpawnPoint }),
		"sizeX":            lua.LuaGenericGetSetNumber[IEntityConfig](func(v IEntityConfig) *int { return &v.GetEntityConfig().SizeX }),
		"sizeY":            lua.LuaGenericGetSetNumber[IEntityConfig](func(v IEntityConfig) *int { return &v.GetEntityConfig().SizeY }),
		"sizeZ":            lua.LuaGenericGetSetNumber[IEntityConfig](func(v IEntityConfig) *int { return &v.GetEntityConfig().SizeZ }),
		"canBeActivated":   lua.LuaGenericGetSetBool[IEntityConfig](func(v IEntityConfig) *bool { return &v.GetEntityConfig().CanBeActivated }),
		"respawnWhenClear": lua.LuaGenericGetSetBool[IEntityConfig](func(v IEntityConfig) *bool { return &v.GetEntityConfig().RespawnWhenClear }),
		"blocking":         lua.LuaGenericGetSetBool[IEntityConfig](func(v IEntityConfig) *bool { return &v.GetEntityConfig().Blocking }),
		"tableSelector":    lua.LuaGenericGetSetNumber[IEntityConfig](func(v IEntityConfig) *int { return &v.GetEntityConfig().TableSelector }),
		"color":            lua.LuaGenericGetSetNumber[IEntityConfig](func(v IEntityConfig) *uint { return &v.GetEntityConfig().Color }),
		"zoneStart":        lua.LuaGenericGetSetBool[IEntityConfig](func(v IEntityConfig) *bool { return &v.GetEntityConfig().ZoneStart }),
		"level":            lua.LuaGenericGetSetNumber[IEntityConfig](func(v IEntityConfig) *int { return &v.GetEntityConfig().Level }),
		"autoRespawn":      lua.LuaGenericGetSetBool[IEntityConfig](func(v IEntityConfig) *bool { return &v.GetEntityConfig().AutoRespawn }),
		"worldEntityTable": lua.LuaGenericGetSetString[IEntityConfig](func(v IEntityConfig) *string { return &v.GetEntityConfig().WorldEntityTable }),
		"respawnRate":      lua.LuaGenericGetSetNumber[IEntityConfig](func(v IEntityConfig) *int { return &v.GetEntityConfig().RespawnRate }),
		// -------------------------------------------------------------------------------------------------------------
		// Unsupported field type Animations
		// -------------------------------------------------------------------------------------------------------------
		// -------------------------------------------------------------------------------------------------------------
		// Unsupported field type Type
		// -------------------------------------------------------------------------------------------------------------
		"fullGCType": lua.LuaGenericGetSetString[IEntityConfig](func(v IEntityConfig) *string { return &v.GetEntityConfig().FullGCType }),
		// -------------------------------------------------------------------------------------------------------------
		// Unsupported field type Desc
		// -------------------------------------------------------------------------------------------------------------
		// -------------------------------------------------------------------------------------------------------------
		// Unsupported field type Behaviour
		// -------------------------------------------------------------------------------------------------------------

		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IEntityConfig](l, 1)
			obj := objInterface.GetEntityConfig()
			obj.Init(
				lua.CheckReferenceValue[drconfigtypes.DRClass](l, 2),
			)

			return 0
		},

		"getEntityConfig": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IEntityConfig](l, 1)
			obj := objInterface.GetEntityConfig()
			res0 := obj.GetEntityConfig()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaEntityConfig(l *lua2.LState) int {
	obj := NewEntityConfig()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("EntityConfig"))
	l.Push(ud)
	return 1
}

func (e *EntityConfig) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = e

	l.SetMetatable(ud, l.GetTypeMetatable("EntityConfig"))
	return ud
}
