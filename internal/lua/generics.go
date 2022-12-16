package lua

import "github.com/yuin/gopher-lua"

func LuaGenericGetSetNumber[T any, K byte | uint16 | uint32 | int8 | int16 | int32 | int | uint | float32 | float64](
	valueCallback func(val T) *K,
) lua.LGFunction {
	return func(state *lua.LState) int {
		obj := CheckInterfaceValue[T](state, 1)
		val := valueCallback(obj)

		if state.GetTop() == 1 {
			state.Push(lua.LNumber(*val))
			return 1
		}

		*val = K(state.CheckNumber(2))
		return 0
	}
}

func LuaGenericGetSetString[T any](
	valueCallback func(val T) *string,
) lua.LGFunction {
	return func(state *lua.LState) int {
		obj := CheckInterfaceValue[T](state, 1)
		val := valueCallback(obj)

		if state.GetTop() == 1 {
			state.Push(lua.LString(*val))
			return 1
		}

		*val = state.CheckString(2)
		return 0
	}
}

func LuaGenericGetSetBool[T any](
	valueCallback func(val T) *bool,
) lua.LGFunction {
	return func(state *lua.LState) int {
		obj := CheckInterfaceValue[T](state, 1)
		val := valueCallback(obj)

		if state.GetTop() == 1 {
			state.Push(lua.LBool(*val))
			return 1
		}

		*val = state.CheckBool(2)
		return 0
	}
}

func LuaGenericGetSetValue[T any, K ILuaConvertible](
	valueCallback func(val T) *K,
) lua.LGFunction {
	return func(state *lua.LState) int {
		obj := CheckInterfaceValue[T](state, 1)
		val := valueCallback(obj)

		if state.GetTop() == 1 {
			state.Push((*val).ToLua(state))
			return 1
		}

		*val = CheckValue[K](state, 2)
		return 0
	}
}
