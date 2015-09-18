package chipmunk

import (
	"bytes"
	"encoding/binary"
	"time"
)

// TimeValue holds the Timestamp and Value to be stored
type TimeValue struct {
	Timestamp time.Time
	Value     float32
}

// PackValues will pack TimeValue records into a byte array. Times are stored at second resolution.
func PackValues(timeVals []TimeValue) []byte {
	bp := newBitPacker()

	// Find most frequent delta
	deltaFreq := make(map[int64]int64)
	lasttv := timeVals[0]
	for _, tv := range timeVals[1:] {
		delta := int64(tv.Timestamp.Sub(lasttv.Timestamp).Seconds())
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

	startTime := timeVals[0].Timestamp
	startValue := timeVals[0].Value

	var b bytes.Buffer

	binary.Write(&b, binary.BigEndian, uint16(deltaSample))
	binary.Write(&b, binary.BigEndian, startTime.Unix())
	binary.Write(&b, binary.BigEndian, startValue)

	lastTime := startTime
	lastValue := startValue

	for _, tv := range timeVals[1:] {
		length, packVal := timePack(deltaSample, lastTime.Unix(), tv.Timestamp.Unix())
		bp.Add(length, packVal)
		length, packVal = valuePack(lastValue, tv.Value)
		bp.Add(length, packVal)

		lastTime = tv.Timestamp
		lastValue = tv.Value
	}

	b.Write(bp.Bytes())

	return b.Bytes()
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
