package objects

import (
	"github.com/goccy/go-json"
	"os"
)

var ArmourMap map[string][]string

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
