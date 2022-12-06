package game

import (
	"RainbowRunner/internal/config"
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/database"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/objects"
	byter "RainbowRunner/pkg/byter"
	"fmt"
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

func handleCharacterChannelMessages(conn *connections.RRConn, msgType byte, reader *byter.Byter) error {
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

func handleCharacterList(conn *connections.RRConn) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.CharacterChannel)) // Character channel
	body.WriteByte(byte(CharacterGetList))          // Get character list (GotCharacter)

	count := len(objects.Players.Players[conn.GetID()].Characters)

	body.WriteByte(byte(count))

	for _, character := range objects.Players.Players[conn.GetID()].Characters {
		body.WriteUInt32(uint32(character.EntityProperties.ID)) // ID?
		sendPlayer(character, conn.Client, body)
	}

	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
}

func handleCharacterCreate(conn *connections.RRConn, reader *byter.Byter) {
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

func handleCharacterPlay(conn *connections.RRConn, reader *byter.Byter) {
	reader.UInt8()
	reader.UInt8()
	slot := reader.UInt8()
	objects.Players.Players[conn.GetID()].CurrentCharacter = objects.Players.Players[conn.GetID()].Characters[slot]

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.CharacterChannel))
	body.WriteByte(byte(CharacterPlay))
	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
}

func handleCharacterConnected(conn *connections.RRConn) {
	count := 2
	objects.Players.Players[conn.GetID()].Characters = make([]*objects.Player, 0, count)

	for i := 0; i < count; i++ {
		player := loadPlayer(conn.Client)
		player.EntityProperties.Conn = conn
		player.EntityProperties.ID = uint32(i + 1)
		player.EntityProperties.OwnerID = uint16(conn.GetID())
		objects.Players.Players[conn.GetID()].Characters = append(objects.Players.Players[conn.GetID()].Characters, player)
	}

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.CharacterChannel)) // Character channel
	body.WriteByte(byte(CharacterConnected))        // Connected
	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
}

func sendPlayer(character *objects.Player, client *connections.RRConnClient, body *byter.Byter) {
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
	//player.GCLabel = "Player Name"

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

	character.AddChild(avatar)

	//avatar2 := loadAvatar(character)
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

func loadAvatar(player *objects.Player) *objects.Avatar {
	avatar := getAvatar(player.RREntityProperties().Conn)

	objects.Entities.RegisterAll(player.EntityProperties.Conn, avatar)
	return avatar
}

func loadPlayer(client *connections.RRConnClient) *objects.Player {
	player := objects.NewPlayer("Ellie")
	objects.Entities.RegisterAll(client, player.Children()...)
	return player
}

func getAvatar(conn connections.Connection) *objects.Avatar {
	avatar := objects.NewAvatar("avatar.classes.FighterFemale")
	avatar.GCLabel = "Avatar Name"
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
	modifiers.GCLabel = "Mod Name"
	modifiers.Properties = []objects.GCObjectProperty{
		objects.Uint32Prop("IDGenerator", 0x01),
	}

	dialogManager := objects.NewGCObject("DialogManager")
	dialogManager.GCLabel = "EllieDialogManager"

	//animationList := objects.NewGCObject("AnimationList")
	//animationList.Name = "EllieAnimations"

	avatarSkills := objects.NewGCObject("Skills")
	avatarSkills.GCType = "avatar.base.skills"
	avatarSkills.GCLabel = "EllieSkills"

	//avatarDesc := objects.NewGCObject("AvatarDesc")
	//avatarDesc.GCType = "avatar.classes.fighterfemale.description"
	//avatarDesc.Name = "EllieAvatarDesc"

	//avatarEquipment := objects.NewGCObject("Equipment")
	//avatarEquipment.GCType = "avatar.base.Equipment"
	//avatarEquipment.GCLabel = "EllieEquipment"

	avatarEquipment := objects.NewInventoryEquipment("avatar.base.Equipment", avatar)
	avatarEquipment.GCLabel = "EllieEquipment"

	//.text:0058E550     ; struct DFCClass *__thiscall Armor::getClass(Armor *__hidden this)
	//.text:0058E550 000 mov     eax, ?Class@Armor@@2PAVDFCClass@@A ; DFCClass * Armor::Class

	//weapon := objects.NewGCObject("MeleeWeapon")
	//weapon.GCType = "1HSwordMythicPAL.1HSwordMythic6"
	//weapon.GCLabel = "EllieWeapon"

	//weaponDesc := objects.NewGCObject("MeleeWeaponDesc")
	//weaponDesc.GCType = "1HMace1PAL.1HMace1-1.Description"
	//weaponDesc.GCLabel = "EllieWeaponDesc"
	//weaponDesc.Properties = []objects.GCObjectProperty{
	//	objects.Uint32Prop("SlotType", uint32(objects.EquipmentSlotWeapon)),
	//}

	//TODO finish
	//weapon.Properties = []objects.GCObjectProperty{
	////	objects.Uint32Prop("ItemDesc.SlotType", 0x0a),
	//	objects.Uint32Prop("EquipmentSlot", uint32(EquipmentSlotWeapon)),
	//}

	//armor.Properties = []objects.GCObjectProperty{
	//	objects.Uint32Prop("Slot", uint32(EquipmentSlotTorso)),
	//}

	manipulator := objects.NewGCObject("Manipulator")
	//manipulator.GCType = "base.MeleeUnit.Manipulators.PrimaryWeapon"
	objects.Entities.RegisterAll(conn, manipulator)

	unitContainer := objects.NewUnitContainer(manipulator, "EllieUnitContainer", avatar)
	//unitContainer.GCType = "unitcontainer"
	//unitContainer.Name = "EllieUnitContainer"

	baseInventory := objects.NewInventory("avatar.base.Inventory", 11)
	baseInventory.GCLabel = "EllieBaseInventory"

	bankInventory := objects.NewInventory("avatar.base.Bank", 12)
	bankInventory.GCLabel = "EllieBankInventory"

	tradeInventory := objects.NewInventory("avatar.base.TradeInventory", 13)
	tradeInventory.GCLabel = "EllieTradeInventory"

	manipulators := objects.NewManipulators("Manipulators")
	manipulators.GCLabel = "ManipulateMe"

	// Items in inventories
	//randomItem := objects.NewEquipment("PlateMythicPAL.PlateMythicBoots1", "PlateMythicPAL.PlateMythicBoots1.Mod1", objects.ItemArmour, types.EquipmentSlotFoot)
	//baseInventory.AddChild(randomItem)

	//r := rand.New(rand.NewSource(time.Now().Unix()))

	AddEquipment(avatarEquipment, manipulators,
		"PlateArmor3PAL.PlateArmor3-7",
		//objects.ArmourMap["armor"][int(r.Int63())%len(objects.ArmourMap["armor"])],
		"PlateBoots3PAL.PlateBoots3-7",
		//objects.ArmourMap["boots"][int(r.Int63())%len(objects.ArmourMap["boots"])],
		"PlateHelm3PAL.PlateHelm3-7",
		//objects.ArmourMap["helm"][int(r.Int63())%len(objects.ArmourMap["helm"])],
		"PlateGloves3PAL.PlateGloves3-7",
		//objects.ArmourMap["gloves"][int(r.Int63())%len(objects.ArmourMap["gloves"])],
		"CrystalMythicPAL.CrystalMythicShield1",
	)

	//manipulators.AddChild(armor)
	//manipulators.AddChild(weapon)

	//slot := objects.NewGCObject("EquipmentSlot")
	//slot.GCType = "avatar.base.Equipment.Description.PrimaryWeaponSlot"
	//slot.GCLabel = "EllieWeaponSlot"
	//
	//slot.Properties = []objects.GCObjectProperty{
	//	objects.Uint32Prop("SlotID", uint32(types.EquipmentSlotWeapon)),
	//	objects.Uint32Prop("SlotType", uint32(types.EquipmentSlotWeapon)),
	//	objects.Uint32Prop("DefaultItem", 0xAAAABBBB),
	//}

	//avatarDesc.Properties = []objects.GCObjectProperty{
	//	objects.StringProp("PVEStartSpawnPoint", "Start"),
	//}

	//worldEntityDesc := objects.NewGCObject("WorldEntityDesc")
	//worldEntityDesc.Name = "EllieWorldEntityDesc"

	unitBehaviour := objects.NewUnitBehavior("avatar.base.UnitBehavior")
	unitBehaviour.GCLabel = "EllieBehaviour"

	unitContainer.AddChild(baseInventory)
	unitContainer.AddChild(bankInventory)
	unitContainer.AddChild(tradeInventory)

	//avatarEquipment.AddChild(weapon)
	//avatarEquipment.AddChild(armor)

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

func AddEquipment(equipment objects.DRObject, manipulators *objects.Manipulators, armour string, boots string, helm string, gloves string, shield string) {
	randomArmour := objects.AddRandomEquipment(database.Armours, objects.ItemArmour)
	randomArmour.GCLabel = "EllieArmour"

	randomBoots := objects.AddRandomEquipment(database.Boots, objects.ItemArmour)
	randomBoots.GCLabel = "EllieArmour"

	randomHelm := objects.AddRandomEquipment(database.Helmets, objects.ItemArmour)
	randomHelm.GCLabel = "EllieArmour"

	randomGloves := objects.AddRandomEquipment(database.Gloves, objects.ItemArmour)
	randomGloves.GCLabel = "EllieArmour"

	randomWeapon := objects.AddRandomEquipment(database.MeleeWeapons, objects.ItemMeleeWeapon)
	randomWeapon.GCLabel = "EllieWeapon"

	//if len(shield) > 0 {
	//	randomShield := objects.NewEquipment(gloves, "ScaleModPAL.Rare.Mod1", "Armor", types.EquipmentSlotOffhand)
	//	randomShield.GCType = shield
	//	randomShield.GCLabel = "EllieArmour"
	//
	//	equipment.AddChild(randomShield)
	//	manipulators.AddChild(randomShield)
	//}

	//randomArmour = objects.NewEquipment(
	//	//"items.deprecated.DeprecatedChildArmorPAL.Body014",
	//	"items.pal.magebodypal.partialbuiltuniqueseasonal001",
	//	"ScaleModPAL.Rare.Mod1",
	//	objects.ItemArmour,
	//	types.EquipmentSlotTorso,
	//)
	//
	//randomHelm = objects.NewEquipment(
	//	"items.pal.rangerhelmpal.partialbuiltmythicseasonal002",
	//	"ScaleModPAL.Rare.Mod1",
	//	objects.ItemArmour,
	//	types.EquipmentSlotHead,
	//)
	//
	//randomGloves = objects.NewEquipment(
	//	"items.pal.mageglovespal.partialbuiltmythicseasonal001",
	//	"ScaleModPAL.Rare.Mod1",
	//	objects.ItemArmour,
	//	types.EquipmentSlotHand,
	//)
	//
	//randomBoots = objects.NewEquipment(
	//	"items.pal.magebootspal.partialbuiltmythicseasonal001",
	//	"ScaleModPAL.Rare.Mod1",
	//	objects.ItemArmour,
	//	types.EquipmentSlotFoot,
	//)
	//
	//randomWeapon = objects.NewEquipment(
	//	"items.pal.1hswordpal.mythicprebuiltboss001",
	//	"ScaleModPAL.Rare.Mod1",
	//	objects.ItemMeleeWeapon,
	//	types.EquipmentSlotWeapon,
	//)

	equipment.AddChild(randomGloves)
	equipment.AddChild(randomBoots)
	equipment.AddChild(randomHelm)
	equipment.AddChild(randomArmour)
	equipment.AddChild(randomWeapon)

	manipulators.AddChild(randomGloves)
	manipulators.AddChild(randomBoots)
	manipulators.AddChild(randomHelm)
	manipulators.AddChild(randomArmour)
	manipulators.AddChild(randomWeapon)

	if config.Config.Logging.LogRandomEquipment {
		fmt.Printf(`Random equipment for today is:
Helm: %s
Armour: %s
Gloves: %s
Boots: %s
Weapon: %s
`, randomHelm.GCType, randomArmour.GCType, randomGloves.GCType, randomBoots.GCType, randomWeapon.GCType)
	}
}
