package datatypes

import "fmt"

type Vector2Float32 struct {
	X, Y float32
}

func (f Vector2Float32) String() string {
	return fmt.Sprintf("(%f, %f)", f.X, f.Y)
}

func (f Vector2Float32) DivideByFloat32(f2 float32) Vector2Float32 {
	return Vector2Float32{
		X: f.X / f2,
		Y: f.Y / f2,
	}
}
