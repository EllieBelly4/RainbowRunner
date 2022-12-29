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
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{})
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
