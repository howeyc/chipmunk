package chipmunkdb

import (
	"fmt"
	"testing"
	"time"
)

func TestPacker(t *testing.T) {
	vp := NewValuePacker()

	var valList []TimeValue

	curtime := time.Now()
	valList = append(valList, TimeValue{curtime, 32})
	curtime = curtime.Add(12 * time.Second)
	valList = append(valList, TimeValue{curtime, 34})
	curtime = curtime.Add(12 * time.Second)
	valList = append(valList, TimeValue{curtime, 36})
	curtime = curtime.Add(15 * time.Second)
	valList = append(valList, TimeValue{curtime, 36})
	curtime = curtime.Add(18 * time.Second)
	valList = append(valList, TimeValue{curtime, 41})

	fmt.Println(valList)

	for _, tv := range valList {
		vp.AddTimeValue(tv)
	}

	fmt.Println(UnPackValues(vp.Bytes()))
}
