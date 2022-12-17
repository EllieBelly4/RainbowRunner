package objects

//go:generate go run ../../scripts/generatelua -type=AnimationsList
type AnimationsList struct {
	Animations []*Animation
}

func NewAnimationsList() *AnimationsList {
	return &AnimationsList{}
}
