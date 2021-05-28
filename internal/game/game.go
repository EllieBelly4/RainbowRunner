package game

import (
	"RainbowRunner/internal/byter"
	"RainbowRunner/internal/objects"
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

type Channel byte

const (
	NoChannel Channel = iota
	Unk1
	Unk2
	UserChannel
	CharacterChannel
	Unk5
	ChatChannel
	ClientEntityChannel
	Unk8
	GroupChannel
	TradeChannel
	UnkB
	UnkC
	ZoneChannel
	UnkE
	PosseChannel
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

type GroupChannelMessage byte

const (
	GroupConnected GroupChannelMessage = iota
)

type ZoneChannelMessage byte

const (
	ZoneUnk0 ZoneChannelMessage = iota
	ZoneUnk1
	ZoneUnk2
	ZoneUnk3
	ZoneUnk4
	ZoneUnk5
	ZoneUnk6
)

func StartGameServer() {
	listen, err := net.Listen("tcp", "0.0.0.0:2603")

	if err != nil {
		panic(err)
	}

	defer func() {
		err := listen.Close()
		if err != nil {
			panic(err)
		}
	}()

	for {
		conn, err := listen.Accept()

		if err != nil {
			panic(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	//parser := message.NewParser(conn)
	buf := make([]byte, 1024*10)

	fmt.Println("Client connected to gameserver")
	var clientID uint32 = 0

	for {
		read, err := conn.Read(buf)

		if err != nil {
			log.Info(fmt.Sprintf("failed to read from connection: %e\n", err))
			break
		}

		log.Info(fmt.Sprintf("(GameServer)Received: \n%s\n", hex.Dump(buf[0:read])))

		reader := byter.NewLEByter(buf[0:read])

		msgType := reader.UInt8() // Message Type?

		if msgType == 0x0a {
			clientID = reader.UInt24() // Unk
			reader.UInt32()            // Packet Length
			reader.UInt8()
			msgTypeA := reader.UInt8()
			reader.UInt8()
			reader = ReadCompressedA(reader)

			log.Infof("Uncompressed A:\n%s", hex.Dump(reader.Buffer))

			if msgTypeA == 0x00 {
				reader.UInt8()      // Some type?
				_ = reader.UInt32() // One Time Key
				reader.Bytes(1)     // Null

				body := byter.NewLEByter(make([]byte, 0, 1024))

				body.WriteByte(0x03)
				WriteMessage(0x10, clientID, 0x0a, conn, body)

				body = byter.NewLEByter(make([]byte, 0, 1024))
				// Unk
				body.WriteUInt24(0xB2B3B4)
				// Unk
				body.WriteByte(0x00)
				WriteCompressedA(clientID, 0x00, 0x03, body, conn)
			} else if msgTypeA == 0x02 {
			} else {
				log.Panicf("Unhandled 0x0a message type %x", msgTypeA)
			}
		} else if msgType == 0x06 || msgType == 0x0e {
			if msgType == 0x0e {
				first := true
				// TODO when we can reliably get the length of all packets we can ensure this is done for everything
				for reader.HasRemainingData() {
					if !first {
						reader.UInt8()
					}
					msgReader := ReadCompressedE(reader)

					log.Infof("Uncompressed E:\n%s", hex.Dump(msgReader.Buffer))

					handleChannelMessage(conn, msgReader, clientID)
					first = false
				}
			} else {
				reader.UInt24() // Unk
				reader.UInt24() // Size
				reader.UInt8()
				reader.UInt24() // Client ID
				reader.UInt8()  // Channel?
				reader.UInt8()  // Sub type?
				reader.UInt24() // Unk

				handleChannelMessage(conn, reader, clientID)
			}
		} else {
			log.Info(fmt.Sprintf("Unhandled message type %x\n", msgType))
		}
	}
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
	avatar := getAvatar(0x01)
	//avatar2 := getAvatar(0x02)

	player := objects.NewGCObject("Player")
	player.ID = 0xBABAFAAB
	player.Name = "Player Name"

	player.AddChild(avatar)
	//player.AddChild(avatar2)
	player.Properties = []objects.GCObjectProperty{
		objects.StringProp("Name", "Ellie"),
	}

	player.Serialise(body)

	body.WriteCString("Unk")  // Specific to player::readObject
	body.WriteCString("Unk2") // Specific to player::readObject
	body.WriteCString("Unk3") // Specific to player::readObject
	body.WriteCString("Unk4") // Specific to player::readObject
	body.WriteUInt32(0x01)    // Specific to player::readObject
	body.WriteUInt32(0x01)    // Specific to player::readObject

	avatar.Serialise(body)

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
	modifiers.ID = 0xBABAFAAC
	modifiers.Name = "Mod Name"
	//modifiers.Properties = []objects.GCObjectProperty{
	//	objects.Uint32Prop("IDGenerator", 0x01),
	//}

	manipulators := objects.NewGCObject("Manipulators")
	manipulators.ID = 0xBABAFACC
	manipulators.Name = "ManipulateMe"

	//animationList := objects.NewGCObject("AnimationList")
	//animationList.ID = 0xBABAF00E
	//animationList.Name = "EllieAnimations"

	avatarSkills := objects.NewGCObject("Skills")
	avatarSkills.GCType = "avatar.base.skills"
	avatarSkills.ID = 0xBAADBABA
	avatarSkills.Name = "EllieSkills"

	avatarDesc := objects.NewGCObject("AvatarDesc")
	avatarDesc.GCType = "avatar.classes.fighterfemale.description"
	avatarDesc.ID = 0xBABAF00D
	avatarDesc.Name = "EllieAvatarDesc"

	//worldEntityDesc := objects.NewGCObject("WorldEntityDesc")
	//worldEntityDesc.ID = 0xBABABABA
	//worldEntityDesc.Name = "EllieWorldEntityDesc"

	unitBehaviour := objects.NewGCObject("UnitBehavior")
	unitBehaviour.GCType = "avatar.base.UnitBehavior"
	unitBehaviour.ID = 0xBABAF00F
	unitBehaviour.Name = "EllieBehaviour"

	//avatar.AddChild(visual)
	//avatar.AddChild(rpgSettings)
	avatar.AddChild(modifiers)
	avatar.AddChild(avatarSkills)
	avatar.AddChild(manipulators)
	avatar.AddChild(avatarDesc)
	avatar.AddChild(unitBehaviour)
	return avatar
}

func sendCharacterList(conn net.Conn, clientID uint32) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(CharacterChannel)) // Character channel
	body.WriteByte(byte(CharacterGetList)) // Get character list (GotCharacter)

	count := 0x01

	body.WriteByte(byte(count))

	for i := 0; i < count; i++ {
		body.WriteUInt32(uint32(i + 1)) // ID?
		sendPlayer(body)
	}

	WriteCompressedA(clientID, 0x01, 0x0f, body, conn)
}
