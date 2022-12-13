package main

import "text/template"

func mergeFuncMaps(maps ...template.FuncMap) template.FuncMap {
	merged := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
}
