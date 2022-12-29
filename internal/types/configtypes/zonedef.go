package configtypes

//go:generate go run ../../../scripts/generatelua -type=ZoneDefConfig
type ZoneDefConfig struct {
	Label                 string
	Name                  string
	UpdateFrequency       int
	Private               bool
	IsLegendary           bool
	IsTown                bool
	DeathPenalty          bool
	UseEliteGenerators    bool
	SendBankContents      bool
	MaxOccupancy          int
	MaxLevel              int
	MinLevel              int
	RespawnZone           string
	RespawnSpawnPoint     string
	AllowPvPAnnouncements bool
	PVPType               int
	PVPMatchType          string
	AllowDuelRequest      bool
	EntryModifier         string
}

func NewZoneDefConfig() *ZoneDefConfig {
	return &ZoneDefConfig{}
}
