package objects

import "RainbowRunner/pkg/byter"

// GCClassRegistry::readType()
func writeGCType(b *byter.Byter, t string) {
	b.WriteByte(0xFF) // GetType
	b.WriteCString(t)
}
