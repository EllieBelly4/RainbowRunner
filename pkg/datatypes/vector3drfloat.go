package datatypes

import (
	"RainbowRunner/pkg/datatypes/drfloat"
)

type Vector3DRFloat struct {
	X drfloat.DRFloat
	Y drfloat.DRFloat
	Z drfloat.DRFloat
}

func (f Vector3DRFloat) String() string {
	return f.X.String() + "," + f.Y.String() + "," + f.Z.String()
}
