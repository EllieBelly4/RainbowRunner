package database

type DRNamespace struct {
	Name     string
	Classes  map[string]*DRClass
	Children map[string]*DRNamespace
}

func NewDRNamespace(name string) *DRNamespace {
	return &DRNamespace{
		Name:     name,
		Children: make(map[string]*DRNamespace),
		Classes:  make(map[string]*DRClass),
	}
}
