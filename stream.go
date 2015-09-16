package chipmunkdb

import (
	"bytes"
	"encoding/binary"
)

type BitStream struct {
	array []uint64

	bitNumber uint
}

func NewBitStream(b []byte) *BitStream {
	buf := bytes.NewBuffer(b)

	size := buf.Len()
	iLen := size / 8

	arr := make([]uint64, iLen)
	var k uint64
	for i := 0; i < iLen; i++ {
		binary.Read(buf, binary.BigEndian, &k)
		arr[i] = k
	}

	return &BitStream{arr, 63}
}

func (bs *BitStream) AdvanceBits(count uint) {
	if bs.bitNumber >= count {
		bs.bitNumber -= count
		return
	} else if bs.bitNumber == (count + 1) {
		bs.bitNumber = 63
		bs.array = bs.array[1:]
		return
	}

	count -= (bs.bitNumber + 1)
	bs.array = bs.array[1:]
	bs.bitNumber = 63
	bs.bitNumber -= count
}

func (bs *BitStream) Peek() uint64 {
	sl := 63 - bs.bitNumber
	sr := 64 - sl
	peekVal := bs.array[0] << sl
	peekVal |= bs.array[1] >> sr

	return peekVal
}
