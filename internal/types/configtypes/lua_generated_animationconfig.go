// Code generated by scripts/generatelua DO NOT EDIT.
package configtypes

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/pkg/datatypes"
	lua2 "github.com/yuin/gopher-lua"
)

type IAnimationConfig interface {
	GetAnimationConfig() *AnimationConfig
}

func (a *AnimationConfig) GetAnimationConfig() *AnimationConfig {
	return a
}

func registerLuaAnimationConfig(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("AnimationConfig")
	state.SetGlobal("AnimationConfig", mt)
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsAnimationConfig(),
	))
}

func luaMethodsAnimationConfig() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"id":               lua.LuaGenericGetSetNumber[IAnimationConfig](func(v IAnimationConfig) *int { return &v.GetAnimationConfig().ID }),
		"animationID":      lua.LuaGenericGetSetNumber[IAnimationConfig](func(v IAnimationConfig) *int { return &v.GetAnimationConfig().AnimationID }),
		"numFrames":        lua.LuaGenericGetSetNumber[IAnimationConfig](func(v IAnimationConfig) *int { return &v.GetAnimationConfig().NumFrames }),
		"triggerTime":      lua.LuaGenericGetSetNumber[IAnimationConfig](func(v IAnimationConfig) *int { return &v.GetAnimationConfig().TriggerTime }),
		"soundTriggerTime": lua.LuaGenericGetSetNumber[IAnimationConfig](func(v IAnimationConfig) *int { return &v.GetAnimationConfig().SoundTriggerTime }),
		"sourceNode":       lua.LuaGenericGetSetString[IAnimationConfig](func(v IAnimationConfig) *string { return &v.GetAnimationConfig().SourceNode }),
		"startFrame":       lua.LuaGenericGetSetNumber[IAnimationConfig](func(v IAnimationConfig) *int { return &v.GetAnimationConfig().StartFrame }),
		"soundID":          lua.LuaGenericGetSetNumber[IAnimationConfig](func(v IAnimationConfig) *int { return &v.GetAnimationConfig().SoundID }),
		"sourceOffset":     lua.LuaGenericGetSetValueAny[IAnimationConfig](func(v IAnimationConfig) *datatypes.Vector3Float32 { return &v.GetAnimationConfig().SourceOffset }),
		"looping":          lua.LuaGenericGetSetBool[IAnimationConfig](func(v IAnimationConfig) *bool { return &v.GetAnimationConfig().Looping }),

		"getAnimationConfig": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IAnimationConfig](l, 1)
			obj := objInterface.GetAnimationConfig()
			res0 := obj.GetAnimationConfig()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}

func (a *AnimationConfig) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("AnimationConfig"))
	return ud
}
