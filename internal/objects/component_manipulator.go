package objects

//go:generate go run ../../scripts/generatelua -type=Manipulator -extends=Component
type Manipulator struct {
	*Component
	Slot int // Not entirely sure this is a "slot", but that's how it's currently being used by Equipment and Skills
}

func NewManipulator(gctype, gcnativetype string) *Manipulator {
	component := NewComponent(gctype, gcnativetype)
	return &Manipulator{Component: component}
}
