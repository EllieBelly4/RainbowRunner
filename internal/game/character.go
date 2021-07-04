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
	avatar := getAvatar()
	//avatar2 := getAvatar(0x02)

	player := objects.NewGCObject("Player")
	player.Name = "Player Name"

	player.AddChild(avatar)
	player.Properties = []objects.GCObjectProperty{
		objects.StringProp("Name", "Ellie"),
	}

	//slot6 := objects.NewGCObject("EquipmentSlot")
	//slot6.GCType = "avatar.base.Equipment.Description.Armor"
	//slot6.Name = "EllieArmorSlot"
	//
	//armor := objects.NewGCObject("Armor")
	//armor.GCType = "ChainArmor1PAL.ChainArmor1-1"
	//armor.Name = "EllieArmour"

	//slot6.AddChild(armor)
	//player.AddChild(slot6)

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

func getAvatar() *objects.GCObject {
	avatar := objects.NewGCObject("Avatar")
	avatar.GCType = "avatar.classes.FighterMale"
	//avatar.GCType = "avatar.base.avatar"
	avatar.Name = "Avatar Name"
	avatar.Properties = []objects.GCObjectProperty{
		objects.Uint32Prop("Hair", 0x01),
		objects.Uint32Prop("HairColor", 0x00),
		objects.Uint32Prop("Face", 0),
		objects.Uint32Prop("FaceFeature", 0),
		objects.Uint32Prop("Skin", 0x01),
		objects.Uint32Prop("Level", 50),
	}

	metrics := objects.NewAvatarMetrics(0xFE34BE34, "EllieMetrics")

	modifiers := objects.NewGCObject("Modifiers")
	modifiers.GCType = "Modifiers"
	modifiers.Name = "Mod Name"
	modifiers.Properties = []objects.GCObjectProperty{
		objects.Uint32Prop("IDGenerator", 0x01),
	}

	manipulators := objects.NewGCObject("Manipulators")
	manipulators.Name = "ManipulateMe"

	dialogManager := objects.NewGCObject("DialogManager")
	dialogManager.Name = "EllieDialogManager"

	//animationList := objects.NewGCObject("AnimationList")
	//animationList.Name = "EllieAnimations"

	avatarSkills := objects.NewGCObject("Skills")
	avatarSkills.GCType = "avatar.base.skills"
	avatarSkills.Name = "EllieSkills"

	//avatarDesc := objects.NewGCObject("AvatarDesc")
	//avatarDesc.GCType = "avatar.classes.fighterfemale.description"
	//avatarDesc.Name = "EllieAvatarDesc"

	avatarEquipment := objects.NewGCObject("Equipment")
	avatarEquipment.GCType = "avatar.base.Equipment"
	avatarEquipment.Name = "EllieEquipment"

	//.text:0058E550     ; struct DFCClass *__thiscall Armor::getClass(Armor *__hidden this)
	//.text:0058E550 000 mov     eax, ?Class@Armor@@2PAVDFCClass@@A ; DFCClass * Armor::Class

	weapon := objects.NewGCObject("MeleeWeapon")
	weapon.ID = 0xAAAABBBB
	weapon.GCType = "1HMace1PAL.1HMace1-1"
	weapon.Name = "EllieWeapon"

	weaponDesc := objects.NewGCObject("MeleeWeaponDesc")
	weaponDesc.ID = 0xAAAABBBB
	weaponDesc.GCType = "1HMace1PAL.1HMace1-1.Description"
	weaponDesc.Name = "EllieWeaponDesc"
	weaponDesc.Properties = []objects.GCObjectProperty{
		objects.Uint32Prop("SlotType", uint32(EquipmentSlotWeapon)),
	}

	//TODO finish
	//weapon.Properties = []objects.GCObjectProperty{
	////	objects.Uint32Prop("ItemDesc.SlotType", 0x0a),
	//	objects.Uint32Prop("EquipmentSlot", uint32(EquipmentSlotWeapon)),
	//}

	armor := objects.NewGCObject("Armor")
	armor.GCType = "ScaleArmor1PAL.ScaleArmor1-1"
	armor.Name = "EllieArmour"

	//armor.Properties = []objects.GCObjectProperty{
	//	objects.Uint32Prop("Slot", uint32(EquipmentSlotTorso)),
	//}

	unitContainer := objects.NewUnitContainer(1, "EllieUnitContainer")
	//unitContainer.GCType = "unitcontainer"
	//unitContainer.Name = "EllieUnitContainer"

	baseInventory := objects.NewGCObject("Inventory")
	baseInventory.GCType = "avatar.base.Inventory"
	baseInventory.Name = "EllieBaseInventory"

	bankInventory := objects.NewGCObject("Inventory")
	bankInventory.GCType = "avatar.base.Bank"
	bankInventory.Name = "EllieBankInventory"

	//armorSlot := objects.NewGCObject("EquipmentSlot")
	//armorSlot.GCType = "avatar.base.Equipment.Description.Armor"
	//armorSlot.Name = "EquipmentSlot6"
	//
	//armorSlot.Properties = []objects.GCObjectProperty{
	//	objects.Uint32Prop("SlotID", uint32(EquipmentSlotTorso)),
	//}

	//armorSlotDesc := objects.NewGCObject("equipmentdesc")
	//armorSlotDesc.GCType = "avatar.base.Equipment.Description.Armor"
	//armorSlotDesc.Name = "EquipmentSlot6Desc"
	//
	//armorSlot.AddChild(armorSlotDesc)
	//armorSlot.AddChild(armor)
	//armor.AddChild(armorSlot)

	//avatarEquipment.AddChild(armor)

	//armorVisual := objects.NewGCObject("MountedVisual")
	//armorVisual.GCType = "avatar.races.humanfemale.HumanFemaleVisuals.ChainArmor1.Torso"
	//armorVisual.Name = "EllieArmorVisual"
	//
	//avatar.AddChild(armorVisual)

	slot := objects.NewGCObject("EquipmentSlot")
	slot.GCType = "avatar.base.Equipment.Description.PrimaryWeaponSlot"
	slot.Name = "EllieWeaponSlot"

	slot.Properties = []objects.GCObjectProperty{
		objects.Uint32Prop("SlotID", uint32(EquipmentSlotWeapon)),
		objects.Uint32Prop("SlotType", uint32(EquipmentSlotWeapon)),
		objects.Uint32Prop("DefaultItem", 0xAAAABBBB),
	}

	//0xbbc86ef4
	//slot6 := objects.NewGCObject("EquipmentSlot")
	//slot6.GCType = "avatar.base.Equipment.Description.Armor"
	//slot6.Name = "EllieArmorSlot"

	//slot6.Properties = []objects.GCObjectProperty{
	//	objects.Uint32Prop("SlotID", uint32(EquipmentSlotTorso)),
	//	objects.Uint32Prop("SlotType", uint32(EquipmentSlotTorso)),
	//	//objects.Uint32Prop("DefaultItem", 0x01),
	//}
	//

	//slot6.AddChild(armor)
	//slot.AddChild(weapon)
	//
	//avatarEquipment.AddChild(weapon)
	//avatar.AddChild(slot6)

	//avatar.AddChild(slot)
	//avatar.AddChild(slot6)

	//avatarDesc.Properties = []objects.GCObjectProperty{
	//	objects.StringProp("PVEStartSpawnPoint", "Start"),
	//}

	//worldEntityDesc := objects.NewGCObject("WorldEntityDesc")
	//worldEntityDesc.Name = "EllieWorldEntityDesc"

	unitBehaviour := objects.NewGCObject("UnitBehavior")
	unitBehaviour.GCType = "avatar.base.UnitBehavior"
	unitBehaviour.Name = "EllieBehaviour"

	unitContainer.AddChild(baseInventory)
	unitContainer.AddChild(bankInventory)

	avatarEquipment.AddChild(weapon)
	avatarEquipment.AddChild(armor)

	//avatar.AddChild(visual)
	//avatar.AddChild(rpgSettings)
	avatar.AddChild(avatarEquipment)
	avatar.AddChild(avatarSkills)
	avatar.AddChild(unitContainer)
	avatar.AddChild(unitBehaviour)
	avatar.AddChild(modifiers)
	avatar.AddChild(manipulators)
	avatar.AddChild(metrics)
	avatar.AddChild(dialogManager)
	//avatar.AddChild(avatarDesc)
	return avatar
}
