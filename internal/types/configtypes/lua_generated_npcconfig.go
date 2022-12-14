// Code generated by scripts/generatelua DO NOT EDIT.
package configtypes

import (
	lua "RainbowRunner/internal/lua"
	lua2 "github.com/yuin/gopher-lua"
)

type INPCConfig interface {
	GetNPCConfig() *NPCConfig
}

func (n *NPCConfig) GetNPCConfig() *NPCConfig {
	return n
}

func registerLuaNPCConfig(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("NPCConfig")
	state.SetGlobal("NPCConfig", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaNPCConfig))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsNPCConfig(),
	))
}

func luaMethodsNPCConfig() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{

		"getNPCConfig": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[INPCConfig](l, 1)
			obj := objInterface.GetNPCConfig()
			res0 := obj.GetNPCConfig()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	}, luaMethodsEntityConfig)
}
func newLuaNPCConfig(l *lua2.LState) int {
	obj := NewNPCConfig()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("NPCConfig"))
	l.Push(ud)
	return 1
}

func (n *NPCConfig) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = n

	l.SetMetatable(ud, l.GetTypeMetatable("NPCConfig"))
	return ud
}
