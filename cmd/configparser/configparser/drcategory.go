package configparser

import (
	"RainbowRunner/internal/database"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

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

func GetGCTypesByCategory(category string, categories map[string]*DRCategory, minDepth int, filter *regexp.Regexp) ([]string, error) {
	rootElement := getElementByCategory(strings.Split(category, "."), categories)

	if rootElement == nil {
		return nil, errors.New(fmt.Sprintf("could not find root element with category %s", category))
	}

	results := make([]string, 0, 1024)

	rootElement.WalkChildGCTypes(func(gcType string, depth int) {
		if minDepth > -1 && depth < minDepth {
			return
		}

		if filter != nil && !filter.MatchString(gcType) {
			return
		}

		results = append(results, gcType)
	}, 0)

	return results, nil
}

func getElementByCategory(splitCategory []string, root map[string]*DRCategory) *DRCategory {
	if child, ok := root[splitCategory[0]]; ok {
		if len(splitCategory) == 1 {
			return child
		}

		return getElementByCategory(splitCategory[1:], root[splitCategory[0]].Children)
	}

	return nil
}

func NewDRCategory() *DRCategory {
	return &DRCategory{
		Children: map[string]*DRCategory{},
		Classes:  map[string]bool{},
	}
}

func (c *DRConfig) GenerateCategoryMap() (map[string]*DRCategory, error) {
	output := map[string]*DRCategory{}

	c.generateCategoryMap(c.Classes.Children, output, []string{})

	return output, nil
}

func (c *DRConfig) generateCategoryMap(classes map[string]*database.DRClassChildGroup, output map[string]*DRCategory, gcTypeName []string) {
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
