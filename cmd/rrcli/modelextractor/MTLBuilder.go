package modelextractor

import (
	"fmt"
	"strings"
)

type MTLTextureType string

const (
	MTLTextureTypeDiffuse = "map_Kd"
)

type MTLTexture struct {
	Type     MTLTextureType
	Filename string
}

type MTLBuilder struct {
	body         strings.Builder
	materialDefs map[string]*strings.Builder
	textureFiles map[string]bool
}

func (b *MTLBuilder) WriteNewMaterial(name string) {
	if _, ok := b.materialDefs[name]; ok {
		return
	}

	b.materialDefs[name] = &strings.Builder{}
	b.materialDefs[name].WriteString(fmt.Sprintf("newmtl %s\n", name))
}

func (b *MTLBuilder) HasMaterial(name string) bool {
	_, ok := b.materialDefs[name]

	return ok
}

func (b *MTLBuilder) String() string {
	complete := strings.Builder{}

	for _, builder := range b.materialDefs {
		complete.WriteString(builder.String())
	}

	return complete.String()
}

func (b *MTLBuilder) WriteNewTexture(materialName string, texture MTLTexture) {
	if texture.Type == "" {
		panic("no texture type provided")
	}

	if texture.Filename == "" {
		panic("no texture filename provided")
	}

	builder := b.materialDefs[materialName]

	builder.WriteString(string(texture.Type))
	builder.WriteRune(' ')

	builder.WriteString(texture.Filename)
	builder.WriteRune('\n')

	b.textureFiles[texture.Filename] = true
}

func (b *MTLBuilder) TextureFilenames() []string {
	list := make([]string, 0)

	for filename, _ := range b.textureFiles {
		list = append(list, filename)
	}

	return list
}

func NewMTLBuilder() *MTLBuilder {
	return &MTLBuilder{
		materialDefs: map[string]*strings.Builder{},
		textureFiles: map[string]bool{},
	}
}
