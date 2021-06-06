package byter

import (
	"bytes"
	"encoding/binary"
	"math"
)

type Byter struct {
	I            int
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

func (b *Byter) Byte() byte {
	i := b.getDataIndex(1)

	return b.Buffer[i]
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

func (b *Byter) Fixed32() float32 {
	i := b.getDataIndex(4)

	if b.littleEndian {
		frac := binary.LittleEndian.Uint16(b.Buffer[i:])
		intNum := binary.LittleEndian.Uint16(b.Buffer[i+2:])

		fracNum := math.Floor(math.Log10(math.Abs(float64(frac)))) + 1
		div := math.Pow(10, fracNum)

		finalFrac := float32(0)

		if frac != 0 {
			finalFrac = float32(frac) / float32(div)
		}

		full := float32(intNum) + finalFrac

		return full
	} else {
		panic("not implemented")
		//return float32(binary.BigEndian.Uint32(b.Buffer[i:]))
	}
}

func (b *Byter) Float32() float32 {
	i := b.getDataIndex(4)

	if b.littleEndian {
		return float32(binary.LittleEndian.Uint32(b.Buffer[i:]))
	} else {
		return float32(binary.BigEndian.Uint32(b.Buffer[i:]))
	}
}

func (b *Byter) Int32() int32 {
	return int32(b.UInt32())
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

func (b Byter) GetLength() int {
	return len(b.Buffer)
}

func (b *Byter) getDataIndex(num int) int {
	if b.Buffer == nil || len(b.Buffer)-b.I < num {
		panic("Not enough data remaining in buffer!")
	}

	i := b.I
	b.I += num
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

func (b *Byter) WriteUInt64(i uint64) error {
	b.Buffer = append(b.Buffer, []byte{0, 0, 0, 0, 0, 0, 0, 0}...)

	if b.littleEndian {
		binary.LittleEndian.PutUint64(b.Buffer[len(b.Buffer)-8:], i)
	} else {
		binary.BigEndian.PutUint64(b.Buffer[len(b.Buffer)-8:], i)
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

func (b *Byter) String() string {
	str := ""
	i := b.getDataIndex(1)

	for b.Buffer[i] != 0x00 {
		str += string(b.Buffer[i])
		i = b.getDataIndex(1)
	}

	return str
}

func (b *Byter) WriteBytes(data []byte) {
	b.Buffer = append(b.Buffer, data...)
}

func (b *Byter) WriteString(s string) {
	b.Buffer = append(b.Buffer, []byte(s)...)
}

func (b *Byter) WriteCString(s string) {
	b.WriteString(s + "\x00")
}

func (b *Byter) WriteNull() error {
	return b.WriteByte(0x00)
}

func (b *Byter) Compare(comparison []byte) bool {
	for _, b1 := range comparison {
		if b1 != b.Byte() {
			return false
		}
	}

	return true
}

func (b *Byter) CompareString(str string) bool {
	return b.Compare([]byte(str))
}

func (b *Byter) Seek(address int) {
	b.I = address
}

func (b *Byter) WriteInt32(i int32) error {
	b.Buffer = append(b.Buffer, []byte{0, 0, 0, 0}...)

	if b.littleEndian {
		binary.LittleEndian.PutUint32(b.Buffer[len(b.Buffer)-4:], uint32(i))
	} else {
		binary.BigEndian.PutUint32(b.Buffer[len(b.Buffer)-4:], uint32(i))
	}

	return nil
}

func (b *Byter) RemainingBytes() []byte {
	return b.Bytes(len(b.Buffer) - b.I)
}

func (b *Byter) HasRemainingData() bool {
	return len(b.Buffer)-b.I > 1
}

func (b *Byter) CString() string {
	str := ""

	for {
		c := b.Byte()

		if c == 0 {
			break
		}

		str += string(c)
	}

	return str
}

func (b *Byter) Bool() bool {
	v := b.Byte()

	if v == 1 {
		return true
	}

	return false
}

func (b *Byter) WriteFloat32(i float32) {
	n := math.Float32bits(i)

	if b.littleEndian {
		b.WriteByte(byte(n))
		b.WriteByte(byte(n >> 8))
		b.WriteByte(byte(n >> 16))
		b.WriteByte(byte(n >> 24))
	} else {
		b.WriteByte(byte(n >> 24))
		b.WriteByte(byte(n >> 16))
		b.WriteByte(byte(n >> 8))
		b.WriteByte(byte(n))
	}
}

func (b *Byter) Int16() int16 {
	return int16(b.UInt16())
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
