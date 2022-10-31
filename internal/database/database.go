package database

import (
	"RainbowRunner/cmd/rrcli/configurator"
	"RainbowRunner/internal/types"
	"RainbowRunner/internal/types/configtypes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type DRClassCollection map[string]*configtypes.DRClass

type EquipmentMap map[string]*configtypes.DRClass

type DRWeapon struct {
	*configtypes.DRClass
}

var Config *configtypes.DRConfig

var Weapons []*configtypes.DRClassChildGroup
var Armour []*configtypes.DRClassChildGroup

var MeleeWeapons EquipmentMap
var RangedWeapons EquipmentMap
var Helmets EquipmentMap
var Armours EquipmentMap
var Gloves EquipmentMap
var Boots EquipmentMap

func FindItem(db []*configtypes.DRClassChildGroup, gcType string) *configtypes.DRClass {
	gcType = strings.ToLower(gcType)
	for _, group := range db {
		if strings.ToLower(group.GCType) == gcType {
			return group.Entities[0]
		}
	}

	return nil
}

func LoadConfigFiles() {
	config, err := configurator.LoadFromDumpedConfigFile("resources/Dumps/generated/finalconf.json")

	if err != nil {
		panic(err)
	}

	Config = config
}

func LoadEquipmentFixtures() {
	fmt.Println("loading equipment fixtures")

	data, err := ioutil.ReadFile("resources/Dumps/generated/armour.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &Armour)

	if err != nil {
		panic(err)
	}

	AddArmours()

	data, err = ioutil.ReadFile("resources/Dumps/generated/weapons.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &Weapons)

	if err != nil {
		panic(err)
	}

	AddWeapons()

	fmt.Println("equipment fixtures loaded")
}

func AddWeapons() {
	for _, sub := range Weapons {
		subType := sub.Entities[0]
		desc := subType.Find([]string{"description"})

		// Mods do not have descriptions
		if desc == nil {
			continue
		}

		if strings.HasSuffix(desc.Properties["WeaponClass"], "MELEE") {
			if MeleeWeapons == nil {
				MeleeWeapons = make(EquipmentMap)
			}

			MeleeWeapons[sub.GCType] = subType
		} else {
			if RangedWeapons == nil {
				RangedWeapons = make(EquipmentMap)
			}

			RangedWeapons[sub.GCType] = subType
		}
	}
}

var armourTypeMap = map[types.EquipmentSlot]*EquipmentMap{
	types.EquipmentSlotHead:  &Helmets,
	types.EquipmentSlotTorso: &Armours,
	types.EquipmentSlotHand:  &Gloves,
	types.EquipmentSlotFoot:  &Boots,
	//types.EquipmentSlotAmulet:
	//types.EquipmentSlotLRing:
	//types.EquipmentSlotRRing:
	//types.EquipmentSlotShoulder:
	//types.EquipmentSlotNone2:
	//types.EquipmentSlotWeapon:
	//types.EquipmentSlotOffhand:
}

func AddArmours() {
	for _, sub := range Armour {
		subType := sub.Entities[0]
		desc := subType.Find([]string{"description"})

		// Mods do not have descriptions
		if desc == nil {
			continue
		}

		if m, ok := armourTypeMap[subType.Slot()]; ok {
			if *m == nil {
				*m = make(EquipmentMap)
			}

			(*m)[sub.GCType] = subType
		}
	}
}
