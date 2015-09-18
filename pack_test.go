package chipmunk

import (
	"fmt"
	"testing"
	"time"
)

func TestPacker(t *testing.T) {
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

	fmt.Println(UnPackValues(PackValues(valList)))
}
