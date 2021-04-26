package message

import (
	"RainbowRunner/internal/byter"
	"RainbowRunner/internal/connection"
	"RainbowRunner/internal/crypt"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
)

type Parser struct {
	connection *connection.Connection
}

func (p *Parser) Parse(read []byte, count int) {
	var readBytes = 0

	for readBytes < count {
		packetLength := int(binary.LittleEndian.Uint16(read[readBytes:]))

		var decrypted = crypt.DecryptBlowfish(read[readBytes+2:readBytes+packetLength-2], packetLength-2)
		//var encrypted = crypt.EncryptBlowfish(decrypted, len(decrypted))
		//fmt.Printf("Old:\n%s\nNew:\n%s\n", hex.Dump(read[readBytes+2:readBytes+packetLength]), hex.Dump(encrypted))

		var byteReader = byter.NewByter(decrypted)
		byteReader.LittleEndian()

		if count-readBytes >= packetLength {
			p.processMessage(byteReader)
			readBytes += packetLength
		} else {
			panic("Can't handle the split packets right now")
		}
	}
}

func (p *Parser) processMessage(reader *byter.Byter) {
	fmt.Printf("Read:\n%s\n", hex.Dump(reader.Buffer))

	messageTypeID := reader.UInt8()

	var err error

	switch int(messageTypeID) {
	case 0:
		err = HandleLoginMessage(p.connection, reader)
	case 5:
		err = HandleServerListMessage(p.connection, reader)
	}

	if err != nil {
		panic(err)
	}
}

func NewParser(conn net.Conn) *Parser {
	return &Parser{
		connection: connection.NewConnection(conn),
	}
}
