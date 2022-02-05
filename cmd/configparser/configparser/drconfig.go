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

			for _, entity := range class.Entities {
				c.mergeProperties(entity, parentEntity)

				if entity.Children == nil {
					entity.Children = map[string]*database.DRClassChildGroup{}
				}

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

	found, err := c.GetSimple(fullgctype)

	if err != nil {
		return nil, err
	}

	for _, child := range found {
		child.Name = splitGCType[len(splitGCType)-1]
		child.GCType = fullgctype
		c.MergeParents(child)
	}

	return found, nil
}

func (c *DRConfig) GetSimple(splitGCType string) ([]*database.DRClassChildGroup, error) {
	found := c.getFromGCType(strings.Split(splitGCType, "."), c.Classes.Children)

	if found == nil {
		return nil, errors.New("child not found")
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
				for _, singleFoundFromSubChild := range foundFromSubChild {
					foundSubChildren = append(foundSubChildren, singleFoundFromSubChild)
				}
			}
		}

		return foundSubChildren
	}

	return nil
}

func (c *DRConfig) getParents(extends string) []string {
	splitKey := strings.Split(extends, ".")
	parent, err := c.GetSimple(extends)

	if err != nil || len(parent) == 0 {
		return splitKey
	}

	if len(parent) > 1 {
		panic("wrong number of parents found")
	}

	if parent[0].Entities[0].Extends != "" && parent[0].Entities[0].Extends != splitKey[0] {
		parentTypes := c.getParents(parent[0].Entities[0].Extends)
		splitKey = append(parentTypes, splitKey...)
	}

	return splitKey
}

//func (c *DRConfig) getParentTypes(extends string) []string {
//	return c.getParentTypesArray(strings.Split(extends, "."))
//}

/*func (c *DRConfig) getParentTypesArray(gcType []string) []string {
	result := gcType
	parent, err := c.Get(strings.Join(gcType, "."))

	if err != nil {
		return gcType
	}

	if len(parent) == 0 {
		return gcType
	}

	if len(parent) > 1 {
		panic("wrong number of parents found")
	}

	if parent[0].Entities[0].Extends != "" {
		return c.getParentTypesArray(append(result, parent[0].Name))
	}

	return result
}*/

func NewDRConfig() *DRConfig {
	return &DRConfig{
		Classes: database.NewDRClass("root"),
	}
}
