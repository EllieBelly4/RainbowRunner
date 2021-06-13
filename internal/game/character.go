package game

import (
	"RainbowRunner/internal/byter"
	"RainbowRunner/internal/objects"
	log "github.com/sirupsen/logrus"
)

type CharacterMessage byte

const (
	CharacterConnected CharacterMessage = iota
	CharacterDisconnected
	CharacterCreate
	CharacterGetList
	CharacterDelete
	CharacterPlay
)

func handleCharacterChannelMessages(conn *RRConn, msgType byte, reader *byter.Byter) error {
	switch CharacterMessage(msgType) {
	case CharacterConnected:
		handleCharacterConnected(conn)
	case CharacterPlay:
		handleCharacterPlay(conn)
	case CharacterGetList:
		sendCharacterList(conn)
	case CharacterCreate:
		handleCharacterCreate(conn, reader)
	default:
		return UnhandledChannelMessageError
	}

	return nil
}

func handleCharacterCreate(conn *RRConn, reader *byter.Byter) {
	name := reader.String()
	class := reader.String()
	reader.UInt8() // Unk
	reader.UInt8() // Face
	reader.UInt8() // Hair
	reader.UInt8() // Hair Colour

	log.Infof("New character created %s (%s)", name, class)

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(CharacterChannel)) // Character channel
	body.WriteByte(byte(CharacterCreate))
	body.WriteUInt32(0x01)

	body.WriteCString("Ellie")

	sendPlayer(body)

	WriteCompressedA(conn, 0x01, 0x0f, body)
}

func handleCharacterPlay(conn *RRConn) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(CharacterChannel))
	body.WriteByte(byte(CharacterPlay))
	WriteCompressedA(conn, 0x01, 0x0f, body)
}

func handleCharacterConnected(conn *RRConn) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(CharacterChannel))   // Character channel
	body.WriteByte(byte(CharacterConnected)) // Connected
	WriteCompressedA(conn, 0x01, 0x0f, body)
}

func sendPlayer(body *byter.Byter) {
	//hero := objects.NewGCObject("Hero")
	//hero.ID = 0xBABAF00B
	//hero.Name = "EllieHero"
	//hero.Properties = []objects.GCObjectProperty{
	//	objects.Uint32Prop("Level", 10),
	//	objects.Uint32Prop("Experience", 1000),
	//}

	//professionTitle := objects.NewGCObject("ProfessionTitle")
	//professionTitle.ID = 0xBAB5BAB5
	//professionTitle.Name = "FIGHTER"
	//professionTitle.Properties = []objects.GCObjectProperty{
	//	objects.Uint32Prop("Elements", 0x01),
	//}

	//rpgSettings := objects.NewGCObject("RPGSettings")
	//rpgSettings.ID = 0x55665566
	//rpgSettings.Name = "EllieRPG"

	//rpgSettings.AddChild(professionTitle)

	//heroDesc := objects.NewGCObject("HeroDesc")
	//heroDesc.ID = 0xF00DB0B0
	//heroDesc.Name = "EllieHeroDesc"

	//hero.AddChild(heroDesc)
	//hero.AddChild(rpgSettings)

	//player.AddChild(hero)
	avatar := getAvatar(0x02)
	//avatar2 := getAvatar(0x02)

	player := objects.NewGCObject("Player")
	player.ID = 0x01
	player.Name = "Player Name"

	player.AddChild(avatar)
	player.Properties = []objects.GCObjectProperty{
		objects.StringProp("Name", "Ellie"),
	}

	player.Serialise(body)

	body.WriteCString("Unk")  // Specific to player::readObject
	body.WriteCString("Unk2") // Specific to player::readObject
	//body.WriteCString("Unk3") // Specific to player::readObject
	//body.WriteCString("Unk4") // Specific to player::readObject
	body.WriteUInt32(0x01) // Specific to player::readObject
	body.WriteUInt32(0x01) // Specific to player::readObject

	avatar.Serialise(body)

	body.WriteByte(0x01)
	body.WriteByte(0x01)

	body.WriteCString("Formidable")

	body.WriteByte(0x01)
	body.WriteByte(0x01)

	body.WriteUInt32(0x01)
}

func getAvatar(ID uint32) *objects.GCObject {
	avatar := objects.NewGCObject("Avatar")
	avatar.GCType = "avatar.classes.fighterfemale"
	//avatar.GCType = "avatar.base.avatar"
	avatar.ID = ID
	avatar.Name = "Avatar Name"
	avatar.Properties = []objects.GCObjectProperty{
		objects.Uint32Prop("Hair", 0x01),
		objects.Uint32Prop("HairColor", 0x00),
		objects.Uint32Prop("Face", 0x01),
		objects.Uint32Prop("FaceFeature", 0x01),
		objects.Uint32Prop("Skin", 0x01),
		objects.Uint32Prop("Level", 100),
	}

	modifiers := objects.NewGCObject("Modifiers")
	modifiers.GCType = "Modifiers"
	modifiers.ID = 0x02
	modifiers.Name = "Mod Name"
	modifiers.Properties = []objects.GCObjectProperty{
		objects.Uint32Prop("IDGenerator", 0x01),
	}

	manipulators := objects.NewGCObject("Manipulators")
	manipulators.ID = 0x03
	manipulators.Name = "ManipulateMe"

	//animationList := objects.NewGCObject("AnimationList")
	//animationList.ID = 0xBABAF00E
	//animationList.Name = "EllieAnimations"

	avatarSkills := objects.NewGCObject("Skills")
	avatarSkills.GCType = "avatar.base.skills"
	avatarSkills.ID = 0x04
	avatarSkills.Name = "EllieSkills"

	avatarDesc := objects.NewGCObject("AvatarDesc")
	avatarDesc.GCType = "avatar.classes.fighterfemale.description"
	avatarDesc.ID = 0xBABAF00D
	avatarDesc.Name = "EllieAvatarDesc"

	avatarEquipment := objects.NewGCObject("Equipment")
	avatarEquipment.GCType = "avatar.base.Equipment"
	avatarEquipment.ID = 0x5000BAAD
	avatarEquipment.Name = "EllieEquipment"

	//.text:0058E550     ; struct DFCClass *__thiscall Armor::getClass(Armor *__hidden this)
	//.text:0058E550 000 mov     eax, ?Class@Armor@@2PAVDFCClass@@A ; DFCClass * Armor::Class

	weapon := objects.NewGCObject("MeleeWeapon")
	weapon.GCType = "1HSword8PAL.1HSword8-1"
	weapon.ID = 0x12354567
	weapon.Name = "EllieWeapon"

	//TODO finish
	weapon.Properties = []objects.GCObjectProperty{
		//objects.StringProp("Label", "Hello"),
	}

	slot := objects.NewGCObject("EquipmentSlot")
	slot.GCType = "avatar.base.Equipment.Description.PrimaryWeaponSlot"
	slot.ID = 0x12354567
	slot.Name = "EllieWeaponSlot"

	slot.Properties = []objects.GCObjectProperty{
		objects.Uint32Prop("SlotID", 0x00),
		objects.Uint32Prop("SlotType", 0xFFFFFFFF),
	}

	slot.AddChild(weapon)
	avatar.AddChild(slot)

	//avatarDesc.Properties = []objects.GCObjectProperty{
	//	objects.StringProp("PVEStartSpawnPoint", "Start"),
	//}

	//worldEntityDesc := objects.NewGCObject("WorldEntityDesc")
	//worldEntityDesc.ID = 0xBABABABA
	//worldEntityDesc.Name = "EllieWorldEntityDesc"

	unitBehaviour := objects.NewGCObject("UnitBehavior")
	unitBehaviour.GCType = "avatar.base.UnitBehavior"
	unitBehaviour.ID = 0x05
	unitBehaviour.Name = "EllieBehaviour"

	//avatar.AddChild(visual)
	//avatar.AddChild(rpgSettings)
	avatar.AddChild(modifiers)
	avatar.AddChild(avatarSkills)
	avatar.AddChild(manipulators)
	avatar.AddChild(avatarDesc)
	avatar.AddChild(avatarEquipment)
	avatar.AddChild(unitBehaviour)
	return avatar
}
