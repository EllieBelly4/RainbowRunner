package game

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	"RainbowRunner/internal/objects"
	"RainbowRunner/internal/serverconfig"
	byter "RainbowRunner/pkg/byter"
	"RainbowRunner/pkg/events"
	log "github.com/sirupsen/logrus"
)

type ZoneChannelMessage byte

const (
	ZoneUnk0 ZoneChannelMessage = iota
	ZoneUnk1
	ZoneUnk2
	ZoneUnk3
	ZoneUnk4
	ZoneUnk5
	ZoneJoin
	ZoneUnk7
	ZoneUnk8
)

func handleZoneChannelMessages(conn *connections.RRConn, msgSubType uint8, reader *byter.Byter) error {
	switch ZoneChannelMessage(msgSubType) {
	case ZoneJoin:
		handleZoneJoin(conn)
	case ZoneUnk8:
		body := byter.NewLEByter(make([]byte, 0, 1024))
		body.WriteByte(byte(messages.ZoneChannel))
		body.WriteByte(0x08)
		connections.WriteCompressedA(conn, 0x01, 0x0f, body)
	default:
		return UnhandledChannelMessageError
	}
	return nil
}

func handleZoneJoin(conn *connections.RRConn) {
	player := objects.Players.GetPlayer(uint16(conn.GetID()))

	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(byte(messages.ZoneMessageReady))
	//body.WriteByte(0x02) // Other acceptable values
	//body.WriteByte(0x05) // Other acceptable values
	body.WriteUInt32(player.Zone().ID) // World ID

	// MiniMapExplored::ReadExploredBits
	// dungeon00_level01 - 0x12
	exploredBitCount := uint16(0x12)
	body.WriteUInt16(exploredBitCount)

	for i := 0; i < int(exploredBitCount); i++ {
		body.WriteUInt32(0xFFFFFFFF)
	}

	connections.WriteCompressedA(conn, 0x01, 0x0f, body)

	body = byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.ZoneChannel))
	body.WriteByte(byte(messages.ZoneMessageInstanceCount))

	// Adds two separate values into the ZoneClient
	body.WriteUInt32(0x01)
	body.WriteUInt32(0x01)
	connections.WriteCompressedA(conn, 0x01, 0x0f, body)

	SendInterval(conn)

	player.CurrentCharacter.SendCreateNewPlayerEntity(player)
	player.CurrentCharacter.OnZoneJoin()

	events.Emit(objects.PlayerEnteredZoneEvent{
		Player: player.CurrentCharacter,
		Zone:   player.CurrentCharacter.Zone,
	})

	//	WriteCompressedA(conn, 0x01, 0x0f, NewLEByterFromCommandString(`
	//07
	//# ClientEntityManager::processInterval
	//# This seems to log messages about the pathManager budget
	//# I don't think this is meant to sync each tick
	//0D # ID
	//
	//# ClientEntityManager::processInterval
	//01 00 00 00
	//01 00 00 00
	//01 00 00 00
	//
	//# PathManager::ReadBudget
	//# These values are syncing the path budget
	//FF FF FF FF
	//FF FF # Per Update (idk)
	//FF FF # Per Path (idk)
	//
	//06 #end`))

	avatar := player.CurrentCharacter.GetChildByGCNativeType("Avatar").(*objects.Avatar)

	avatar.SendFollowClient()
	avatar.IsSpawned = true

	if serverconfig.Config.Welcome.SendWelcomeMessage {
		SendWelcomeMessage(player)
	}
}

func sendGoToZone(conn *connections.RRConn, zoneName string) {
	rrPlayer := objects.Players.Players[conn.GetID()]

	tZone := objects.Zones.GetOrCreateZone(zoneName)

	if tZone == nil {
		log.Errorf("could not find zone %s", zoneName)
		return
	}

	rrPlayer.CurrentCharacter.JoinZone(tZone)
}

func SendWelcomeMessage(player *objects.RRPlayer) {
	msg := messages.ChatMessage{
		Channel: messages.MessageChannelSourceGlobalAnnouncement,
		Message: serverconfig.Config.Welcome.Message,
	}

	player.Conn.SendMessage(msg)
}
