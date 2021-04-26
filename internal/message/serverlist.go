package message

import (
	"RainbowRunner/internal/byter"
	"RainbowRunner/internal/connection"
)

func HandleServerListMessage(c *connection.Connection, reader *byter.Byter) error {
	// Session1
	reader.UInt32()
	// Session2
	reader.UInt32()

	/**
	00000000 linACSendServerListExPacket struc ; (sizeof=0x24, align=0x4, copyof_819)
	00000000 baseclass_0 msgMessage ?
	00000010 m_lastServerId db ?
	00000011 db ? ; undefined
	00000012 db ? ; undefined
	00000013 db ? ; undefined
	00000014 m_serverInfoEx std::vector<dm_serverInfoEx,std::allocator<dm_serverInfoEx> > ?
	00000024 linACSendServerListExPacket ends
	00000024
	00000000 dm_serverInfoEx struc ; (sizeof=0x14, align=0x4, copyof_746)
	00000000
	00000000 m_id db ?
	00000000
	00000001 db ? ; undefined
	00000002 db ? ; undefined
	00000003 db ? ; undefined
	00000004 m_ip dd ?
	00000004
	00000008 m_port dd ?
	00000008
	0000000C m_ageLimit db ?
	0000000C
	0000000D m_pkFlag db ?
	0000000D
	0000000E m_currentUser dw ?
	0000000E
	00000010 m_maxLimitUser dw ?
	00000010
	00000012 m_serverStatus db ?
	00000012
	00000013 db ? ; undefined
	00000014 dm_serverInfoEx ends
	*/
	response := byter.NewByter(make([]byte, 0, 128))
	response.WriteByte(0x04)
	response.WriteByte(0x01) // If this is 0 then "World is down"
	response.WriteByte(0x01) // Unk

	response.WriteByte(0x0D) // Server ID

	// Server Entry
	response.WriteUInt32(0x7F000001) // IP 127.0.0.1
	response.WriteUInt32(0x00000A2B) // Port 2603
	response.WriteBool(false)        // Age limit
	response.WriteBool(false)        // PKFlag
	response.WriteUInt16(0x0000)     // Current User ? User count?
	response.WriteUInt16(0xFFFF)     // Max user count
	response.WriteByte(0x01)         // Server Status

	err := c.WriteMessageByter(response)

	if err != nil {
		return err
	}

	return nil
}
