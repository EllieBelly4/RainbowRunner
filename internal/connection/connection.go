package connection

import (
	"RainbowRunner/internal/byter"
	"RainbowRunner/internal/crypt"
	"encoding/binary"
	"net"
)

type Connection struct {
	Conn net.Conn
}

func (c *Connection) WriteMessageBytes(bytes []byte) ([]byte, error) {
	length := len(bytes)
	remainder := length % 8

	if remainder != 0 {
		padding := 8 - remainder
		bytes = append(bytes, make([]byte, padding)...)
		length += padding
	}

	packetLength := length + 2 + 8

	_, err := c.Conn.Write([]byte{
		byte(packetLength),
		byte(packetLength >> 8),
	})

	if err != nil {
		return nil, err
	}

	checksum := uint32(0)

	for i := 0; i < length; i += 4 {
		value := binary.LittleEndian.Uint32(bytes[i:])
		checksum ^= value
	}

	checksumBytes := make([]byte, 4)

	binary.LittleEndian.PutUint32(checksumBytes, checksum)

	bytes = append(bytes, checksumBytes...)
	bytes = append(bytes, []byte{0x0, 0x0, 0x0, 0x0}...)

	//written, err := c.Conn.Write(bytes)
	encryptedPayload := crypt.EncryptBlowfish(bytes, len(bytes))
	_, err = c.Conn.Write(encryptedPayload)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func (c *Connection) WriteMessageByter(response *byter.Byter) ([]byte, error) {
	return c.WriteMessageBytes(response.Buffer)
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{
		Conn: conn,
	}
}
