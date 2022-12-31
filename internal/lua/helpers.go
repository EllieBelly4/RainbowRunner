package lua

import lua2 "github.com/yuin/gopher-lua"

func ValueToLValue(state *lua2.LState, value interface{}) lua2.LValue {
	switch v := value.(type) {
	case ILuaConvertible:
		return v.ToLua(state)
	case lua2.LValue:
		return v
	case int:
		return lua2.LNumber(v)
	case int32:
		return lua2.LNumber(v)
	case uint32:
		return lua2.LNumber(v)
	case uint:
		return lua2.LNumber(v)
	case uint64:
		return lua2.LNumber(v)
	case int64:
		return lua2.LNumber(v)
	case float64:
		return lua2.LNumber(v)
	case float32:
		return lua2.LNumber(v)
	case string:
		return lua2.LString(v)
	case bool:
		return lua2.LBool(v)
	case nil:
		return lua2.LNil
	case []interface{}:
		table := state.NewTable()
		for i, v := range v {
			table.RawSetInt(i+1, ValueToLValue(state, v))
		}
		return table
	case map[interface{}]interface{}:
		table := state.NewTable()
		for k, v := range v {
			table.RawSet(ValueToLValue(state, k), ValueToLValue(state, v))
		}
		return table
	default:
		panic("Unsupported type")
	}
}
