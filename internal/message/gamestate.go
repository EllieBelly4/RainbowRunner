package message

import (
	"RainbowRunner/internal/global"
	byter "RainbowRunner/pkg/byter"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func HandleAboutToPlay(p *AuthMessageParser, reader *byter.Byter) error {
	// SessionID
	reader.UInt64()
	serverID := reader.UInt8()

	log.Info(fmt.Sprintf("Wants to join server %d\n", serverID))

	response := byter.NewLEByter(make([]byte, 0, 0xFF))

	key := global.GenerateOneTimeKey()

	err := global.AddLoginRequest(key, p.Username)

	if err != nil {
		return err
	}

	response.WriteUInt32(key)        // OneTimeKey
	response.WriteUInt32(0x5678DEFA) // UID?
	response.WriteByte(serverID)     // Server ID
	p.WriteAuthMessage(AuthServerPlayOkPacket, response)

	//response.WriteUInt32(0x7F000001) // IP 127.0.0.1
	//response.WriteUInt32(0x00000A2B) // Port 2603
	//response.WriteUInt32(uint32(serverID))
	//response.WriteUInt32(0xDEADBEEF) // OneTimeKey
	//
	//p.WriteAuthMessage(AuthServerHandoffToGamePacket, response)

	return nil
}
