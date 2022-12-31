package lua

import (
	log "github.com/sirupsen/logrus"
	lua2 "github.com/yuin/gopher-lua"
	"reflect"
)

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
	}

	rv := reflect.ValueOf(value)

	switch rv.Kind() {
	case reflect.Map:
		table := state.NewTable()

		for _, key := range rv.MapKeys() {
			table.RawSet(ValueToLValue(state, key.Interface()), ValueToLValue(state, rv.MapIndex(key).Interface()))
		}

		return table
	}

	log.Error("unsupported type: ", value)
	return lua2.LNil
}
