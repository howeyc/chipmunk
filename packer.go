package chipmunkdb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

	bitopts "github.com/cmchao/go-bitops"
)

type BitPacker interface {
	Add(length uint, value uint64)
	Bytes() []byte
}

type bitPacker struct {
	array      []uint64
	currentIdx int
	currentBit uint
}

func NewBitPacker() BitPacker {
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

type ValuePacker interface {
	Add(timestamp time.Time, value float32)
	Bytes() []byte
}

type timeVal struct {
	t int64
	v float32
}

type valPacker struct {
	bp BitPacker

	timeVals []timeVal
}

func NewValuePacker() ValuePacker {
	return &valPacker{bp: NewBitPacker()}
}

func (vp *valPacker) Add(timestamp time.Time, value float32) {
	vp.timeVals = append(vp.timeVals, timeVal{timestamp.Unix(), value})
}

func (vp *valPacker) Bytes() []byte {
	// TODO: calculate deltaSample
	deltaSample := int64(15)

	startTime := vp.timeVals[0].t
	startValue := vp.timeVals[0].v

	var b bytes.Buffer

	binary.Write(&b, binary.BigEndian, uint16(deltaSample))
	binary.Write(&b, binary.BigEndian, startTime)
	binary.Write(&b, binary.BigEndian, startValue)

	lastTime := startTime
	lastValue := startValue

	for _, tv := range vp.timeVals[1:] {
		length, packVal := timePack(deltaSample, lastTime, tv.t)
		vp.bp.Add(length, packVal)
		length, packVal = valuePack(lastValue, tv.v)
		vp.bp.Add(length, packVal)

		lastTime = tv.t
		lastValue = tv.v
	}

	b.Write(vp.bp.Bytes())

	return b.Bytes()
}

func unPackValues(b []byte) []timeVal {
	var deltaSample uint16
	var lastTime int64
	var lastValue float32

	buf := bytes.NewBuffer(b)

	binary.Read(buf, binary.BigEndian, &deltaSample)
	binary.Read(buf, binary.BigEndian, &lastTime)
	binary.Read(buf, binary.BigEndian, &lastValue)

	bs := NewBitStream(buf.Bytes())

	vals := []timeVal{{lastTime, lastValue}}
	fmt.Println(vals)

	for {
		packed := bs.Peek()
		tlength, nTime := timeUnPack(int64(deltaSample), lastTime, packed)
		bs.AdvanceBits(tlength)
		packed = bs.Peek()
		vlength, nVal := valueUnPack(lastValue, packed)
		bs.AdvanceBits(vlength)

		lastValue = nVal
		lastTime = nTime

		fmt.Println(timeVal{nTime, nVal})

		vals = append(vals, timeVal{nTime, nVal})
	}
}
