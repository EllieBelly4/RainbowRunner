package objects

import (
	"RainbowRunner/internal/types/configtypes"
)

//go:generate go run ../../scripts/generatelua -type=Animation
type Animation struct {
	ID               int
	NumFrames        int
	TriggerTime      int
	SoundTriggerTime int
	AnimationID      int
}

func NewAnimationFromConfig(a *configtypes.AnimationConfig) *Animation {
	animation := NewAnimation()

	animation.ID = a.ID
	animation.AnimationID = a.AnimationID
	animation.NumFrames = a.NumFrames
	animation.TriggerTime = a.TriggerTime
	animation.SoundTriggerTime = a.SoundTriggerTime

	return animation
}

func NewAnimation() *Animation {
	return &Animation{}
}
