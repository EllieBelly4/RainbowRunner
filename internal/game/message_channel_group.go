package game

import (
	"RainbowRunner/internal/connections"
	"RainbowRunner/internal/game/messages"
	byter "RainbowRunner/pkg/byter"
)

type GroupChannelMessage byte

const (
	GroupConnected GroupChannelMessage = iota
)

func handleGroupChannelMessages(conn *connections.RRConn, msgType byte, reader *byter.Byter) error {
	switch GroupChannelMessage(msgType) {
	case GroupConnected:
		handleGroupConnected(conn)
	default:
		return UnhandledChannelMessageError
	}

	return nil
}

func handleGroupConnected(conn *connections.RRConn) {
	body := byter.NewLEByter(make([]byte, 0, 1024))
	body.WriteByte(byte(messages.GroupChannel))
	body.WriteByte(48)

	//sendGoToZone(conn, body, "TestTilesets")
	//sendGoToZone(conn, body, "thehub")
	//sendGoToZone(conn, body, "dungeon01_level01")
	//sendGoToZone(conn, body, "d06_l01_q05")//Sitar Hero-Easy(2)
	//sendGoToZone(conn, body, "d06_l07_q05")//Sitar Hero-Expert(2)
	//sendGoToZone(conn, body, "epic01_central") //Sitar Hero-Expert(2)

	//sendGoToZone(conn, body, "dungeon02_level00")//Algor's Terror-Dome Base(2)
	//sendGoToZone(conn, body, "dungeon03_level00")//Vexation Station(2)
	//sendGoToZone(conn, body, "dungeon04_level00")//The Widower's Nest Base Camp(2)
	//sendGoToZone(conn, body, "dungeon05_level00")//Frump's Interdungenal Resort(2)
	//sendGoToZone(conn, body, "dungeon06_level00")//Mutanous Malaise Mezzanine(2)
	//sendGoToZone(conn, body, "dungeon09_level00")//The Shadows Gate(2)
	//sendGoToZone(conn, body, "dungeon11_level00")//Ballzack's Base Camp(2)
	//sendGoToZone(conn, body, "dungeon15_level00")//Ratsputin's Holding Cell(2)
	//sendGoToZone(conn, body, "Tutorial")
	//sendGoToZone(conn, body, "TestVendorLevelSpecArmor")
	//sendGoToZone(conn, body, "TheHubPortals_Dungeon01")
	//sendGoToZone(conn, body, "dungeon08_level00")//The Embercore Portal(2)
	//sendGoToZone(conn, body, "pvp_start")//Pwnston
	//sendGoToZone(conn, body, "dungeon02_level08_boss")//Lair of the Snow Chieftain (2)

	//sendGoToZone(conn, body, "dungeon16_level00")//The Mutantmania Training Center (2)
	sendGoToZone(conn, body, "town") //townston

}
