package chipmunkdb

import (
	"bytes"
	"encoding/binary"
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
	AddTimeValue(tv TimeValue)
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

func (vp *valPacker) AddTimeValue(tv TimeValue) {
	vp.timeVals = append(vp.timeVals, timeVal{tv.Timestamp.Unix(), tv.Value})
}

func (vp *valPacker) Add(timestamp time.Time, value float32) {
	vp.timeVals = append(vp.timeVals, timeVal{timestamp.Unix(), value})
}

func (vp *valPacker) Bytes() []byte {
	// Find most frequent delta
	deltaFreq := make(map[int64]int64)
	lasttv := vp.timeVals[0]
	for _, tv := range vp.timeVals[1:] {
		delta := tv.t - lasttv.t
		freq := deltaFreq[delta]
		freq++
		deltaFreq[delta] = freq

		lasttv = tv
	}
	var maxFreq, deltaSample int64
	for delta, freq := range deltaFreq {
		if freq > maxFreq {
			deltaSample = delta
			maxFreq = freq
		}
	}

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

type TimeValue struct {
	Timestamp time.Time
	Value     float32
}

func UnPackValues(b []byte) []TimeValue {
	var deltaSample uint16
	var lastTime int64
	var lastValue float32

	buf := bytes.NewBuffer(b)

	binary.Read(buf, binary.BigEndian, &deltaSample)
	binary.Read(buf, binary.BigEndian, &lastTime)
	binary.Read(buf, binary.BigEndian, &lastValue)

	bs := NewBitStream(buf.Bytes())

	vals := []TimeValue{{time.Unix(lastTime, 0), lastValue}}

	for {
		packed := bs.Peek()
		tlength, nTime := timeUnPack(int64(deltaSample), lastTime, packed)
		if err := bs.AdvanceBits(tlength); err != nil {
			break
		}
		packed = bs.Peek()
		vlength, nVal := valueUnPack(lastValue, packed)
		if err := bs.AdvanceBits(vlength); err != nil {
			break
		}

		lastValue = nVal
		lastTime = nTime

		vals = append(vals, TimeValue{time.Unix(nTime, 0), nVal})
	}

	return vals
}
