package chipmunk

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type bitStream struct {
	array []uint64

	bitNumber uint
}

func newBitStream(b []byte) *bitStream {
	buf := bytes.NewBuffer(b)

	size := buf.Len()
	iLen := size / 8

	arr := make([]uint64, iLen)
	var k uint64
	for i := 0; i < iLen; i++ {
		binary.Read(buf, binary.BigEndian, &k)
		arr[i] = k
	}

	return &bitStream{arr, 63}
}

func (bs *bitStream) AdvanceBits(count uint) error {
	if bs.bitNumber >= count {
		bs.bitNumber -= count
		return nil
	} else if bs.bitNumber == (count + 1) {
		bs.bitNumber = 63
		if len(bs.array) <= 1 {
			return errors.New("end of stream")
		}
		bs.array = bs.array[1:]
		return nil
	}

	if len(bs.array) <= 1 {
		return errors.New("end of stream")
	}

	count -= (bs.bitNumber + 1)
	bs.array = bs.array[1:]
	bs.bitNumber = 63
	bs.bitNumber -= count

	return nil
}

func (bs *bitStream) Peek() (peekVal uint64) {
	sl := 63 - bs.bitNumber
	sr := 64 - sl
	if len(bs.array) > 0 {
		peekVal = bs.array[0] << sl
		if len(bs.array) > 1 {
			peekVal |= bs.array[1] >> sr
		}
	}

	return peekVal
}
