package objects

type IComponent interface {
	IsComponent()
}

type Component struct {
}

func (Component) IsComponent() {
}
