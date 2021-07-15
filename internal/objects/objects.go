package objects

import (
	"RainbowRunner/pkg/byter"
	"encoding/json"
	"os"
)

var ArmourMap map[string][]string

type DRObject interface {
	RREntityProperties() *RREntityProperties

	WriteFullGCObject(b *byter.Byter)
	WriteInit(b *byter.Byter)
	WriteUpdate(b *byter.Byter)
	WriteSynch(b *byter.Byter)

	ReadUpdate(reader *byter.Byter) error

	AddChild(object DRObject)
	Children() []DRObject
	GetChildByGCType(s string) DRObject
	GetChildByGCNativeType(s string) DRObject

	GetGCObject() *GCObject
	Tick()
}

func Init() {
	file, err := os.Open("resources/Dumps/armor_dump.json")

	if err != nil {
		panic(err)
	}

	j := json.NewDecoder(file)

	ArmourMap = make(map[string][]string)

	err = j.Decode(&ArmourMap)

	if err != nil {
		panic(err)
	}
}
