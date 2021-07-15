package database

import (
	"RainbowRunner/internal/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type DRClassCollection []*DRClass

type EquipmentMap map[string]*DRClass

type DRWeapon struct {
	*DRClass
}

var Weapons DRClassCollection
var Armour DRClassCollection

var MeleeWeapons EquipmentMap
var RangedWeapons EquipmentMap
var Helmets EquipmentMap
var Armours EquipmentMap
var Gloves EquipmentMap
var Boots EquipmentMap

func (d *DRClassCollection) Find(fqGCType string) *DRClass {
	split := strings.Split(fqGCType, ".")

	for _, class := range []*DRClass(*d) {
		if class.Name == split[0] {
			if len(split) > 1 {
				return class.Find(split[1:])
			} else {
				return class
			}
		}
	}

	return nil
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
	for _, weapon := range Weapons {
		root := weapon.Name

		for _, subType := range weapon.Children {
			desc := subType.Find([]string{"Description"})

			// Mods do not have descriptions
			if desc == nil {
				continue
			}

			key := strings.Join([]string{root, subType.Name}, ".")

			if strings.HasSuffix(desc.Properties["WeaponClass"], "MELEE") {
				if MeleeWeapons == nil {
					MeleeWeapons = make(EquipmentMap)
				}

				MeleeWeapons[key] = subType
			} else {
				if RangedWeapons == nil {
					RangedWeapons = make(EquipmentMap)
				}

				RangedWeapons[key] = subType
			}
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
	for _, armour := range Armour {
		root := armour.Name

		for _, subType := range armour.Children {
			desc := subType.Find([]string{"Description"})

			// Mods do not have descriptions
			if desc == nil {
				continue
			}

			key := strings.Join([]string{root, subType.Name}, ".")

			if m, ok := armourTypeMap[subType.Slot()]; ok {
				if *m == nil {
					*m = make(EquipmentMap)
				}

				(*m)[key] = subType
			}
		}
	}
}
