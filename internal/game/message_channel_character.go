package game

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/objects"
	"RainbowRunner/internal/types/drobjecttypes"
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
		body.WriteUInt32(character.EntityProperties.ID) // ID?
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

	body.WriteCString(conn.LoginName)

	sendPlayer(loadPlayer(conn.LoginName), conn.Client, body)

	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
}

func handleCharacterPlay(conn *connections.RRConn, reader *byter.Byter) {
	reader.UInt8()
	reader.UInt8()
	slot := reader.UInt8()

	character := objects.Players.Players[conn.GetID()].Characters[slot]
	objects.Players.Players[conn.GetID()].CurrentCharacter = character

	character.WalkChildren(func(object drobjecttypes.DRObject) {
		props := object.(objects.IRREntityPropertiesHaver).GetRREntityProperties()

		props.Conn = conn
		props.OwnerID = uint16(conn.GetID())
	})

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.CharacterChannel))
	body.WriteByte(byte(CharacterPlay))
	connections.WriteCompressedA(conn, 0x01, 0x0f, body)
}

func handleCharacterConnected(conn *connections.RRConn) {
	count := 2
	objects.Players.Players[conn.GetID()].Characters = make([]*objects.Player, 0, count)

	for i := 0; i < count; i++ {
		player := loadPlayer(conn.LoginName)
		player.EntityProperties.Conn = conn
		player.EntityProperties.ID = uint32(objects.NewID())
		//player.EntityProperties.ID = uint32(i + 1)
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
	avatar := objects.LoadAvatar()
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

func loadPlayer(name string) *objects.Player {
	player := objects.NewPlayer(name)
	//objects.Entities.RegisterAll(client, player.Children()...)
	return player
}
