// Code generated by scripts/generateluaregistrations DO NOT EDIT.
package configtypes

import lua2 "github.com/yuin/gopher-lua"

func RegisterAllLuaFunctions(state *lua2.LState) {
	registerLuaAnimationConfig(state)
	registerLuaEntityConfig(state)
	registerLuaNPCConfig(state)
	registerLuaWorldConfig(state)
	registerLuaZoneDefConfig(state)
}