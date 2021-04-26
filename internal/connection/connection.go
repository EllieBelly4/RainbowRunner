package connection

import (
	"RainbowRunner/internal/byter"
	"RainbowRunner/internal/crypt"
	"encoding/binary"
	"fmt"
	"net"
)

type Connection struct {
	Conn net.Conn
}

func (c *Connection) WriteMessageBytes(bytes []byte) error {
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

	fmt.Printf("%x", []byte{
		byte(packetLength),
		byte(packetLength >> 8),
	})

	if err != nil {
		return err
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

	fmt.Printf("Checksum: %x\n", checksum)

	//written, err := c.Conn.Write(bytes)
	encryptedPayload := crypt.EncryptBlowfish(bytes, len(bytes))
	written, err := c.Conn.Write(encryptedPayload)

	if err != nil {
		return err
	}

	fmt.Printf("%x\n", bytes)
	fmt.Printf("%x\n", encryptedPayload)

	fmt.Printf("Send message with payload length %d\n", written)

	return nil
}

func (c *Connection) WriteMessageByter(response *byter.Byter) error {
	return c.WriteMessageBytes(response.Buffer)
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{
		Conn: conn,
	}
}
