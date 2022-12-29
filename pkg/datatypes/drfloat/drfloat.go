package drfloat

type DRFloat int32

func FromFloat32(val float32) DRFloat {
	return DRFloat(val * 256)
}

func FromUInt32(val uint32) DRFloat {
	return DRFloat(val * 256)
}

func FromInt32(val int32) DRFloat {
	return DRFloat(val * 256)
}

func (f DRFloat) ToUInt() uint32 {
	return uint32(f / 256)
}

func (f DRFloat) ToInt() int32 {
	return int32(f / 256)
}

func (f DRFloat) ToFloat32() float32 {
	return float32(f) / 256
}

func (f DRFloat) ToWire() uint32 {
	return uint32(f)
}
