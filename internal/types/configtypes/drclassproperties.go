package configtypes

import "regexp"

type DRClassProperties map[string]string

var quoteRemoveRegex = regexp.MustCompile("^[\"']|[\"']$")

func (p *DRClassProperties) StringVal(key string) string {
	return quoteRemoveRegex.ReplaceAllString((*p)[key], "")
}
