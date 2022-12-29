package configtypes

//go:generate go run ../../../scripts/generatelua -type AnimationConfig
type AnimationConfig struct {
	ID int
	// This is used for remapping IDs to another animation
	// e.g. I think melee attacks by default have 3 animations but not all units have 3 attack animations,
	// so they map all 3 to the first animationID
	AnimationID int
	NumFrames   int

	// Unk
	TriggerTime int

	// Unk if we can actually use this
	SoundTriggerTime int

	// TODO extract this data from the config, animations do not have names but they usually have comments explaining what they are
	//Comment string
}
