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

type MTLColourType string

const (
	MTLColourTypeDiffuse            = "Kd"
	MTLColourTypeSpecular           = "Ks"
	MTLColourTypeAmbient            = "Ka"
	MTLColourTypeTransmissionFilter = "Tf"
)

type MTLColour struct {
	Type    MTLColourType
	R, G, B float32
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

func (b *MTLBuilder) WriteNewColour(materialName string, colour MTLColour) {
	if colour.Type == "" {
		panic("no colour type provided")
	}

	builder := b.materialDefs[materialName]

	builder.WriteString(string(colour.Type))
	builder.WriteRune(' ')

	builder.WriteString(fmt.Sprintf("%f", colour.R))
	builder.WriteRune(' ')

	builder.WriteString(fmt.Sprintf("%f", colour.G))
	builder.WriteRune(' ')

	builder.WriteString(fmt.Sprintf("%f", colour.B))
	builder.WriteRune(' ')

	builder.WriteRune('\n')
}

func (b *MTLBuilder) WriteNewAlpha(materialName string, f float32) {
	builder := b.materialDefs[materialName]

	builder.WriteString("d ")

	builder.WriteString(fmt.Sprintf("%f", f))
	builder.WriteRune('\n')
}

func NewMTLBuilder() *MTLBuilder {
	return &MTLBuilder{
		materialDefs: map[string]*strings.Builder{},
		textureFiles: map[string]bool{},
	}
}
