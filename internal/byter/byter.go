package byter

import (
	"bytes"
	"encoding/binary"
)

type Byter struct {
	i            int
	Buffer       []byte
	littleEndian bool
}

func (b *Byter) BigEndian() {
	b.littleEndian = false
}

func (b *Byter) LittleEndian() {
	b.littleEndian = true
}

func (b *Byter) Bytes(count int) []byte {
	i := b.getDataIndex(count)

	return b.Buffer[i : i+count]
}

func (b *Byter) UInt16() uint16 {
	i := b.getDataIndex(2)

	if b.littleEndian {
		return binary.LittleEndian.Uint16(b.Buffer[i:])
	} else {
		return binary.BigEndian.Uint16(b.Buffer[i:])
	}
}

func (b *Byter) UInt32() uint32 {
	i := b.getDataIndex(4)

	if b.littleEndian {
		return binary.LittleEndian.Uint32(b.Buffer[i:])
	} else {
		return binary.BigEndian.Uint32(b.Buffer[i:])
	}
}

func (b *Byter) UInt24() uint32 {
	i := b.getDataIndex(3)

	result := uint32(0)

	if b.littleEndian {
		result |= uint32(b.Buffer[i])
		i++
		result |= uint32(binary.LittleEndian.Uint16(b.Buffer[i:])) << 8
	} else {
		result |= uint32(b.Buffer[i]) << 16
		i++
		result |= uint32(binary.LittleEndian.Uint16(b.Buffer[i:]))
	}

	return result
}

func (b *Byter) UInt64() uint64 {
	var result uint64 = 0

	if b.littleEndian {
		i := b.getDataIndex(4)
		result |= uint64(binary.LittleEndian.Uint32(b.Buffer[i:])) << 32

		i = b.getDataIndex(4)
		result |= uint64(binary.LittleEndian.Uint32(b.Buffer[i:]))
	} else {
		i := b.getDataIndex(4)
		result |= uint64(binary.BigEndian.Uint32(b.Buffer[i:])) << 32

		i = b.getDataIndex(4)
		result |= uint64(binary.BigEndian.Uint32(b.Buffer[i:]))
	}

	return result
}

func (b *Byter) getDataIndex(num int) int {
	if b.Buffer == nil || len(b.Buffer)-b.i < num {
		panic("Not enough data remaining in buffer!")
	}

	i := b.i
	b.i += num
	return i
}

func (b *Byter) UInt8() uint8 {
	i := b.getDataIndex(1)

	return b.Buffer[i]
}

func (b *Byter) WriteByte(i byte) error {
	b.Buffer = append(b.Buffer, i)

	return nil
}

func (b *Byter) WriteBool(i bool) error {
	if i {
		b.Buffer = append(b.Buffer, 0x01)
	} else {
		b.Buffer = append(b.Buffer, 0x00)
	}

	return nil
}

func (b *Byter) WriteUInt32(i uint32) error {
	b.Buffer = append(b.Buffer, []byte{0, 0, 0, 0}...)

	if b.littleEndian {
		binary.LittleEndian.PutUint32(b.Buffer[len(b.Buffer)-4:], i)
	} else {
		binary.BigEndian.PutUint32(b.Buffer[len(b.Buffer)-4:], i)
	}

	return nil
}

func (b *Byter) WriteUInt24(num uint) error {
	if b.littleEndian {
		b.Buffer = append(b.Buffer, byte(num))

		b.Buffer = append(b.Buffer, []byte{0, 0}...)
		binary.LittleEndian.PutUint16(b.Buffer[len(b.Buffer)-2:], uint16(num>>8))
	} else {
		b.Buffer = append(b.Buffer, byte(num>>16))

		b.Buffer = append(b.Buffer, []byte{0, 0}...)
		binary.BigEndian.PutUint16(b.Buffer[len(b.Buffer)-2:], uint16(num))
	}

	return nil
}

func (b *Byter) WriteUInt16(i uint16) error {
	b.Buffer = append(b.Buffer, []byte{0, 0}...)

	if b.littleEndian {
		binary.LittleEndian.PutUint16(b.Buffer[len(b.Buffer)-2:], i)
	} else {
		binary.BigEndian.PutUint16(b.Buffer[len(b.Buffer)-2:], i)
	}

	return nil
}

func (b *Byter) Data() []byte {
	return b.Buffer[0:len(b.Buffer)]
}

func (b *Byter) Write(body *Byter) {
	b.Buffer = append(b.Buffer, body.Data()...)
}

func (b *Byter) Clear() {
	b.Buffer = b.Buffer[:0]
}

func (b *Byter) WriteBuffer(b2 bytes.Buffer) {
	b.Buffer = append(b.Buffer, b2.Bytes()...)
}

func (b *Byter) String() {
	str := ""

	for b.Buffer[b.i] != 0x00 {
		i := b.getDataIndex(1)
		str += string(b.Buffer[i])
	}
}

func NewByter(buffer []byte) *Byter {
	return &Byter{
		Buffer: buffer,
	}
}

func NewLEByter(buffer []byte) *Byter {
	return &Byter{
		Buffer:       buffer,
		littleEndian: true,
	}
}
