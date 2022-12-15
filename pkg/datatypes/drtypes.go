package datatypes

type DRFloat float32

func (f DRFloat) ToUInt() uint32 {
	return uint32(f * 256)
}

func DRFloatFromUInt(i uint32) DRFloat {
	return DRFloat(float32(i) / 256)
}
