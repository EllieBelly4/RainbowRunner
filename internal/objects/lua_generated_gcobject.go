// Code generated by scripts/generatelua DO NOT EDIT.
package objects

import (
	"RainbowRunner/internal/connections"
	lua "RainbowRunner/internal/lua"
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
	lua2 "github.com/yuin/gopher-lua"
)

type IGCObject interface {
	GetGCObject() *GCObject
}

func (g *GCObject) GetGCObject() *GCObject {
	return g
}

func registerLuaGCObject(state *lua2.LState) {
	// Ensure the import is referenced in code
	_ = lua.LuaScript{}

	mt := state.NewTypeMetatable("GCObject")
	state.SetGlobal("GCObject", mt)
	state.SetField(mt, "new", state.NewFunction(newLuaGCObject))
	state.SetField(mt, "__index", state.SetFuncs(state.NewTable(),
		luaMethodsGCObject(),
	))
}

func luaMethodsGCObject() map[string]lua2.LGFunction {
	return lua.LuaMethodsExtend(map[string]lua2.LGFunction{
		"entityProperties": lua.LuaGenericGetSetValueAny[IGCObject](func(v IGCObject) *RREntityProperties { return &v.GetGCObject().EntityProperties }),
		"version":          lua.LuaGenericGetSetNumber[IGCObject](func(v IGCObject) *uint8 { return &v.GetGCObject().Version }),
		"gcnativeType":     lua.LuaGenericGetSetString[IGCObject](func(v IGCObject) *string { return &v.GetGCObject().GCNativeType }),
		"gclabel":          lua.LuaGenericGetSetString[IGCObject](func(v IGCObject) *string { return &v.GetGCObject().GCLabel }),
		"gcchildren":       lua.LuaGenericGetSetValueAny[IGCObject](func(v IGCObject) *[]drobjecttypes.DRObject { return &v.GetGCObject().GCChildren }),
		"gctype":           lua.LuaGenericGetSetString[IGCObject](func(v IGCObject) *string { return &v.GetGCObject().GCType }),
		"properties":       lua.LuaGenericGetSetValueAny[IGCObject](func(v IGCObject) *[]GCObjectProperty { return &v.GetGCObject().Properties }),
		"entityHandler":    lua.LuaGenericGetSetValueAny[IGCObject](func(v IGCObject) *EntityMessageHandler { return &v.GetGCObject().EntityHandler }),
		"gcparent":         lua.LuaGenericGetSetValueAny[IGCObject](func(v IGCObject) *drobjecttypes.DRObject { return &v.GetGCObject().GCParent }),

		"removeChild": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.RemoveChild(
				lua.CheckValue[drobjecttypes.DRObject](l, 2),
			)
			l.Push(lua2.LBool(res0))

			return 1
		},

		"getGCType": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.GetGCType()
			l.Push(lua2.LString(res0))

			return 1
		},

		"getGCNativeType": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.GetGCNativeType()
			l.Push(lua2.LString(res0))

			return 1
		},

		"getChildrenFiltered": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.GetChildrenFiltered(
				lua.CheckValue[func(drobjecttypes.DRObject) bool](l, 2),
			)
			res0Array := l.NewTable()

			for _, res0 := range res0 {
				if res0 != nil {
					res0Array.Append(res0.ToLua(l))
				} else {
					res0Array.Append(lua2.LNil)
				}
			}

			l.Push(res0Array)

			return 1
		},

		"getPlayerOwner": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.GetPlayerOwner()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},

		"getChildrenByGCNativeType": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.GetChildrenByGCNativeType(string(l.CheckString(2)))
			res0Array := l.NewTable()

			for _, res0 := range res0 {
				if res0 != nil {
					res0Array.Append(res0.ToLua(l))
				} else {
					res0Array.Append(lua2.LNil)
				}
			}

			l.Push(res0Array)

			return 1
		},

		"getRREntityProperties": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.GetRREntityProperties()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},

		"getParentEntity": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.GetParentEntity()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},

		"setParent": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			obj.SetParent(
				lua.CheckValue[drobjecttypes.DRObject](l, 2),
			)

			return 0
		},

		"gcobject": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.GCObject()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},

		"string": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.String()
			l.Push(lua2.LString(res0))

			return 1
		},

		"setOwner": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			obj.SetOwner(
				lua.CheckReferenceValue[connections.RRConn](l, 2),
			)

			return 0
		},

		"removeChildAt": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			obj.RemoveChildAt(int(l.CheckNumber(2)))

			return 0
		},

		"type": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.Type()
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("drobjecttypes.DRObjectType"))
			l.Push(ud)

			return 1
		},

		"id": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.ID()
			l.Push(lua2.LNumber(res0))

			return 1
		},

		"readUpdate": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.ReadUpdate(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)
			ud := l.NewUserData()
			ud.Value = res0
			l.SetMetatable(ud, l.GetTypeMetatable("error"))
			l.Push(ud)

			return 1
		},

		"writeSynch": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			obj.WriteSynch(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"tick": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			obj.Tick()

			return 0
		},

		"init": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			obj.Init()

			return 0
		},

		"writeData": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			obj.WriteData(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"writeInit": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			obj.WriteInit(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"writeUpdate": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			obj.WriteUpdate(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"ownerID": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.OwnerID()
			l.Push(lua2.LNumber(res0))

			return 1
		},

		"children": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.Children()
			res0Array := l.NewTable()

			for _, res0 := range res0 {
				if res0 != nil {
					res0Array.Append(res0.ToLua(l))
				} else {
					res0Array.Append(lua2.LNil)
				}
			}

			l.Push(res0Array)

			return 1
		},

		"rrentityProperties": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.RREntityProperties()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},

		"writeFullGCObject": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			obj.WriteFullGCObject(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"addChild": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			obj.AddChild(
				lua.CheckValue[drobjecttypes.DRObject](l, 2),
			)

			return 0
		},

		"getChildByGCNativeType": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.GetChildByGCNativeType(string(l.CheckString(2)))
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},

		"getChildByGCType": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.GetChildByGCType(string(l.CheckString(2)))
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},

		"setVersion": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			obj.SetVersion(uint8(l.CheckNumber(2)))

			return 0
		},

		"readData": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			obj.ReadData(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)

			return 0
		},

		"walkChildren": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			obj.WalkChildren(
				lua.CheckValue[func(object drobjecttypes.DRObject)](l, 2),
			)

			return 0
		},

		"removeChildrenByGCNativeType": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.RemoveChildrenByGCNativeType(string(l.CheckString(2)))
			l.Push(lua2.LNumber(res0))

			return 1
		},

		"getChildFromGCTypeRequest": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.GetChildFromGCTypeRequest(
				lua.CheckReferenceValue[byter.Byter](l, 2),
			)
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},

		"getGCObject": func(l *lua2.LState) int {
			objInterface := lua.CheckInterfaceValue[IGCObject](l, 1)
			obj := objInterface.GetGCObject()
			res0 := obj.GetGCObject()
			if res0 != nil {
				l.Push(res0.ToLua(l))
			} else {
				l.Push(lua2.LNil)
			}

			return 1
		},
	})
}
func newLuaGCObject(l *lua2.LState) int {
	obj := NewGCObject(string(l.CheckString(1)))
	ud := l.NewUserData()
	ud.Value = obj

	l.SetMetatable(ud, l.GetTypeMetatable("GCObject"))
	l.Push(ud)
	return 1
}

func (g *GCObject) ToLua(l *lua2.LState) lua2.LValue {
	ud := l.NewUserData()
	ud.Value = g

	l.SetMetatable(ud, l.GetTypeMetatable("GCObject"))
	return ud
}
