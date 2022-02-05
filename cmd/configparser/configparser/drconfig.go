package configparser

import (
	"RainbowRunner/internal/database"
	"errors"
	"fmt"
	"strings"
)

/**
{
	item: {
		classes: {
			abc: {
				properties: {}
				name: string
			}
		}
		namespaces: {
			mods: {

			}
		}
	}
}
*/

type DRConfig struct {
	Classes *database.DRClass `json:"classes"`
}

func (c *DRConfig) MergeParents(class *database.DRClassChildGroup) {
	for _, entity := range class.Entities {
		if entity.Extends != "" {
			parentClass, err := c.Get(entity.Extends)

			if err != nil || parentClass == nil || len(parentClass) == 0 {
				fmt.Println("child not found " + entity.Extends)
				return
			}

			if len(parentClass) != 1 || len(parentClass[0].Entities) != 1 {
				panic(fmt.Sprintf("multiple parents found for %s", class.Name))
				return
			}

			parent := parentClass[0]
			parentEntity := parent.Entities[0]

			c.MergeParents(parent)

			if parentEntity.Properties == nil {
				continue
			}

			for _, entity := range class.Entities {
				c.mergeProperties(entity, parentEntity)

				for parentChildName, parentChild := range parentEntity.Children {
					c.MergeParents(parentChild)

					if _, ok := entity.Children[parentChildName]; !ok {
						entity.Children[parentChildName] = parentChild
					} else {
						if len(entity.Children[parentChildName].Entities) > 1 || len(parentChild.Entities) > 1 {
							fmt.Println("cannot merge children as there are more than 1")
							continue
						}

						c.mergeProperties(entity.Children[parentChildName].Entities[0], parentChild.Entities[0])
					}
				}
			}
		} else {
			for _, child := range entity.Children {
				c.MergeParents(child)
			}
		}
	}
}

func (c *DRConfig) mergeProperties(entity *database.DRClass, parentEntity *database.DRClass) {
	if entity.Properties == nil {
		entity.Properties = map[string]string{}
	}

	for propKey, propVal := range parentEntity.Properties {
		if _, ok := entity.Properties[propKey]; !ok {
			entity.Properties[propKey] = propVal
		}
	}
}

func (c *DRConfig) Get(fullgctype string) ([]*database.DRClassChildGroup, error) {
	splitGCType := strings.Split(fullgctype, ".")

	found := c.getFromGCType(splitGCType, c.Classes.Children)

	if found == nil {
		return nil, errors.New("child not found")
	}

	for _, child := range found {
		child.Name = splitGCType[len(splitGCType)-1]
		child.GCType = fullgctype
		c.MergeParents(child)
	}

	return found, nil
}

func (c *DRConfig) getFromGCType(gcType []string, children map[string]*database.DRClassChildGroup) []*database.DRClassChildGroup {
	if child, ok := children[gcType[0]]; ok {
		if len(gcType) == 1 {
			return []*database.DRClassChildGroup{child}
		}

		foundSubChildren := make([]*database.DRClassChildGroup, 0)

		for _, subChild := range child.Entities {
			foundFromSubChild := c.getFromGCType(gcType[1:], subChild.Children)

			if foundFromSubChild != nil {
				foundSubChildren = append(foundSubChildren, foundFromSubChild...)
			}
		}

		return foundSubChildren
	}

	return nil
}

func NewDRConfig() *DRConfig {
	return &DRConfig{
		Classes: database.NewDRClass("root"),
	}
}
