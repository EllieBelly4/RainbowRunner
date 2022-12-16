package objects

//go:generate go run ../../scripts/generatelua -type=Entity -extends=GCObject
type Entity struct {
	*GCObject
}

func NewEntity(gctype string) *Entity {
	return &Entity{
		GCObject: NewGCObject(gctype),
	}
}
