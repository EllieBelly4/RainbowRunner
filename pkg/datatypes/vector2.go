package datatypes

import "math"

type Vector2 struct {
	X, Y int32
}

func (v Vector2) Distance(other Vector2) float64 {
	xd := other.X - v.X
	yd := other.Y - v.Y

	a := math.Pow(float64(xd), 2)
	b := math.Pow(float64(yd), 2)

	return math.Sqrt(a + b)
}

func (v Vector2) ToVector3Float32() Vector3Float32 {
	return Vector3Float32{
		X: float32(v.X),
		Y: float32(v.Y),
	}
}

func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}
