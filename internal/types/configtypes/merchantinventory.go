package configtypes

type MerchantInventoryConfig struct {
	*InventoryConfig

	MaxItemLevel      int
	MinItemLevel      int
	StaticContents    bool
	AutoGenerateItems bool
	ItemGenerator     string
}

func NewMerchantInventoryConfig(gctype string) *MerchantInventoryConfig {
	return &MerchantInventoryConfig{
		InventoryConfig: NewInventoryConfig(gctype),
	}
}
