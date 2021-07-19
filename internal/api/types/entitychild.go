package types

type EntityChildResolver struct {
	result interface{}
}

func (er *EntityChildResolver) ToEntity() (*Entity, bool) {
	res, ok := er.result.(*Entity)
	return res, ok
}

func (er *EntityChildResolver) ToComponent() (*Component, bool) {
	res, ok := er.result.(*Component)
	return res, ok
}
