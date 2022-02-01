package main

import (
	"RainbowRunner/internal/types"
	"RainbowRunner/pkg/datatypes"
	"fmt"
)

func main() {
	//mat1 := types.Matrix324x4{Values: [16]float32{
	//	1, 0, 0, 0,
	//	0, 1, 0, 0,
	//	0, 0, 1, 0,
	//	3, 2, 2, 1,
	//}}
	//
	//mat2 := types.Matrix324x4{Values: [16]float32{
	//	1, 0, 0, 0,
	//	0, 1, 0, 0,
	//	0, 0, 1, 0,
	//	3, 5, 7, 1,
	//}}

	/**
	  1	  0	  0	  0
	  0	  0.71	  -0.71	  0
	  0	  0.71	  0.71	  0
	  0	  0	  0	  1
	*/

	rotMat := types.Matrix324x4{Values: [16]float32{
		0.3420202, 0.9396926, 0, 0,
		-0.9396926, 0.3420202, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}}

	vector := datatypes.Vector3Float32{
		Y: 2,
	}

	//fmt.Println(mat1.MultiplyVector3Float32(vector).String())
	//fmt.Println(mat1.MultiplyMatrix324x4(mat2).String())
	fmt.Println(rotMat.MultiplyVector3Float32(vector).String())

	//database.LoadEquipmentFixtures()
	//
	//weapon := database.Weapons.Find("1HSwordMythicPAL.1HSwordMythic6")
	//
	//fmt.Printf("%+v\n", weapon)
	//fmt.Printf("%d\n", weapon.ModCount())
	//
	//armour := database.Armour.Find("ChainArmor1PAL.ChainArmor1-2")
	//
	//fmt.Printf("%+v\n", armour)
	//fmt.Printf("%d\n", armour.ModCount())

	//data, err := json.MarshalIndent(armour, "", "  ")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(string(data))
}
