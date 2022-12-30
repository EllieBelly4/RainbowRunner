package drconfigtypes

type DRCategory struct {
	Children map[string]*DRCategory
	Classes  map[string]bool
}

func (c *DRCategory) WalkChildGCTypes(f func(gcType string, depth int), currentDepth int) {
	if c.Children != nil {
		for _, child := range c.Children {
			child.WalkChildGCTypes(f, currentDepth+1)
		}
	}

	for gcType, _ := range c.Classes {
		f(gcType, currentDepth)
	}
}

func NewDRCategory() *DRCategory {
	return &DRCategory{
		Children: map[string]*DRCategory{},
		Classes:  map[string]bool{},
	}
}
