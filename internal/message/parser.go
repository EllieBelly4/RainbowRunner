package message

import (
	"RainbowRunner/internal/byter"
	"RainbowRunner/internal/connection"
	"RainbowRunner/internal/crypt"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
)

type AuthMessageParser struct {
	connection *connection.Connection
}

func (p *AuthMessageParser) Parse(read []byte, count int) {
	var readBytes = 0

	for readBytes < count {
		packetLength := int(binary.LittleEndian.Uint16(read[readBytes:]))

		var decrypted = crypt.DecryptBlowfish(read[readBytes+2:readBytes+packetLength-2], packetLength-2)
		//var encrypted = crypt.EncryptBlowfish(decrypted, len(decrypted))
		//log.Info(fmt.Sprintf("Old:\n%s\nNew:\n%s\n", hex.Dump(read[readBytes+2:readBytes+packetLength]), hex.Dump(encrypted)))

		var byteReader = byter.NewLEByter(decrypted)

		if count-readBytes >= packetLength {
			p.processMessage(byteReader, packetLength)
			readBytes += packetLength
		} else {
			panic("Can't handle the split packets right now")
		}
	}
}

func (p *AuthMessageParser) processMessage(reader *byter.Byter, length int) {
	messageTypeID := reader.UInt8()

	log.Info(fmt.Sprintf(
		"Received %s (%d bytes):\n%s\n",
		AuthClientMessage(messageTypeID).String(), length, hex.Dump(reader.Buffer),
	))

	var err error

	switch AuthClientMessage(messageTypeID) {
	case AuthClientLoginPacket:
		err = HandleLoginMessage(p, reader)
	case AuthClientAboutToPlayPacket:
		err = HandleAboutToPlay(p, reader)
	case AuthClientServerListExtPacket:
		err = HandleServerListMessage(p, reader)
	}

	if err != nil {
		panic(err)
	}
}

func (p *AuthMessageParser) WriteAuthMessage(messageType AuthServerMessage, response *byter.Byter) error {
	sent, err := p.connection.WriteMessageBytes(append([]byte{byte(messageType)}, response.Buffer...))

	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Sent %s (%d bytes):\n%s\n", messageType.String(), len(sent), hex.Dump(sent)))

	return nil
}

func NewAuthMessageParser(conn net.Conn) *AuthMessageParser {
	return &AuthMessageParser{
		connection: connection.NewConnection(conn),
	}
}
