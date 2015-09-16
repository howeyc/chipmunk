package chipmunkdb

import (
	"fmt"
	"unsafe"

	bitopts "github.com/cmchao/go-bitops"
)

func countLeadingZero32(value uint32) (count uint) {
	for value > 0 {
		if value&0x80000000 == 0 {
			count++
			value = value << 1
		} else {
			return
		}
	}

	return
}

func countTrailingZero32(value uint32) (count uint) {
	for value > 0 {
		if value&1 == 0 {
			count++
			value = value >> 1
		} else {
			return
		}
	}

	return
}

// Pack a floating point value change for placement into a uint64 "bit-stream"
//
// If there is no change, packed value is a one-bit long value of 0
// If there is change, packed value is a variable-length bit value:
//	1-bits - Set to 1
//	5-bits - Number of leading zeros in XOR of fPrev and fCurr
//	6-bits - Length of significant bits of XOR of fPrev and fCurr
//	X-bits - Value of XOR of fPrev and fCurr
//
// Packed value is returned such that the LSB of packed value starts at bit 0
// (that is, the packed value is shifted such that there are a bunch of leading
// zeros in the MSB before the first bit, which is set to 1 in the case of a change)
func valuePack(fPrev, fCurr float32) (length uint, packed uint64) {
	xorval := (*(*uint32)(unsafe.Pointer(&fPrev))) ^ (*(*uint32)(unsafe.Pointer(&fCurr)))

	if xorval == 0 {
		return 1, 0
	}

	fmt.Println("valuePack: ", xorval)

	leadzeroLen := countLeadingZero32(xorval)
	trailzeroLen := countTrailingZero32(xorval)
	sigbitLen := 32 - (leadzeroLen + trailzeroLen)

	packed, _ = bitopts.SetBit64(packed, 63)
	packed, _ = bitopts.SetField64(packed, 62, 58, uint64(leadzeroLen))
	packed, _ = bitopts.SetField64(packed, 57, 52, uint64(sigbitLen))
	packed, _ = bitopts.SetField64(packed, 51, 51-(sigbitLen-1), uint64(xorval)>>trailzeroLen)

	packLen := sigbitLen + 12

	return packLen, packed >> (64 - packLen)
}

func valueUnPack(fPrev float32, packed uint64) (length uint, value float32) {
	if one, _ := bitopts.TestBit64(packed, 63); !one {
		return 1, fPrev
	}

	leadzeroLen, _ := bitopts.GetField64(packed, 62, 58)
	sigbitLen, _ := bitopts.GetField64(packed, 57, 52)
	xorval, _ := bitopts.GetField64(packed, 51, uint(51-(sigbitLen-1)))

	trailzeroLen := 32 - (leadzeroLen + sigbitLen)
	xor32 := uint32(xorval << trailzeroLen)

	fmt.Println("valueUnPack: ", xor32)

	bval := (*(*uint32)(unsafe.Pointer(&fPrev))) ^ (*(*uint32)(unsafe.Pointer(&xor32)))

	value = (*(*float32)(unsafe.Pointer(&bval)))

	return uint(sigbitLen + 12), value
}
