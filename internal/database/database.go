package database

import (
	"RainbowRunner/cmd/rrcli/configurator"
	"RainbowRunner/internal/types"
	"RainbowRunner/internal/types/configtypes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

type DRClassCollection map[string]*configtypes.DRClass

type EquipmentMap map[string]*configtypes.DRClass

type DRWeapon struct {
	*configtypes.DRClass
}

var config *configtypes.DRConfig
var checkpointConfigs map[string]map[string]*CheckpointConfig

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
	log.Info("loading extracted config files")
	var err error

	config, err = configurator.LoadFromDumpedConfigFile("resources/Dumps/generated/finalconf.json")

	if err != nil {
		panic(err)
	}

	log.Info("loading checkpoint configs")

	rawCheckpointConfigs, err := config.Get("world.checkpoints")

	if err != nil {
		panic(err)
	}

	checkpointConfigs = sortCheckpoints(rawCheckpointConfigs)

	log.Info("config files loaded")
}

func LoadEquipmentFixtures() {
	log.Info("loading equipment fixtures")

	log.Info("loading armour fixtures")
	data, err := ioutil.ReadFile("resources/Dumps/generated/armour.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &Armour)

	if err != nil {
		panic(err)
	}

	AddArmours()

	log.Info("loading weapon fixtures")
	data, err = ioutil.ReadFile("resources/Dumps/generated/weapons.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &Weapons)

	if err != nil {
		panic(err)
	}

	AddWeapons()

	log.Info("equipment fixtures loaded")
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
