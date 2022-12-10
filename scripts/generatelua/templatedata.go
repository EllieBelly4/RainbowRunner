package main

import "fmt"

type TemplateData struct {
	Struct *StructDef
}

func (t *TemplateData) StructTypeNameVar() string {
	return fmt.Sprintf("lua%sTypeName", t.Struct.Name)
}
