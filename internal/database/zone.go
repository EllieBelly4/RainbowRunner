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
	NPCs        map[string]configtypes.INPCConfig
	Waypoints   map[string]configtypes.IWaypointConfig
	Checkpoints map[string]configtypes.ICheckpointEntityConfig
	World       *configtypes.WorldConfig
	GCType      string
	Entities    map[string]configtypes.IEntityConfig
}

func (z *ZoneConfig) GetAllNPCs() []*configtypes.NPCConfig {
	l := make([]*configtypes.NPCConfig, 0)

	for _, entity := range z.NPCs {
		l = append(l, entity.GetNPCConfig())
	}

	return l
}

func (z *ZoneConfig) GetAllCheckpointEntities() []*configtypes.CheckpointEntityConfig {
	l := make([]*configtypes.CheckpointEntityConfig, 0)

	for _, entity := range z.Checkpoints {
		l = append(l, entity.GetCheckpointEntityConfig())
	}

	return l
}

func (z *ZoneConfig) GetAllWaypoints() []*configtypes.WaypointConfig {
	l := make([]*configtypes.WaypointConfig, 0)

	for _, entity := range z.Waypoints {
		l = append(l, entity.GetWaypointConfig())
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

		//if checkConfig, ok := checkpointConfigs[zoneConfig.Name]; ok {
		//	zoneConfig.Checkpoints = checkConfig
		//}
	} else {
		log.Warnf("could not find configuration for zone %s, using defaults", name)
	}

	return zoneConfig, nil
}

func addWorldEntities(zoneConfig *ZoneConfig, worldConfig *configtypes.WorldConfig, gcroot []string) {
	if worldConfig.Entities == nil {
		return
	}

	for _, entity := range worldConfig.Entities {
		shortGCType := strings.ToLower(entity.GetEntityConfig().Name)

		zoneConfig.Entities[shortGCType] = entity

		if npcConfig, ok := entity.(configtypes.INPCConfig); ok {
			zoneConfig.NPCs[shortGCType] = npcConfig.GetNPCConfig()
		}

		if checkpointConfig, ok := entity.(configtypes.ICheckpointEntityConfig); ok {
			zoneConfig.Checkpoints[shortGCType] = checkpointConfig.GetCheckpointEntityConfig()
		}

		if waypointConfig, ok := entity.(configtypes.IWaypointConfig); ok {
			zoneConfig.Waypoints[shortGCType] = waypointConfig.GetWaypointConfig()
		}
	}
}

func NewZoneConfig(name string, gctype string) *ZoneConfig {
	return &ZoneConfig{
		Name:        name,
		GCType:      gctype,
		Checkpoints: map[string]configtypes.ICheckpointEntityConfig{},
		NPCs:        map[string]configtypes.INPCConfig{},
		Waypoints:   map[string]configtypes.IWaypointConfig{},
		Entities:    map[string]configtypes.IEntityConfig{},
	}
}
