package database

import (
	"RainbowRunner/internal/types/configtypes"
	"RainbowRunner/internal/types/drconfigtypes"
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

//go:generate go run ../../scripts/generatelua -type=ZoneConfig
type ZoneConfig struct {
	Name        string
	NPCs        map[string]*configtypes.NPCConfig
	Checkpoints map[string]*CheckpointConfig
	World       *configtypes.WorldConfig
	GCType      string
}

func (z *ZoneConfig) GetAllNPCs() []*configtypes.NPCConfig {
	l := make([]*configtypes.NPCConfig, 0)

	for _, npc := range z.NPCs {
		l = append(l, npc)
	}

	return l
}

func GetZoneConfig(name string) (*ZoneConfig, error) {
	var rawConfig []*drconfigtypes.DRClassChildGroup
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

	zoneConfig := NewZoneConfig(name, strings.Join(gcRoot, "."))

	worldConfig, ok := worlds[name]

	if !ok {
		log.Errorf("Zone does not have a world config: %s", name)
	} else {
		zoneConfig.World = worldConfig
	}

	if worldConfig.Entities != nil {
		addWorldEntities(zoneConfig, worldConfig, append(gcRoot, "npc"))
	}

	if rawConfig != nil {
		//configEntities := rawConfig[0].Entities[0].Children

		//if npcConfig, ok := configEntities["npc"]; ok {
		//	handleNPCs(zoneConfig, npcConfig, append(gcRoot, "npc"))
		//}

		if checkConfig, ok := checkpointConfigs[zoneConfig.Name]; ok {
			zoneConfig.Checkpoints = checkConfig
		}
	} else {
		log.Warnf("could not find configuration for zone %s, using defaults", name)
	}

	return zoneConfig, nil
}

func addWorldEntities(zoneConfig *ZoneConfig, worldConfig *configtypes.WorldConfig, gcroot []string) {
	if worldConfig.Entities == nil {
		return
	}

	gcrootString := strings.Join(gcroot, ".")

	if zoneConfig.NPCs == nil {
		zoneConfig.NPCs = make(map[string]*configtypes.NPCConfig)
	}

	for _, entity := range worldConfig.Entities {
		if entity.Type == configtypes.EntityConfigTypeNPC {
			shortGCType := strings.ToLower(strings.TrimPrefix(entity.FullGCType, gcrootString+"."))
			zoneConfig.NPCs[shortGCType] = configtypes.NewNPCConfigFromEntity(entity)
		}
	}
}

func NewZoneConfig(name string, gctype string) *ZoneConfig {
	return &ZoneConfig{
		Name:   name,
		GCType: gctype,
	}
}
