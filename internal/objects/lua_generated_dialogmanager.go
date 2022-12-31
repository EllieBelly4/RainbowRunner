// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IDialogManager interface {
	GetDialogManager() *DialogManager
}

func (d *DialogManager) GetDialogManager() *DialogManager {
	return d
}

func registerLuaDialogManager(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("DialogManager")
	state.SetGlobal("DialogManager", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaDialogManager))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsDialogManager(),
	))
}

func luaMethodsDialogManager() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{

		"type": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IDialogManager](l, 1)
			obj := objInterface.GetDialogManager()
			res0 := obj.Type()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("drobjecttypes.DRObjectType"))
			l.Push(ud)

			return 1
		},

		"writeInit": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IDialogManager](l, 1)
			obj := objInterface.GetDialogManager()
			obj.WriteInit(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"writeUpdate": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IDialogManager](l, 1)
			obj := objInterface.GetDialogManager()
			obj.WriteUpdate(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},
	}, luaMethodsGCObject)
}
func newLuaDialogManager(l *lua2.LState) int {
	obj := NewDialogManager()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("DialogManager"))
	l.Push(ud)
	return 1
}

func (d *DialogManager) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = d

	l.SetMetatable(ud, l.GetTypeMetatable("DialogManager"))
	return ud
}
