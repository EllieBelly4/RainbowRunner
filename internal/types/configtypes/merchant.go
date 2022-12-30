package configtypes

type MerchantConfig struct {
	SellValueMod    float32
	BuyValueMod     float32
	RegenerateItems bool
	IDGenerator     int

	Inventories map[string]*MerchantInventoryConfig
	GCType      string
}

func (c *MerchantConfig) AddInventory(name string, config *MerchantInventoryConfig) {
	c.Inventories[name] = config
}

func NewMerchantConfig(gctype string) *MerchantConfig {
	return &MerchantConfig{
		GCType:      gctype,
		Inventories: map[string]*MerchantInventoryConfig{},
	}
}
