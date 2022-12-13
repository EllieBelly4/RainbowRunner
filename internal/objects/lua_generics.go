package objects

import (
	"RainbowRunner/internal/lua"
	lua2 "github.com/yuin/gopher-lua"
)

func luaGenericGetSetNumber[T any, K byte | uint16 | uint32 | int8 | int16 | int32 | int | uint | float32 | float64](
	valueCallback func(val T) *K,
) lua2.LGFunction {
	return func(state *lua2.LState) int {
		obj := lua.CheckInterfaceValue[T](state, 1)
		val := valueCallback(obj)

		if state.GetTop() == 1 {
			state.Push(lua2.LNumber(*val))
			return 1
		}

		*val = K(state.CheckNumber(2))
		return 0
	}
}

func luaGenericGetSetString[T any](
	valueCallback func(val T) *string,
) lua2.LGFunction {
	return func(state *lua2.LState) int {
		obj := lua.CheckInterfaceValue[T](state, 1)
		val := valueCallback(obj)

		if state.GetTop() == 1 {
			state.Push(lua2.LString(*val))
			return 1
		}

		*val = state.CheckString(2)
		return 0
	}
}

func luaGenericGetSetBool[T any](
	valueCallback func(val T) *bool,
) lua2.LGFunction {
	return func(state *lua2.LState) int {
		obj := lua.CheckInterfaceValue[T](state, 1)
		val := valueCallback(obj)

		if state.GetTop() == 1 {
			state.Push(lua2.LBool(*val))
			return 1
		}

		*val = state.CheckBool(2)
		return 0
	}
}

func luaGenericGetSetValue[T any, K ILuaConvertible](
	valueCallback func(val T) *K,
) lua2.LGFunction {
	return func(state *lua2.LState) int {
		obj := lua.CheckInterfaceValue[T](state, 1)
		val := valueCallback(obj)

		if state.GetTop() == 1 {
			state.Push((*val).ToLua(state))
			return 1
		}

		*val = lua.CheckValue[K](state, 2)
		return 0
	}
}
