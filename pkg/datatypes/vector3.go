package datatypes

import "fmt"

type Vector3 struct {
	X, Y, Z int32
}

func (v Vector3) String() string {
	return fmt.Sprintf("(%d, %d, %d)", v.X, v.Y, v.Z)
}

func (v Vector3) ToVector2() Vector2 {
	return Vector2{
		X: v.X,
		Y: v.Y,
	}
}
