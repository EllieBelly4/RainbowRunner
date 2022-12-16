package datatypes

import (
	"fmt"
	"github.com/yuin/gopher-lua"
)

type Vector3Float32 struct {
	X, Y, Z float32
}

func (v Vector3Float32) ToVector2Float32() Vector2Float32 {
	return Vector2Float32{
		X: v.X,
		Y: v.Y,
	}
}

func (f Vector3Float32) Add(other Vector3Float32) Vector3Float32 {
	return Vector3Float32{
		X: f.X + other.X,
		Y: f.Y + other.Y,
		Z: f.Z + other.Z,
	}
}

func (f Vector3Float32) Sub(other Vector3Float32) Vector3Float32 {
	return Vector3Float32{
		X: f.X - other.X,
		Y: f.Y - other.Y,
		Z: f.Z - other.Z,
	}
}

func (f Vector3Float32) String() string {
	return fmt.Sprintf("(%f, %f, %f)", f.X, f.Y, f.Z)
}

func (f Vector3Float32) DivideByFloat32(i float32) Vector3Float32 {
	return Vector3Float32{
		X: f.X / i,
		Y: f.Y / i,
		Z: f.Z / i,
	}
}

func (f Vector3Float32) MultiplyByFloat32(i float32) Vector3Float32 {
	return Vector3Float32{
		X: f.X * i,
		Y: f.Y * i,
		Z: f.Z * i,
	}
}

func (f Vector3Float32) ToLua(state *lua.LState) lua.LValue {
	ud := state.NewUserData()
	ud.Value = f
	state.SetMetatable(ud, state.GetTypeMetatable("Vector3"))
	return ud
}
