package database

import (
	"RainbowRunner/cmd/rrcli/configurator"
	"RainbowRunner/internal/gosucks"
	"RainbowRunner/internal/types"
	"RainbowRunner/internal/types/configtypes"
	"RainbowRunner/pkg/datatypes"
	"RainbowRunner/pkg/datatypes/drfloat"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"reflect"
	"strings"
)

type DRClassCollection map[string]*configtypes.DRClass

type EquipmentMap map[string]*configtypes.DRClass

type DRWeapon struct {
	*configtypes.DRClass
}

var config *configtypes.DRConfig
var checkpointConfigs map[string]map[string]*CheckpointConfig
var zones map[string]*configtypes.ZoneDefConfig
var worlds map[string]*configtypes.WorldConfig

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

		setPropertiesOnStruct(zoneDefConfig, props)

		zones[zoneID] = zoneDefConfig
	}

	gosucks.VAR(zonesConfig)

	return zones
}

func setPropertiesOnStruct(
	obj any,
	props configtypes.DRClassProperties,
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
			return
		}

		if field.Type() == reflect.TypeOf(datatypes.Vector3Float32{}) {
			field.Set(reflect.ValueOf(props.Vector3Val(propKey)))
			return
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
