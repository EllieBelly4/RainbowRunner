// Code generated by scripts/generatelua DO NOT EDIT.
package actions

import (
	lua "RainbowRunner/internal/lua"
	byter "RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/datatypes"
	lua2 "github.com/yuin/gopher-lua"
)

type IActionUsePosition interface {
	GetActionUsePosition() *ActionUsePosition
}

func (a *ActionUsePosition) GetActionUsePosition() *ActionUsePosition {
	return a
}

func registerLuaActionUsePosition(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("ActionUsePosition")
	state.SetGlobal("ActionUsePosition", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaActionUsePosition))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsActionUsePosition(),
	))
}

func luaMethodsActionUsePosition() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{

		"actionID": lua.LuaGenericGetSetValueAny[IActionUsePosition](func(v IActionUsePosition) *byte { return &v.GetActionUsePosition().ActionID }),

		"position": lua.LuaGenericGetSetValueAny[IActionUsePosition](func(v IActionUsePosition) *datatypes.Vector3Float32 { return &v.GetActionUsePosition().Position }),

		"opCode": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionUsePosition](l, 1)
			obj := objInterface.GetActionUsePosition()
			res0 := obj.OpCode()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("BehaviourAction"))
			l.Push(ud)

			return 1
		},

		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionUsePosition](l, 1)
			obj := objInterface.GetActionUsePosition()
			obj.Init(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"getActionUsePosition": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IActionUsePosition](l, 1)
			obj := objInterface.GetActionUsePosition()
			res0 := obj.GetActionUsePosition()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaActionUsePosition(l *lua2.LState) int {
	obj := NewActionUsePosition()
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("ActionUsePosition"))
	l.Push(ud)
	return 1
}

func (a *ActionUsePosition) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = a

	l.SetMetatable(ud, l.GetTypeMetatable("ActionUsePosition"))
	return ud
}
