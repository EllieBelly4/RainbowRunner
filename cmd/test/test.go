package main

import (
	"RainbowRunner/internal/database"
	"fmt"
)

func main() {
	database.LoadEquipmentFixtures()

	weapon := database.Weapons.Find("1HSwordMythicPAL.1HSwordMythic6")

	fmt.Printf("%+v\n", weapon)
	fmt.Printf("%d\n", weapon.ModCount())

	armour := database.Armour.Find("ChainArmor1PAL.ChainArmor1-2")

	fmt.Printf("%+v\n", armour)
	fmt.Printf("%d\n", armour.ModCount())

	//data, err := json.MarshalIndent(armour, "", "  ")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(string(data))
}
