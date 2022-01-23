package pkg

type Vector2 struct {
	X, Y int32
}

type Vector2Float32 struct {
	X, Y float32
}

type Vector3 struct {
	X, Y, Z int32
}

func (v Vector3) ToVector2() Vector2 {
	return Vector2{
		X: v.X,
		Y: v.Y,
	}
}

type Vector3Short struct {
	X, Y, Z int16
}

type Vector3Float32 struct {
	X, Y, Z float32
}

type Transform struct {
	Position Vector3
	Rotation int
}
