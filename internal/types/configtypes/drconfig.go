package configtypes

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
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
	Classes *DRClass `json:"classes"`
}

func (c *DRConfig) MergeParents(class *DRClassChildGroup) {
	for _, entity := range class.Entities {
		if entity.Extends != "" {
			parentClass, err := c.Get(entity.Extends)

			if err != nil || parentClass == nil || len(parentClass) == 0 {
				log.Debug("child not found " + entity.Extends)
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
					entity.Children = map[string]*DRClassChildGroup{}
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

func (c *DRConfig) mergeProperties(entity *DRClass, parentEntity *DRClass) {
	if entity.Properties == nil {
		entity.Properties = map[string]string{}
	}

	for propKey, propVal := range parentEntity.Properties {
		if _, ok := entity.Properties[propKey]; !ok {
			entity.Properties[propKey] = propVal
		}
	}
}

func (c *DRConfig) Get(fullgctype string) ([]*DRClassChildGroup, error) {
	fullgctype = strings.ToLower(fullgctype)
	splitGCType := strings.Split(fullgctype, ".")

	found, err := c.GetSimple(fullgctype)

	if err != nil {
		return nil, err
	}

	if len(found) == 0 {
		return nil, errors.New(fmt.Sprintf("could not find %s", fullgctype))
	}

	for _, child := range found {
		child.Name = splitGCType[len(splitGCType)-1]
		child.GCType = fullgctype

		c.MergeParents(child)
	}

	return found, nil
}

func (c *DRConfig) GetSimple(splitGCType string) ([]*DRClassChildGroup, error) {
	found := c.getFromGCType(strings.Split(splitGCType, "."), c.Classes.Children)

	if found == nil {
		return nil, errors.New("child not found")
	}
	return found, nil

}

func (c *DRConfig) getFromGCType(gcType []string, children map[string]*DRClassChildGroup) []*DRClassChildGroup {
	if child, ok := children[gcType[0]]; ok {
		if len(gcType) == 1 {
			return []*DRClassChildGroup{child}
		}

		foundSubChildren := make([]*DRClassChildGroup, 0)

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
	extends = strings.ToLower(extends)
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
		Classes: NewDRClass("root"),
	}
}

func (c *DRConfig) GenerateCategoryMap() (map[string]*DRCategory, error) {
	output := map[string]*DRCategory{}

	c.generateCategoryMap(c.Classes.Children, output, []string{})

	return output, nil
}

func (c *DRConfig) generateCategoryMap(classes map[string]*DRClassChildGroup, output map[string]*DRCategory, gcTypeName []string) {
	for className, classChildGroup := range classes {
		for _, entity := range classChildGroup.Entities {
			c.generateCategoryMap(entity.Children, output, append(gcTypeName, className))

			if entity.Extends != "" {
				parentsGCTypes := c.getParents(entity.Extends)
				parentGCType := strings.Join(parentsGCTypes, ".")

				fullGCType := parentGCType + "." + className

				curMap := output
				var curCategory *DRCategory = nil
				curGCType := ""

				for i := 0; i < len(parentsGCTypes); i++ {
					curGCType = parentsGCTypes[i]

					if _, ok := curMap[curGCType]; !ok {
						curMap[curGCType] = NewDRCategory()
					}

					curCategory = curMap[curGCType]
					curMap = curMap[curGCType].Children
				}

				realGCTypeName := strings.Join(append(gcTypeName, className), ".")

				if curCategory == nil {
					fmt.Println("nil category " + fullGCType)
				} else {
					curCategory.Classes[realGCTypeName] = true
				}

				fmt.Println(realGCTypeName)
			}
		}
	}
}

func (c *DRConfig) List(maxDepth int, predicate func(group *DRClassChildGroup) bool) ([]*DRClassChildGroup, error) {
	results := make([]*DRClassChildGroup, 0, 1024)

	results = addChildrenUntilDepth(maxDepth,
		predicate,
		c.Classes.Children, results)

	return results, nil
}

func addChildrenUntilDepth(
	maxDepth int,
	predicate func(group *DRClassChildGroup) bool,
	children map[string]*DRClassChildGroup,
	results []*DRClassChildGroup,
) []*DRClassChildGroup {
	for gName, group := range children {
		group.Name = gName

		if !predicate(group) {
			continue
		}

		truncateAtDepth(group, maxDepth, 0)

		results = append(results, group)
	}

	return results
}

func truncateAtDepth(group *DRClassChildGroup, maxDepth int, currentDepth int) {
	for _, entity := range group.Entities {
		if currentDepth >= maxDepth {
			entity.Children = nil
		} else {
			for gName, childGroup := range entity.Children {
				childGroup.Name = gName
				truncateAtDepth(childGroup, maxDepth, currentDepth+1)
			}
		}
	}
}
