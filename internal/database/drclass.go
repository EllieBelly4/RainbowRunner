package database

import (
	"RainbowRunner/internal/types"
	"fmt"
	"regexp"
	"strconv"
)

type DRClass struct {
	Name       string
	Properties map[string]string
	Children   []*DRClass
}

func (c *DRClass) Find(class []string) *DRClass {
	for _, child := range c.Children {
		if child.Name == class[0] {
			if len(class) > 1 {
				return c.Find(class[1:])
			} else {
				return child
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
	desc := c.Find([]string{"Description"})

	// Mods do not have descriptions
	if desc == nil {
		panic(fmt.Sprintf("%s does not have a description", c.Name))
	}

	slotInt, err := strconv.Atoi(desc.Properties["SlotType"])

	if err != nil {
		panic(err)
	}

	return types.EquipmentSlot(slotInt)
}

func NewDRClass(className string) *DRClass {
	return &DRClass{
		Name:       className,
		Properties: make(map[string]string),
		Children:   make([]*DRClass, 0),
	}
}
