package configtypes

//go:generate go run ../../../scripts/generatelua -type=InventoryConfig
type InventoryConfig struct {
	ID int

	Description *InventoryDescConfig
	GCType      string
}

func NewInventoryConfig(gctype string) *InventoryConfig {
	return &InventoryConfig{
		GCType: gctype,
	}
}

type InventoryDescConfig struct {
	Height int
	Width  int
	Label  string
}

func NewInventoryDescConfig() *InventoryDescConfig {
	return &InventoryDescConfig{}
}
