package database

type ZoneConfig struct {
	Name string
}

func GetZoneConfig(name string) error {
	//rawConfig, err := config.Get("world." + name)
	//
	//if err != nil {
	//	return err
	//}
	//
	//zoneConfig := NewZoneConfig()
	//
	//configEntities := rawConfig[0].Entities[0].Children
	//
	//if npcConfig, ok := configEntities["npc"]; ok {
	//	handleNPCs(zoneConfig, npcConfig)
	//}
	//
	//gosucks.VAR(rawConfig)
	//
	////npcs := rawConfig

	return nil
}

//func handleNPCs(zoneConfig *ZoneConfig, npcConfig *configtypes.DRClassChildGroup) {
//	for _, npcConfig := range npcConfig.Entities[0].Children {
//		npc := NewNPCCOnfig(npcConfig)
//	}
//}

func NewZoneConfig() *ZoneConfig {
	return &ZoneConfig{}
}
