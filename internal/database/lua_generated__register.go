// Code generated by scripts/generateluaregistrations DO NOT EDIT.
package database

import lua2 "github.com/yuin/gopher-lua"

func RegisterAllLuaFunctions(state *lua2.LState) {
	registerLuaAnimationConfig(state)
	registerLuaNPCConfig(state)
}
