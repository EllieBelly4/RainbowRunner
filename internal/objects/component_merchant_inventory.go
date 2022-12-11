package objects

//go:generate go run ../../scripts/generatelua -type=MerchantInventory -extends=Inventory
type MerchantInventory struct {
	*Inventory
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
