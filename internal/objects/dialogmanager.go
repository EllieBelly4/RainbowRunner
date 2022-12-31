package objects

import (
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
)

//go:generate go run ../../scripts/generatelua -type=DialogManager -extends=GCObject
type DialogManager struct {
	*GCObject
}

func (q DialogManager) Type() drobjecttypes.DRObjectType {
	return drobjecttypes.DRObjectManager
}

func NewDialogManager() *DialogManager {
	q := &DialogManager{
		GCObject: NewGCObject("DialogManager"),
	}

	q.GCType = "DialogManager"

	return q
}

func (q DialogManager) WriteInit(b *byter.Byter) {

}

func (q DialogManager) WriteUpdate(b *byter.Byter) {
	panic("implement me")
}
