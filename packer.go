package chipmunkdb

import (
	"bytes"
	"encoding/binary"
	"time"
)

// ValuePacker will compress TimeValue records into a byte array
type ValuePacker interface {
	Add(tv TimeValue)
	Bytes() []byte
}

type timeVal struct {
	t int64
	v float32
}

type valPacker struct {
	bp *bitPacker

	timeVals []timeVal
}

// NewValuePacker will create a new ValuePacker
// ValuePacker is used to pack TimeValue records into a byte array
func NewValuePacker() ValuePacker {
	return &valPacker{bp: newBitPacker()}
}

// Add a TimeValue to the packed byte array
func (vp *valPacker) Add(tv TimeValue) {
	vp.timeVals = append(vp.timeVals, timeVal{tv.Timestamp.Unix(), tv.Value})
}

// Retrieve the compressed values as a byte array
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

// TimeValue holds the Timestamp and Value to be stored
type TimeValue struct {
	Timestamp time.Time
	Value     float32
}

// UnPackValues will return an array of TimeValue records for a given byte array
func UnPackValues(b []byte) []TimeValue {
	var deltaSample uint16
	var lastTime int64
	var lastValue float32

	buf := bytes.NewBuffer(b)

	binary.Read(buf, binary.BigEndian, &deltaSample)
	binary.Read(buf, binary.BigEndian, &lastTime)
	binary.Read(buf, binary.BigEndian, &lastValue)

	bs := newBitStream(buf.Bytes())

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
