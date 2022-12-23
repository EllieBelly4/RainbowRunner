package objects

import (
	"RainbowRunner/internal/types/drobjecttypes"
	"RainbowRunner/pkg/byter"
)

type DialogManager struct {
	*GCObject
}

func (q DialogManager) Type() drobjectypes.DRObjectType {
	return drobjectypes.DRObjectManager
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
