package datatypes

import (
	"fmt"
	"math"
)

type Vector2Float32 struct {
	X, Y float32
}

func (f Vector2Float32) String() string {
	return fmt.Sprintf("(%f, %f)", f.X, f.Y)
}

func (v Vector2Float32) Distance(other Vector2Float32) float64 {
	xd := other.X - v.X
	yd := other.Y - v.Y

	a := math.Pow(float64(xd), 2)
	b := math.Pow(float64(yd), 2)

	return math.Sqrt(a + b)
}

func (f Vector2Float32) DivideByFloat32(f2 float32) Vector2Float32 {
	return Vector2Float32{
		X: f.X / f2,
		Y: f.Y / f2,
	}
}

func (f Vector2Float32) Sub(other Vector2Float32) Vector2Float32 {
	f.X -= other.X
	f.Y -= other.Y

	return f
}

func (f Vector2Float32) Normalize() Vector2Float32 {
	magnitude := f.Magnitude()

	return Vector2Float32{
		X: f.X / magnitude,
		Y: f.Y / magnitude,
	}
}

func (f Vector2Float32) Magnitude() float32 {
	return float32(math.Sqrt(float64(f.X*f.X + f.Y*f.Y)))
}

func (f Vector2Float32) Mul(i float32) Vector2Float32 {
	return Vector2Float32{
		X: f.X * i,
		Y: f.Y * i,
	}
}

func (f Vector2Float32) ToVector3Float32() Vector3Float32 {
	return Vector3Float32{
		X: f.X,
		Y: f.Y,
		Z: 0,
	}
}
