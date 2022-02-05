package configparser

import "RainbowRunner/internal/database"

/**
{
	item: {
		classes: {
			abc: {
				properties: {}
				name: string
			}
		}
		namespaces: {
			mods: {

			}
		}
	}
}
*/

type DRConfig struct {
	Classes *database.DRClass `json:"classes"`
}

func (c *DRConfig) MergeParents() {
	//c.mergeParents(c.Classes)
}

func NewDRConfig() *DRConfig {
	return &DRConfig{
		Classes: database.NewDRClass("root"),
	}
}
