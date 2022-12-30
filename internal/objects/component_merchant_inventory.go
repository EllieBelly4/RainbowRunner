package objects

import "RainbowRunner/internal/types/configtypes"

//go:generate go run ../../scripts/generatelua -type=MerchantInventory -extends=Inventory
type MerchantInventory struct {
	*Inventory
	BaseConfig *configtypes.MerchantInventoryConfig
}

func NewMerchantInventory(gcType string, index byte) *MerchantInventory {
	gcObject := NewGCObject("MerchantInventory")
	gcObject.GCType = gcType

	return &MerchantInventory{
		Inventory: &Inventory{
			GCObject:    gcObject,
			InventoryID: index,
		},
	}
}

func NewMerchantInventoryFromConfig(config *configtypes.MerchantInventoryConfig) *MerchantInventory {
	inventory := NewMerchantInventory(config.GCType, byte(config.ID))

	inventory.BaseConfig = config

	return inventory
}
