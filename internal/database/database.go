package database

import (
	"RainbowRunner/cmd/rrcli/configurator"
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/internal/types"
	"RainbowRunner/internal/types/configtypes"
	drconfigtypes2 "RainbowRunner/internal/types/drconfigtypes"
	"fmt"
	"github.com/goccy/go-json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
)

type DRClassCollection map[string]*drconfigtypes2.DRClass

type EquipmentMap map[string]*drconfigtypes2.DRClass

type DRWeapon struct {
	*drconfigtypes2.DRClass
}

var config *drconfigtypes2.DRConfig
var checkpointConfigs map[string]map[string]*CheckpointConfig
var zones map[string]*configtypes.ZoneDefConfig
var worlds map[string]*configtypes.WorldConfig

var Weapons []*drconfigtypes2.DRClassChildGroup
var Armour []*drconfigtypes2.DRClassChildGroup

var MeleeWeapons EquipmentMap
var RangedWeapons EquipmentMap
var Helmets EquipmentMap
var Armours EquipmentMap
var Gloves EquipmentMap
var Boots EquipmentMap

func FindItem(db []*drconfigtypes2.DRClassChildGroup, gcType string) *drconfigtypes2.DRClass {
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

	worlds = LoadWorldConfigs()
	zones = LoadZoneConfigs()

	log.Info("config files loaded")
}

func LoadZoneConfigs() map[string]*configtypes.ZoneDefConfig {
	log.Info("loading zone configs")

	zones := make(map[string]*configtypes.ZoneDefConfig)

	zonesConfig, err := configurator.LoadFromDumpedConfigFile("resources/Dumps/generated/zones.json")

	if err != nil {
		panic(err)
	}

	for zoneID, zoneGroup := range zonesConfig.Classes.Children {
		if len(zoneGroup.Entities) > 1 {
			panic(fmt.Sprintf("zone %s has more than one entity", zoneID))
		}

		zoneDef := zoneGroup.Entities[0]

		zoneDefConfig := configtypes.NewZoneDefConfig()
		props := zoneDef.Properties

		configtypes.SetPropertiesOnStruct(zoneDefConfig, props)

		zones[zoneID] = zoneDefConfig
	}

	gosucks.VAR(zonesConfig)

	return zones
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

		slot, err := subType.Slot()

		if err != nil {
			log.Error(err)
			continue
		}

		if m, ok := armourTypeMap[slot]; ok {
			if *m == nil {
				*m = make(EquipmentMap)
			}

			(*m)[sub.GCType] = subType
		}
	}
}
