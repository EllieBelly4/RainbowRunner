package types

type Component struct {
	id       int32
	typeName string
}

func (c *Component) TypeName() *string {
	return &c.typeName
}

func (c *Component) ID() *int32 {
	return &c.id
}
