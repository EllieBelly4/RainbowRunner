package configtypes

import (
	"RainbowRunner/internal/types"
	"fmt"
	"regexp"
	"strconv"
)

type DRClass struct {
	Name             string                        `json:"name,omitempty"`
	Extends          string                        `json:"extends,omitempty"`
	Properties       DRClassProperties             `json:"properties,omitempty"`
	Children         map[string]*DRClassChildGroup `json:"children,omitempty"`
	CustomProperties map[string]interface{}        `json:"customProperties,omitempty"`
}

func (c *DRClass) Find(class []string) *DRClass {
	for childName, child := range c.Children {
		if childName == class[0] {
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

	for childName, _ := range c.Children {
		//if modRegexp.MatchString(child.Name) {
		//	modCount++
		//}

		if childName != "description" {
			modCount++
		}
	}

	return modCount
}

func (c *DRClass) Slot() types.EquipmentSlot {
	desc, ok := c.Children["description"]

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
