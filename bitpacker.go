package chipmunk

import (
	"bytes"
	"encoding/binary"

	bitopts "github.com/cmchao/go-bitops"
)

type bitPacker struct {
	array      []uint64
	currentIdx int
	currentBit uint
}

func newBitPacker() *bitPacker {
	vp := &bitPacker{array: make([]uint64, 1), currentBit: 63}
	return vp
}

func (vp *bitPacker) Add(length uint, value uint64) {
	lowbit := int64(vp.currentBit) - (int64(length) - 1)
	if lowbit < 0 {
		sr := uint64(length) - (uint64(vp.currentBit) + 1)
		fieldVal := value >> sr
		vp.array[vp.currentIdx], _ = bitopts.SetField64(vp.array[vp.currentIdx], vp.currentBit, 0, uint64(fieldVal))
		vp.array = append(vp.array, uint64(value<<(64-sr)))
		vp.currentBit = 63 - uint(sr)
		vp.currentIdx++
	} else {
		vp.array[vp.currentIdx], _ = bitopts.SetField64(vp.array[vp.currentIdx], vp.currentBit, uint(lowbit), value)
		if lowbit == 0 {
			vp.array = append(vp.array, uint64(0))
			vp.currentBit = 63
			vp.currentIdx++
		} else {
			vp.currentBit = uint(lowbit) - 1
		}
	}
}

func (vp *bitPacker) Bytes() []byte {
	var b bytes.Buffer

	for _, val := range vp.array {
		binary.Write(&b, binary.BigEndian, val)
	}
	return b.Bytes()
}
