package configtypes

import (
	"RainbowRunner/internal/types/drconfigtypes"
	"RainbowRunner/pkg/datatypes"
	"RainbowRunner/pkg/datatypes/drfloat"
	"fmt"
	"reflect"
	"strings"
)

func SetPropertiesOnStruct(
	obj any,
	props drconfigtypes.DRClassProperties,
) {
	rval := reflect.ValueOf(obj)

	objName := reflect.TypeOf(obj).Elem().Name()

	for propKey, val := range props {
		field := rval.Elem().FieldByName(propKey)

		sField, _ := rval.Type().Elem().FieldByName(propKey)

		tag := getStructTag(sField)
		tagInfo := parseTag(tag)

		if !field.IsValid() {
			//panic(fmt.Sprintf("unhandled property %s", propKey))
			fmt.Printf("%s unhandled property %s = %s\n", objName, propKey, val)
			continue
		}

		if field.Type() == reflect.TypeOf(drfloat.DRFloat(0)) {
			field.Set(reflect.ValueOf(drfloat.FromFloat32(float32(props.FloatVal(propKey)))))
			continue
		}

		if field.Type() == reflect.TypeOf(datatypes.Vector3Float32{}) {
			field.Set(reflect.ValueOf(props.Vector3Val(propKey)))
			continue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(props.StringVal(propKey))
		case reflect.Int:
			if tagInfo.Parse == "hex" {
				field.SetInt(int64(props.HexVal(propKey)))
			} else {
				field.SetInt(int64(props.IntVal(propKey)))
			}
		case reflect.Uint:
			if tagInfo.Parse == "hex" {
				field.SetUint(uint64(props.HexVal(propKey)))
			} else {
				field.SetUint(uint64(props.IntVal(propKey)))
			}
		case reflect.Bool:
			field.SetBool(props.BoolVal(propKey))
		case reflect.Float64:
			field.SetFloat(props.FloatVal(propKey))
		case reflect.Float32:
			field.SetFloat(props.FloatVal(propKey))
		default:
			panic(fmt.Sprintf("%s unhandled property type %s %s", objName, propKey, field.Kind()))
		}
	}
}

type tagInfo struct {
	Parse string
}

func parseTag(tag string) tagInfo {
	split := strings.Split(tag, " ")

	for _, s := range split {
		if strings.HasPrefix(s, "parse:") {
			parseVal := strings.Trim(s[6:], "\"")
			return tagInfo{
				Parse: parseVal,
			}
		}
	}

	return tagInfo{}
}

func getStructTag(f reflect.StructField) string {
	return string(f.Tag)
}
