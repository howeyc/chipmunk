package chipmunkdb

import (
	bitopts "github.com/cmchao/go-bitops"
)

func timePack(deltaSample, tPrev, tCurr int64) (length uint, packed uint64) {
	tDelta := tCurr - tPrev
	packVal := tDelta - deltaSample

	if packVal == 0 {
		return 1, 0
	}

	if packVal >= -31 && packVal <= 31 {
		packed, _ = bitopts.SetBit64(packed, 63)
		if packVal < 0 {
			packVal = -packVal
			packed, _ = bitopts.SetBit64(packed, 62)
		}
		packed, _ = bitopts.SetField64(packed, 61, 57, uint64(packVal))
	} else {
		panic("unable to pack time")
	}

	return 7, packed >> (64 - 7)
}
