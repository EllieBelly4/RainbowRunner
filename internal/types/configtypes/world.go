package configtypes

//go:generate go run ../../../scripts/generatelua -type=WorldConfig
type WorldConfig struct {
	Name                     string
	EncounterTable           string
	Generated                bool
	MazeDeadEndRemovalChance int
	MazeHeight               int
	MazeRandomness           int
	MazeSparseness           int
	MazeWidth                int
	TileSet                  string
	TileSize                 int
	WorldEntityTable         string
	WorldEntityTable2        string
	WorldEntityTable3        string
	Entities                 []*EntityConfig
}

func NewWorldConfig() *WorldConfig {
	return &WorldConfig{}
}
