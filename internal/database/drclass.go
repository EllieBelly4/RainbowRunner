package database

import (
	"RainbowRunner/internal/types"
	"fmt"
	"regexp"
	"strconv"
)

type DRClassChildGroup struct {
	Name     string     `json:"name"`
	Extends  string     `json:"extends,omitempty"`
	Entities []*DRClass `json:"entities"`
	Entity   *DRClass   `json:"entity,omitempty"`
}

type DRClass struct {
	Name       string                        `json:"name,omitempty"`
	Extends    string                        `json:"extends,omitempty"`
	Properties map[string]string             `json:"properties,omitempty"`
	Children   map[string]*DRClassChildGroup `json:"children,omitempty"`
}

func (c *DRClass) Find(class []string) *DRClass {
	for _, child := range c.Children {
		if child.Name == class[0] {
			if len(class) > 1 {
				return c.Find(class[1:])
			} else {
				return child.Entities[0]
			}
		}
	}

	return nil
}

var modRegexp = regexp.MustCompile("^Mod[0-9]+$")

func (c *DRClass) ModCount() int {
	modCount := 0

	for _, child := range c.Children {
		//if modRegexp.MatchString(child.Name) {
		//	modCount++
		//}

		if child.Name != "Description" {
			modCount++
		}
	}

	return modCount
}

func (c *DRClass) Slot() types.EquipmentSlot {
	desc, ok := c.Children["Description"]

	// Mods do not have descriptions
	if !ok {
		panic(fmt.Sprintf("%s does not have a description", c.Name))
	}

	slotInt, err := strconv.Atoi(desc.Entities[0].Properties["SlotType"])

	if err != nil {
		panic(err)
	}

	return types.EquipmentSlot(slotInt)
}

func NewDRClass(className string) *DRClass {
	return &DRClass{
		Name:       className,
		Properties: make(map[string]string),
		Children:   make(map[string]*DRClassChildGroup),
	}
}

func NewDRClassChildGroup(className string) *DRClassChildGroup {
	return &DRClassChildGroup{
		Name:     className,
		Entities: make([]*DRClass, 0),
	}
}
