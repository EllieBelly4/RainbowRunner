package game

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/managers"
	"RainbowRunner/internal/objects"
	byter "RainbowRunner/pkg/byter"
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
		handleCharacterPlay(conn, reader)
	case CharacterGetList:
		handleCharacterList(conn)
	case CharacterCreate:
		handleCharacterCreate(conn, reader)
	default:
		return UnhandledChannelMessageError
	}

	return nil
}

func handleCharacterList(conn *RRConn) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.CharacterChannel)) // Character channel
	body.WriteByte(byte(CharacterGetList))          // Get character list (GotCharacter)

	count := len(conn.Client.Characters)

	body.WriteByte(byte(count))

	for _, character := range conn.Client.Characters {
		body.WriteUInt32(uint32(character.EntityProperties.ID)) // ID?
		sendPlayer(character, conn.Client, body)
	}

	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
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
	body.WriteByte(byte(messages.CharacterChannel)) // Character channel
	body.WriteByte(byte(CharacterCreate))
	body.WriteUInt32(0x01)

	body.WriteCString("Ellie")

	sendPlayer(loadPlayer(conn.Client), conn.Client, body)

	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
}

func handleCharacterPlay(conn *RRConn, reader *byter.Byter) {
	reader.UInt8()
	reader.UInt8()
	slot := reader.UInt8()
	conn.Client.CurrentCharacter = conn.Client.Characters[slot]

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.CharacterChannel))
	body.WriteByte(byte(CharacterPlay))
	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
}

func handleCharacterConnected(conn *RRConn) {
	count := 2
	conn.Client.Characters = make([]*objects.Player, 0, count)

	for i := 0; i < count; i++ {
		conn.Client.Characters = append(conn.Client.Characters, loadPlayer(conn.Client))
	}

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.CharacterChannel)) // Character channel
	body.WriteByte(byte(CharacterConnected))        // Connected
	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
}

func sendPlayer(character *objects.Player, client *RRConnClient, body *byter.Byter) {
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
	//gcAvatar := getAvatar()
	//avatar2 := getAvatar(0x02)

	//player := objects.NewPlayer("Ellie")
	//player := objects.NewGCObject("Player")
	//player.GCName = "Player Name"

	//player.AddChild(avatar)

	//slot6 := objects.NewGCObject("EquipmentSlot")
	//slot6.GCType = "avatar.base.Equipment.Description.Armor"
	//slot6.Name = "EllieArmorSlot"
	//
	//armor := objects.NewGCObject("Armor")
	//armor.GCType = "ChainArmor1PAL.ChainArmor1-1"
	//armor.Name = "EllieArmour"

	//slot6.AddChild(armor)
	//player.AddChild(slot6)

	//player := loadPlayer(conn.Client)
	avatar := loadAvatar(character)

	//player.AddChild(avatar)

	character.WriteFullGCObject(body)
	avatar.WriteFullGCObject(body)

	body.WriteByte(0x01)
	body.WriteByte(0x01)

	body.WriteCString("Normal")

	body.WriteByte(0x01)
	body.WriteByte(0x01)

	body.WriteUInt32(0x01)
}

func loadAvatar(player *objects.Player) *objects.GCObject {
	avatar := getAvatar(player.RREntityProperties().Conn)
	player.AddChild(avatar)

	managers.Entities.RegisterAll(player.EntityProperties.Conn, player)
	return avatar
}

func loadPlayer(client *RRConnClient) *objects.Player {
	player := objects.NewPlayer("Ellie")
	managers.Entities.RegisterAll(client, player)
	return player
}

func getAvatar(conn connections.Connection) *objects.GCObject {
	avatar := objects.NewGCObject("Avatar")
	avatar.GCType = "avatar.classes.FighterMale"
	//avatar.GCType = "avatar.base.avatar"
	avatar.GCName = "Avatar Name"
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
	modifiers.GCName = "Mod Name"
	modifiers.Properties = []objects.GCObjectProperty{
		objects.Uint32Prop("IDGenerator", 0x01),
	}

	manipulators := objects.NewGCObject("Manipulators")
	manipulators.GCName = "ManipulateMe"

	dialogManager := objects.NewGCObject("DialogManager")
	dialogManager.GCName = "EllieDialogManager"

	//animationList := objects.NewGCObject("AnimationList")
	//animationList.Name = "EllieAnimations"

	avatarSkills := objects.NewGCObject("Skills")
	avatarSkills.GCType = "avatar.base.skills"
	avatarSkills.GCName = "EllieSkills"

	//avatarDesc := objects.NewGCObject("AvatarDesc")
	//avatarDesc.GCType = "avatar.classes.fighterfemale.description"
	//avatarDesc.Name = "EllieAvatarDesc"

	avatarEquipment := objects.NewGCObject("Equipment")
	avatarEquipment.GCType = "avatar.base.Equipment"
	avatarEquipment.GCName = "EllieEquipment"

	//.text:0058E550     ; struct DFCClass *__thiscall Armor::getClass(Armor *__hidden this)
	//.text:0058E550 000 mov     eax, ?Class@Armor@@2PAVDFCClass@@A ; DFCClass * Armor::Class

	weapon := objects.NewGCObject("MeleeWeapon")
	weapon.GCType = "1HMace1PAL.1HMace1-1"
	weapon.GCName = "EllieWeapon"

	weaponDesc := objects.NewGCObject("MeleeWeaponDesc")
	weaponDesc.GCType = "1HMace1PAL.1HMace1-1.Description"
	weaponDesc.GCName = "EllieWeaponDesc"
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
	armor.GCName = "EllieArmour"

	//armor.Properties = []objects.GCObjectProperty{
	//	objects.Uint32Prop("Slot", uint32(EquipmentSlotTorso)),
	//}

	manipulator := objects.NewGCObject("Manipulator")
	//manipulator.GCType = "base.MeleeUnit.Manipulators.PrimaryWeapon"
	managers.Entities.RegisterAll(conn, manipulator)

	unitContainer := objects.NewUnitContainer(manipulator, "EllieUnitContainer")
	//unitContainer.GCType = "unitcontainer"
	//unitContainer.Name = "EllieUnitContainer"

	baseInventory := objects.NewGCObject("Inventory")
	baseInventory.GCType = "avatar.base.Inventory"
	baseInventory.GCName = "EllieBaseInventory"

	bankInventory := objects.NewGCObject("Inventory")
	bankInventory.GCType = "avatar.base.Bank"
	bankInventory.GCName = "EllieBankInventory"

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
	slot.GCName = "EllieWeaponSlot"

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
	unitBehaviour.GCName = "EllieBehaviour"

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
