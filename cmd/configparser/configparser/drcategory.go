package configparser

import (
	"RainbowRunner/internal/types/configtypes"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func GetGCTypesByCategory(category string, categories map[string]*configtypes.DRCategory, minDepth int, filter *regexp.Regexp) ([]string, error) {
	var rootElement *configtypes.DRCategory
	var depthMod = 0

	if category != "" {
		rootElement = getElementByCategory(strings.Split(category, "."), categories)
	} else {
		rootElement = &configtypes.DRCategory{
			Children: categories,
			Classes:  nil,
		}

		depthMod -= 1
	}

	if rootElement == nil {
		return nil, errors.New(fmt.Sprintf("could not find root element with category %s", category))
	}

	results := make([]string, 0, 1024)

	rootElement.WalkChildGCTypes(func(gcType string, depth int) {
		d := depth + depthMod
		if minDepth > -1 && d < minDepth {
			return
		}

		if filter != nil && !filter.MatchString(gcType) {
			return
		}

		results = append(results, gcType)
	}, 0)

	return results, nil
}

func getElementByCategory(splitCategory []string, root map[string]*configtypes.DRCategory) *configtypes.DRCategory {
	if child, ok := root[splitCategory[0]]; ok {
		if len(splitCategory) == 1 {
			return child
		}

		return getElementByCategory(splitCategory[1:], root[splitCategory[0]].Children)
	}

	return nil
}
