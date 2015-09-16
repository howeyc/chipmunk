package chipmunkdb

import (
	"fmt"
	"testing"
)

func TestValuePack(t *testing.T) {
	packer := NewBitPacker()

	length, pack := valuePack(2, 4)
	packer.Add(length, pack)
	fmt.Printf("%d: 0x%016X\n", length, pack)

	length, pack = valuePack(12.6, 42)
	packer.Add(length, pack)
	fmt.Printf("%d: 0x%016X\n", length, pack)

	fmt.Printf("0x%X\n", packer.Bytes())
}

func TestTimePack(t *testing.T) {
	length, pack := timePack(15, 300, 315)
	fmt.Printf("%d: 0x%016X\n", length, pack)
	length, pack = timePack(15, 300, 320)
	fmt.Printf("%d: 0x%016X\n", length, pack)
	length, pack = timePack(15, 300, 310)
	fmt.Printf("%d: 0x%016X\n", length, pack)
}
