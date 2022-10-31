package configparser

import "RainbowRunner/internal/database"

type DRConfig struct {
	Classes *database.DRClass `json:"classes"`
}
