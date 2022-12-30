package drconfigtypes

import (
	"RainbowRunner/pkg/datatypes"
	"encoding/binary"
	"encoding/hex"
	"regexp"
	"strconv"
	"strings"
)

type DRClassProperties map[string]string

var quoteRemoveRegex = regexp.MustCompile("^[\"']|[\"']$")

func (p *DRClassProperties) StringVal(key string) string {
	return quoteRemoveRegex.ReplaceAllString((*p)[key], "")
}

func (p *DRClassProperties) IntVal(key string) int {
	val, err := strconv.Atoi((*p)[key])

	if err != nil {
		panic(err)
	}

	return val
}

func (p *DRClassProperties) BoolVal(key string) bool {
	val, err := strconv.ParseBool((*p)[key])

	if err != nil {
		panic(err)
	}

	return val
}

func (p *DRClassProperties) FloatVal(key string) float64 {
	val, err := strconv.ParseFloat((*p)[key], 64)

	if err != nil {
		panic(err)
	}

	return val
}

func (p *DRClassProperties) Vector3Val(key string) datatypes.Vector3Float32 {
	rawVal := (*p)[key]

	splitRawVal := strings.Split(rawVal, ",")

	if len(splitRawVal) != 3 {
		panic("invalid vector3")
	}

	x, err := strconv.ParseFloat(splitRawVal[0], 32)

	if err != nil {
		panic(err)
	}

	y, err := strconv.ParseFloat(splitRawVal[1], 32)

	if err != nil {
		panic(err)
	}

	z, err := strconv.ParseFloat(splitRawVal[2], 32)

	if err != nil {
		panic(err)
	}

	return datatypes.Vector3Float32{
		X: float32(x),
		Y: float32(y),
		Z: float32(z),
	}
}

func (p *DRClassProperties) HexVal(key string) uint32 {
	val, err := hex.DecodeString((*p)[key][2:])

	if err != nil {
		panic(err)
	}

	return binary.BigEndian.Uint32(val)
}
