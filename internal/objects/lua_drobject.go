package objects

import (
	"RainbowRunner/internal/lua"
	lua2 "github.com/yuin/gopher-lua"
)

const luaDrobjectTypeName = "DRObject"

func registerLuaDrobject(state *lua2.LState) {
	mt := state.NewTypeMetatable(luaDrobjectTypeName)
	state.SetGlobal("DRobject", mt)
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaDRObjectMethods,
	))
}

func luaDRObjectExtendMethods(methods map[string]lua2.LGFunction) map[string]lua2.LGFunction {
	newMethods := make(map[string]lua2.LGFunction)

	for key, value := range luaDRObjectMethods {
		newMethods[key] = value
	}

	for key, value := range methods {
		newMethods[key] = value
	}

	return newMethods
}

var luaDRObjectMethods = map[string]lua2.LGFunction{
	"addChild":               luaDRObjectAddChild,
	"children":               luaDRObjectGetChildren,
	"type":                   luaDRObjectGetType,
	"getChildByGCType":       luaDRObjectGetChildByGCType,
	"getChildByGCNativeType": luaDRObjectGetChildByGCNativeType,
	"ownerID":                luaDRObjectGetOwnerID,
}

func luaDRObjectGetOwnerID(state *lua2.LState) int {
	obj := lua.CheckInterfaceValue[DRObject](state, 1)

	state.Push(lua2.LNumber(obj.OwnerID()))
	return 1
}

func luaDRObjectGetChildByGCNativeType(state *lua2.LState) int {
	obj := lua.CheckInterfaceValue[DRObject](state, 1)
	gctype := state.CheckString(2)

	child := obj.GetChildByGCNativeType(gctype)

	ud := state.NewUserData()
	ud.Value = child

	state.SetMetatable(ud, state.GetTypeMetatable(luaDrobjectTypeName))

	state.Push(ud)
	return 1
}

func luaDRObjectGetChildByGCType(state *lua2.LState) int {
	obj := lua.CheckInterfaceValue[DRObject](state, 1)
	gctype := state.CheckString(2)

	child := obj.GetChildByGCType(gctype)

	ud := state.NewUserData()
	ud.Value = child

	state.SetMetatable(ud, state.GetTypeMetatable(luaDrobjectTypeName))

	state.Push(ud)
	return 1
}

func luaDRObjectGetType(state *lua2.LState) int {
	obj := lua.CheckInterfaceValue[DRObject](state, 1)

	state.Push(lua2.LString(obj.Type().String()))

	return 1
}

func luaDRObjectGetChildren(state *lua2.LState) int {
	obj := lua.CheckInterfaceValue[DRObject](state, 1)

	count := 0

	for _, child := range obj.Children() {
		ud := state.NewUserData()
		ud.Value = child

		state.SetMetatable(ud, state.GetTypeMetatable(luaDrobjectTypeName))
		state.Push(ud)
		count++
	}

	return count
}

func luaDRObjectAddChild(state *lua2.LState) int {
	obj := lua.CheckInterfaceValue[DRObject](state, 1)
	child := lua.CheckInterfaceValue[DRObject](state, 2)

	obj.AddChild(child)

	return 0
}
