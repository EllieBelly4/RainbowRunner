package game

import (
	"RainbowRunner/internal/config"
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

	//sendGoToZone(conn, "TestTilesets")
	//sendGoToZone(conn, "thehub")
	//sendGoToZone(conn, "dungeon01_level01")//Algeron(2)
	//sendGoToZone(conn, "d06_l01_q05")//Sitar Hero-Easy(2)
	//sendGoToZone(conn, "d06_l07_q05")//Sitar Hero-Expert(2)
	//sendGoToZone(conn, "epic01_central") //Hundred Town(2)
	//sendGoToZone(conn, "dungeon02_level00")//Algor's Terror-Dome Base(2)
	//sendGoToZone(conn, "dungeon03_level00")//Vexation Station(2)
	//sendGoToZone(conn, "dungeon04_level00")//The Widower's Nest Base Camp(2)
	//sendGoToZone(conn, "dungeon05_level00")//Frump's Interdungenal Resort(2)
	//sendGoToZone(conn, "dungeon06_level00")//Mutanous Malaise Mezzanine(2)
	//sendGoToZone(conn, "dungeon09_level00")//The Shadows Gate(2)
	//sendGoToZone(conn, "dungeon11_level00")//Ballzack's Base Camp(2)
	//sendGoToZone(conn, "dungeon15_level00")//Ratsputin's Holding Cell(2)
	//sendGoToZone(conn, "Tutorial")
	//sendGoToZone(conn, "TestVendorLevelSpecArmor")
	//sendGoToZone(conn, "TheHubPortals_Dungeon01")//d01_Algeron(2)
	//sendGoToZone(conn, "dungeon08_level00")//The Embercore Portal(2)
	//sendGoToZone(conn, "pvp_start")//Pwnston
	//sendGoToZone(conn, "dungeon02_level08_boss")//Lair of the Snow Chieftain (2)
	//sendGoToZone(conn, "dungeon16_level00")//The Mutantmania Training Center (2)
	//sendGoToZone(conn, "dungeon00_level01") //DewValley Forest(2)
	//sendGoToZone(conn, "dungeon00_level02") //Dew Valley Forest - Level 2 (2)
	//sendGoToZone(conn, "dungeon00_level03") //Dew Valley Forest - Level 3(2)
	//sendGoToZone(conn, "dungeon00_level03_boss") //Rattle Tooth's Lair
	//sendGoToZone(conn, "dungeon01_level01_off1a") //Orok Outpost
	//sendGoToZone(conn, "dungeon01_level02") //Algernon - Level 2(2)
	//sendGoToZone(conn, "dungeon01_level03") //Algernon - Level 3(2)
	//sendGoToZone(conn, "dungeon01_level03_off1a") //Capwn's Family Ratstaurant(2)
	//sendGoToZone(conn, "dungeon01_level04") //Algernon - Level 4(2)
	//sendGoToZone(conn, "dungeon01_level04_off1a") //Not the Rumored Refinery(2)
	//sendGoToZone(conn, "dungeon01_level05") //Algernon - Level 5(2)
	//sendGoToZone(conn, "dungeon01_level05_off1a") //Not the Mutant Operation Center
	//sendGoToZone(conn, "dungeon01_level06") //Algernon - Level 6
	//sendGoToZone(conn, "dungeon01_level07") //townston
	//sendGoToZone(conn, "dungeon01_level07_off1a") //Little Dew Valley Forest
	//sendGoToZone(conn, "dungeon01_level07_off2a") //Porthole Laboratories
	//sendGoToZone(conn, "dungeon01_level08_boss") //Sissirat's Lair
	//sendGoToZone(conn, "dungeon02_level01") //Algor's Terror-Dome
	//sendGoToZone(conn, "dungeon02_level01_off1a") //Subterrorania
	//sendGoToZone(conn, "dungeon02_level02") //Algor's Terror-Dome - Level 2
	//sendGoToZone(conn, "elite01_stage06_level02") //Horrific Dungeon of Legend - Stage 6 - Lower Cindercore
	//sendGoToZone(conn, "town") //townston
	//sendGoToZone(conn, "town") //townston
	//sendGoToZone(conn, "town") //townston
	//sendGoToZone(conn, "town") //townston
	//sendGoToZone(conn, "town") //townston
	//sendGoToZone(conn, "town") //townston
	//sendGoToZone(conn, "town") //townston
	//sendGoToZone(conn, "town") //townston
	//sendGoToZone(conn, "town") //townston
	//sendGoToZone(conn, "town") //townston

	sendGoToZone(conn, config.Config.DefaultZone)

}
