package configtypes

type MerchantConfig struct {
	SellValueMod float32
	BuyValueMod  float32
}

func NewMerchantConfig() *MerchantConfig {
	return &MerchantConfig{}
}

type MerchantInventory struct {
}

func NewMerchantInventory() *MerchantInventory {
	return &MerchantInventory{}
}
