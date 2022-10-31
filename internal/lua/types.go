package lua

import lua "github.com/yuin/gopher-lua"

func CheckReferenceValue[T any](state *lua.LState, index int) *T {
	ud := state.CheckUserData(index)
	if v, ok := ud.Value.(*T); ok {
		return v
	}
	state.ArgError(index, "zone expected")
	return nil
}
