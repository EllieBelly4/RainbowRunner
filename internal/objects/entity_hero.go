package objects

//go:generate go run ../../scripts/generatelua -type=Hero -extends=Unit
type Hero struct {
	*Unit
}

func NewHero(gcType string) *Hero {
	return &Hero{Unit: NewUnit(gcType)}
}
