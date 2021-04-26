package byter

import "encoding/binary"

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

	binary.LittleEndian.PutUint32(b.Buffer[len(b.Buffer)-4:], i)

	return nil
}

func (b *Byter) WriteUInt16(i uint16) error {
	b.Buffer = append(b.Buffer, []byte{0, 0}...)

	binary.LittleEndian.PutUint16(b.Buffer[len(b.Buffer)-2:], i)

	return nil
}

func NewByter(buffer []byte) *Byter {
	return &Byter{
		Buffer: buffer,
	}
}
