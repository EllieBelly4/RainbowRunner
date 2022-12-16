package lua

import lua "github.com/yuin/gopher-lua"

type ILuaConvertible interface {
	ToLua(state *lua.LState) lua.LValue
}

func LuaMethodsExtend(child map[string]lua.LGFunction, parents ...func() map[string]lua.LGFunction) map[string]lua.LGFunction {
	newMethods := make(map[string]lua.LGFunction)

	for _, parent := range parents {
		for key, value := range parent() {
			newMethods[key] = value
		}
	}

	for key, value := range child {
		newMethods[key] = value
	}

	return newMethods
}
