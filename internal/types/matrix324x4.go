package types

import (
	"RainbowRunner/pkg/datatypes"
	"fmt"
	"math"
)

type Matrix324x4 struct {
	Values [16]float32
}

var Matrix324x4Identity = Matrix324x4{
	Values: [16]float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	},
}

func (x Matrix324x4) String() string {
	return fmt.Sprintf(`%f %f %f %f
%f %f %f %f
%f %f %f %f
%f %f %f %f
`,
		x.Values[0], x.Values[4], x.Values[8], x.Values[12],
		x.Values[1], x.Values[5], x.Values[9], x.Values[13],
		x.Values[2], x.Values[6], x.Values[10], x.Values[14],
		x.Values[3], x.Values[7], x.Values[11], x.Values[15],
	)
}

func (x Matrix324x4) MultiplyVector3Float32(point datatypes.Vector3Float32) datatypes.Vector3Float32 {
	/**
	From game:
	V20*Z + V10*Y + V00*X + V30
	V21*Z + V11*Y + V01*X + V31
	V22*Z + V12*Y + V02*X + V32

	0 4 8  12
	1 5 9  13
	2 6 10 14
	3 7 11 15

	*/
	v := x.Values
	return datatypes.Vector3Float32{
		X: v[0]*point.X + v[4]*point.Y + v[8]*point.Z + v[12],
		Y: v[1]*point.X + v[5]*point.Y + v[9]*point.Z + v[13],
		Z: v[2]*point.X + v[6]*point.Y + v[10]*point.Z + v[14],
	}
}

func (x Matrix324x4) MultiplyVector3Float32NoTranslate(point datatypes.Vector3Float32) datatypes.Vector3Float32 {
	/**
	From game:
	V20*Z + V10*Y + V00*X + V30
	V21*Z + V11*Y + V01*X + V31
	V22*Z + V12*Y + V02*X + V32

	0 4 8  12
	1 5 9  13
	2 6 10 14
	3 7 11 15

	*/
	v := x.Values
	return datatypes.Vector3Float32{
		X: v[0]*point.X + v[4]*point.Y + v[8]*point.Z,
		Y: v[1]*point.X + v[5]*point.Y + v[9]*point.Z,
		Z: v[2]*point.X + v[6]*point.Y + v[10]*point.Z,
	}
}

func (x Matrix324x4) MultiplyMatrix324x4(matrix Matrix324x4) Matrix324x4 {
	a := x.Values
	b := matrix.Values

	return Matrix324x4{
		Values: [16]float32{
			// First Column
			// ma00*mb00 + ma01*mb10 + ma02*mb20 + ma03*mb30
			a[0]*b[0] + a[4]*b[1] + a[8]*b[2] + a[12]*b[3],
			//ma10*mb00 + ma11*mb10 + ma12*mb20 + ma13*mb30
			a[1]*b[0] + a[5]*b[1] + a[9]*b[2] + a[13]*b[3],
			//ma20*mb00 + ma21*mb10 + ma22*mb20 + ma23*mb30
			a[2]*b[0] + a[6]*b[1] + a[10]*b[2] + a[14]*b[3],
			//ma30*mb00 + ma31*mb10 + ma32*mb20 + ma33*mb30
			a[3]*b[0] + a[7]*b[1] + a[11]*b[2] + a[15]*b[3],

			// Second Column
			// ma00*mb01 + ma01*mb11 + ma02*mb21 + ma03*mb31
			a[0]*b[4] + a[4]*b[5] + a[8]*b[6] + a[12]*b[7],
			//ma10*mb01 + ma11*mb11 + ma12*mb21 + ma13*mb31
			a[1]*b[4] + a[5]*b[5] + a[9]*b[6] + a[13]*b[7],
			//ma20*mb01 + ma21*mb11 + ma22*mb21 + ma23*mb31
			a[2]*b[4] + a[6]*b[5] + a[10]*b[6] + a[14]*b[7],
			//ma30*mb01 + ma31*mb11 + ma32*mb21 + ma33*mb31
			a[3]*b[4] + a[7]*b[5] + a[11]*b[6] + a[15]*b[7],

			// Third Column
			//ma00*mb02 + ma01*mb12 + ma02*mb22 + ma03*mb32
			a[0]*b[8] + a[4]*b[9] + a[8]*b[10] + a[12]*b[11],
			//ma10*mb02 + ma11*mb12 + ma12*mb22 + ma13*mb32
			a[1]*b[8] + a[5]*b[9] + a[9]*b[10] + a[13]*b[11],
			//ma20*mb02 + ma21*mb12 + ma22*mb22 + ma23*mb32
			a[2]*b[8] + a[6]*b[9] + a[10]*b[10] + a[14]*b[11],
			//ma30*mb02 + ma31*mb12 + ma32*mb22 + ma33*mb32
			a[3]*b[8] + a[7]*b[9] + a[11]*b[10] + a[15]*b[11],

			// Fourth Column
			//ma00*mb03 + ma01*mb13 + ma02*mb23 + ma03*mb33
			a[0]*b[12] + a[4]*b[13] + a[8]*b[14] + a[12]*b[15],
			//ma10*mb03 + ma11*mb13 + ma12*mb23 + ma13*mb33
			a[1]*b[12] + a[5]*b[13] + a[9]*b[14] + a[13]*b[15],
			//ma20*mb03 + ma21*mb13 + ma22*mb23 + ma23*mb33
			a[2]*b[12] + a[6]*b[13] + a[10]*b[14] + a[14]*b[15],
			//ma30*mb03 + ma31*mb13 + ma32*mb23 + ma33*mb33
			a[3]*b[12] + a[7]*b[13] + a[11]*b[14] + a[15]*b[15],
		},
	}
}

func (x *Matrix324x4) Fix() {
	for i, value := range x.Values {
		if math.IsNaN(float64(value)) {
			x.Values[i] = 0
		}
	}
}
