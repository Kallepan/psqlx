package psqlx

import (
	"encoding/binary"
)

type Buffer struct {
	data []byte
}

func (b *Buffer) WriteInt32(n int32) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(n))
	b.data = append(b.data, buf...)
}

func (b *Buffer) WriteInt16(n int16) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(n))
	b.data = append(b.data, buf...)
}

func (b *Buffer) WriteString(str string) {
	b.data = append(b.data, []byte(str)...)
	b.data = append(b.data, strSeparator)
}

func (b *Buffer) WriteBytes(data []byte) {
	b.data = append(b.data, data...)
}

func (b *Buffer) WriteByte(data byte) {
	b.data = append(b.data, data)
}

func (b *Buffer) CalculateSize(prefix int) {
	l := len(b.data)
	data := b.data

	binary.BigEndian.PutUint32(data[prefix:], uint32(l-prefix))

	b.data = data
}

func (b *Buffer) Data() []byte {
	return b.data
}
