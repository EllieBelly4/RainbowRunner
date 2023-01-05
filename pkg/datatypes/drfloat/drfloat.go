package drfloat

import "fmt"

type DRFloat int32

func (f DRFloat) String() string {
	return fmt.Sprintf("%f", float32(f)/256.0)
}

func FromFloat32(val float32) DRFloat {
	return DRFloat(val * 256)
}

func FromUInt32(val uint32) DRFloat {
	return DRFloat(val * 256)
}

func FromInt32(val int32) DRFloat {
	return DRFloat(val * 256)
}

func (f DRFloat) Sub(other DRFloat) DRFloat {
	return DRFloat(int32(f) - int32(other))
}

func (f DRFloat) Add(other DRFloat) DRFloat {
	return DRFloat(int32(f) + int32(other))
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
