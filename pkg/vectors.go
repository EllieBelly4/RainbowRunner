package pkg

type Vector2 struct {
	X, Y int32
}

type Vector3 struct {
	X, Y, Z int32
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
