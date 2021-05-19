package main

import (
	"RainbowRunner/internal/byter"
	"RainbowRunner/internal/objects"
	"encoding/hex"
	"fmt"
)

func main() {
	player := objects.NewGCObject("Player")
	player.ID = 0x01
	player.Name = "Ellie"

	avatar := objects.NewGCObject("Avatar")
	avatar.ID = 0x02
	avatar.Name = "EllieAvatar"
	avatar.Properties = []objects.GCObjectProperty{
		objects.Uint8Prop("Hair", 0x00),
		objects.Uint8Prop("HairColor", 0x00),
		objects.Uint8Prop("Face", 0x01),
		objects.Uint8Prop("FaceFeature", 0x01),
		objects.Uint8Prop("Skin", 0x01),
	}

	modifiers := objects.NewGCObject("Modifiers")
	modifiers.ID = 0x03
	modifiers.Name = "ModifiersAreHere"
	modifiers.Properties = []objects.GCObjectProperty{
		objects.Uint8Prop("IDGenerator", 0x01),
	}

	avatar.AddChild(modifiers)
	player.AddChild(avatar)

	body := byter.NewLEByter(make([]byte, 0, 1024))
	player.Serialise(body)

	fmt.Printf("%s\n", hex.Dump(body.Data()))
}
