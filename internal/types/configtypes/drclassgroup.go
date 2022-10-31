package configtypes

type DRClassChildGroup struct {
	Name     string     `json:"name,omitempty"`
	Entities []*DRClass `json:"entities"`
	GCType   string     `json:"gcType,omitempty"`
}

func NewDRClass(className string) *DRClass {
	return &DRClass{
		Name:             className,
		Properties:       make(map[string]string),
		Children:         make(map[string]*DRClassChildGroup),
		CustomProperties: make(map[string]interface{}),
	}
}

func NewDRClassChildGroup(className string) *DRClassChildGroup {
	return &DRClassChildGroup{
		Name:     className,
		Entities: make([]*DRClass, 0),
	}
}
