package lua

import lua "github.com/yuin/gopher-lua"

func CheckReferenceValue[T any](state *lua.LState, index int) *T {
	ud := state.CheckUserData(index)
	if v, ok := ud.Value.(*T); ok {
		return v
	}
	state.ArgError(index, "reference type does not match expected")
	return nil
}

func CheckInterfaceValue[T any](state *lua.LState, index int) T {
	ud := state.CheckUserData(index)
	if v, ok := ud.Value.(T); ok {
		return v
	}
	state.ArgError(index, "interface implementation does not match expected")
	return *new(T)
}

func CheckValue[T any](state *lua.LState, index int) T {
	ud := state.CheckUserData(index)
	if v, ok := ud.Value.(T); ok {
		return v
	}
	state.ArgError(index, "value type does not match expected")
	return *new(T)
}
