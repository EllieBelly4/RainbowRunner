package database

import (
	"RainbowRunner/internal/types/configtypes"
	log "github.com/sirupsen/logrus"
	"strings"
)

type CheckpointConfig struct {
	Name       string
	FullGCType string
	Order      int32
	SpawnPoint string
	Zone       string
	Entity     *CheckpointEntityConfig
}

type CheckpointEntityConfig struct {
	FullGCType string
	Label      string
	Blocking   bool
}

type ZoneConfig struct {
	Name        string
	NPCs        map[string]*NPCConfig
	MOBs        map[string]*MOBConfig //Added by Tim for loading in mobs
	Checkpoints map[string]*CheckpointConfig
}

func GetZoneConfig(name string) (*ZoneConfig, error) {
	var rawConfig []*configtypes.DRClassChildGroup
	var gcRoot []string

	paths := []string{
		strings.Join([]string{"world", name}, "."),
		name,
	}

	for i := 0; i < len(paths); i++ {
		newConfig, err := config.Get(paths[i])

		if err != nil {
			continue
		}

		rawConfig = newConfig
		gcRoot = strings.Split(paths[i], ".")
		break
	}

	zoneConfig := NewZoneConfig(name)

	if rawConfig != nil {
		configEntities := rawConfig[0].Entities[0].Children

		if npcConfig, ok := configEntities["npc"]; ok {
			handleNPCs(zoneConfig, npcConfig, append(gcRoot, "npc"))
		}
		//Added by Tim for loading in mobs
		if mobConfig, ok := configEntities["mob"]; ok {
			handleMOBs(zoneConfig, mobConfig, append(gcRoot, "mob"))
		}

		if checkConfig, ok := checkpointConfigs[zoneConfig.Name]; ok {
			zoneConfig.Checkpoints = checkConfig
		}
	} else {
		log.Warnf("could not find configuration for zone %s, using defaults", name)
	}

	return zoneConfig, nil
}

func handleNPCs(zoneConfig *ZoneConfig, rawNPCConfig *configtypes.DRClassChildGroup, gcRoot []string) {
	zoneConfig.NPCs = make(map[string]*NPCConfig)

	for npcKey, npcConfig := range rawNPCConfig.Entities[0].Children {
		npcGCroot := append(gcRoot, npcKey)
		npc := NewNPCConfig(npcConfig, npcGCroot)

		npc.FullGCType = strings.Join(npcGCroot, ".")

		zoneConfig.NPCs[npcKey] = npc
	}
}

// Added by Tim for loading in mobs
func handleMOBs(zoneConfig *ZoneConfig, rawMOBConfig *configtypes.DRClassChildGroup, gcRoot []string) {
	zoneConfig.MOBs = make(map[string]*MOBConfig)

	for mobKey, mobConfig := range rawMOBConfig.Entities[0].Children {
		mobGCroot := append(gcRoot, mobKey)
		mob := NewMOBConfig(mobConfig, mobGCroot)

		mob.FullGCType = strings.Join(mobGCroot, ".")

		zoneConfig.MOBs[mobKey] = mob
	}
}

func NewZoneConfig(name string) *ZoneConfig {
	return &ZoneConfig{
		Name: name,
	}
}
