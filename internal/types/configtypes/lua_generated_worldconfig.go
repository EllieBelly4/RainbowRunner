// Code generated by scripts/generatelua DO NOT EDIT.
package configtypes

import (
	lua "RainbowRunner/internal/lua"
	lua2 "github.com/yuin/gopher-lua"
)

type IWorldConfig interface {
	GetWorldConfig() *WorldConfig
}

func (w *WorldConfig) GetWorldConfig() *WorldConfig {
	return w
}

func registerLuaWorldConfig(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("WorldConfig")
	state.SetGlobal("WorldConfig", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaWorldConfig))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsWorldConfig(),
	))
}

func luaMethodsWorldConfig() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"name":                     lua.LuaGenericGetSetString[IWorldConfig](func(v IWorldConfig) *string { return &v.GetWorldConfig().Name }),
		"encounterTable":           lua.LuaGenericGetSetString[IWorldConfig](func(v IWorldConfig) *string { return &v.GetWorldConfig().EncounterTable }),
		"generated":                lua.LuaGenericGetSetBool[IWorldConfig](func(v IWorldConfig) *bool { return &v.GetWorldConfig().Generated }),
		"mazeDeadEndRemovalChance": lua.LuaGenericGetSetNumber[IWorldConfig](func(v IWorldConfig) *int { return &v.GetWorldConfig().MazeDeadEndRemovalChance }),
		"mazeHeight":               lua.LuaGenericGetSetNumber[IWorldConfig](func(v IWorldConfig) *int { return &v.GetWorldConfig().MazeHeight }),
		"mazeRandomness":           lua.LuaGenericGetSetNumber[IWorldConfig](func(v IWorldConfig) *int { return &v.GetWorldConfig().MazeRandomness }),
		"mazeSparseness":           lua.LuaGenericGetSetNumber[IWorldConfig](func(v IWorldConfig) *int { return &v.GetWorldConfig().MazeSparseness }),
		"mazeWidth":                lua.LuaGenericGetSetNumber[IWorldConfig](func(v IWorldConfig) *int { return &v.GetWorldConfig().MazeWidth }),
		"tileSet":                  lua.LuaGenericGetSetString[IWorldConfig](func(v IWorldConfig) *string { return &v.GetWorldConfig().TileSet }),
		"tileSize":                 lua.LuaGenericGetSetNumber[IWorldConfig](func(v IWorldConfig) *int { return &v.GetWorldConfig().TileSize }),
		"worldEntityTable":         lua.LuaGenericGetSetString[IWorldConfig](func(v IWorldConfig) *string { return &v.GetWorldConfig().WorldEntityTable }),
		"worldEntityTable2":        lua.LuaGenericGetSetString[IWorldConfig](func(v IWorldConfig) *string { return &v.GetWorldConfig().WorldEntityTable2 }),
		"worldEntityTable3":        lua.LuaGenericGetSetString[IWorldConfig](func(v IWorldConfig) *string { return &v.GetWorldConfig().WorldEntityTable3 }),
		"entities":                 lua.LuaGenericGetSetValueAny[IWorldConfig](func(v IWorldConfig) *[]IEntityConfig { return &v.GetWorldConfig().Entities }),

		"getWorldConfig": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IWorldConfig](l, 1)
			obj := objInterface.GetWorldConfig()
			res0 := obj.GetWorldConfig()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaWorldConfig(l *lua2.LState) int {
	obj := NewWorldConfig()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("WorldConfig"))
	l.Push(ud)
	return 1
}

func (w *WorldConfig) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = w

	l.SetMetatable(ud, l.GetTypeMetatable("WorldConfig"))
	return ud
}
