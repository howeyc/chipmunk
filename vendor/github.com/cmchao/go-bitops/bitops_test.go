package bitops

import "testing"

func TestExtract32(t *testing.T) {
    var value uint32 = 0xF0F0F0F0
    var length uint
    var start uint
    var field uint32

    //check error
    start = 32
    length = 0
    _, err := Extract32(value, start, length)
    if err == nil {
        t.Fail()
        t.Log("invalid start")
    }

    start = 0
    length = 33
    _, err = Extract32(value, start, length)
    if err == nil {
        t.Fail()
        t.Log("invalid length")
    }

    start = 31
    length = 2
    _, err = Extract32(value, start, length)
    if err == nil {
        t.Fail()
        t.Log("invalid length from valid start")
    }

    //check pass case
    start = 0
    length = 1
    field, err = Extract32(value, start, length)
    if err != nil || field != 0x0 {
        t.Fail()
        t.Log("LSB")
    }

    start = 31
    length = 1
    field, err = Extract32(value, start, length)
    if err != nil || field != 0x1 {
        t.Fail()
        t.Log("MSB")
    }

    start = 4
    length = 4
    field, err = Extract32(value, start, length)
    if err != nil || field != 0xF {
        t.Fail()
        t.Log("get valid field")
    }
}

func TestExtract64(t *testing.T) {
    var value uint64 = 0xF0F0F0F0F0F0F0F0
    var length uint
    var start uint
    var field uint64

    //check error
    start = 64
    length = 0
    _, err := Extract64(value, start, length)
    if err == nil {
        t.Fail()
        t.Log("invalid start")
    }

    start = 0
    length = 65
    _, err = Extract64(value, start, length)
    if err == nil {
        t.Fail()
        t.Log("invalid length")
    }

    start = 63
    length = 2
    _, err = Extract64(value, start, length)
    if err == nil {
        t.Fail()
        t.Log("invalid length from valid start")
    }

    //check pass case
    start = 0
    length = 1
    field, err = Extract64(value, start, length)
    if err != nil || field != 0x0 {
        t.Fail()
        t.Log("LSB")
    }

    start = 63
    length = 1
    field, err = Extract64(value, start, length)
    if err != nil || field != 0x1 {
        t.Fail()
        t.Log("MSB")
    }

    start = 12 
    length = 4
    field, err = Extract64(value, start, length)
    if err != nil || field != 0xF {
        t.Fail()
        t.Log("get valid field")
    }
}

func TestGetField32(t *testing.T) {
    var value uint32 = 0xF0F0F0F0
    var high uint
    var low uint
    var field uint32

    //check error
    high = 32
    low = 0 
    _, err := GetField32(value, high, low)
    if err == nil {
        t.Fail()
        t.Log("invalid high")
    }

    high = 0
    low = 32
    _, err = GetField32(value, high, low)
    if err == nil {
        t.Fail()
        t.Log("invalid low")
    }

    high = 10
    low = 20
    _, err = GetField32(value, high, low)
    if err == nil {
        t.Fail()
        t.Log("high < low")
    }

    //check pass case
    high = 0
    low = 0
    field, err = GetField32(value, high, low)
    if err != nil || field != 0x0 {
        t.Fail()
        t.Log("LSB")
    }

    high = 31
    low = 31
    field, err = GetField32(value, high, low)
    if err != nil || field != 0x1 {
        t.Fail()
        t.Log("MSB")
    }

    high = 7
    low = 4
    field, err = GetField32(value, high, low)
    if err != nil || field != 0xF {
        t.Fail()
        t.Log("get valid field")
    }
}

func TestGetField64(t *testing.T) {
    var value uint64 = 0xF0F0F0F0F0F0F0F0
    var high uint
    var low uint
    var field uint64

    //check error
    high = 64
    low = 0 
    _, err := GetField64(value, high, low)
    if err == nil {
        t.Fail()
        t.Log("invalid high")
    }

    high = 0
    low = 64
    _, err = GetField64(value, high, low)
    if err == nil {
        t.Fail()
        t.Log("invalid low")
    }

    high = 10
    low = 20
    _, err = GetField64(value, high, low)
    if err == nil {
        t.Fail()
        t.Log("high < low")
    }

    //check pass case
    high = 0
    low = 0
    field, err = GetField64(value, high, low)
    if err != nil || field != 0x0 {
        t.Fail()
        t.Log("LSB")
    }

    high = 63
    low = 63
    field, err = GetField64(value, high, low)
    if err != nil || field != 0x1 {
        t.Fail()
        t.Log("MSB")
    }

    high = 7
    low = 4
    field, err = GetField64(value, high, low)
    if err != nil || field != 0xF {
        t.Fail()
        t.Log("get valid field")
    }
}

func TestDeposit32(t *testing.T) {
    var value uint32 = 0xF0F0F0F0
    var length uint
    var start uint
    var field uint32

    //check error
    start = 32
    length = 0
    field, err := Deposit32(value, start, length, 0x1234)
    if err == nil {
        t.Fail()
        t.Log("invalid start")
    }

    if value != field {
        t.Fail()
        t.Logf("invalid return %x", field)
    }

    start = 0
    length = 33
    field, err = Deposit32(value, start, length, 0x1234)
    if err == nil {
        t.Fail()
        t.Log("invalid length")
    }

    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    start = 31
    length = 2
    field, err = Deposit32(value, start, length, 0x1234)
    if err == nil {
        t.Fail()
        t.Log("invalid length from valid start")
    }
    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    //check pass case
    start = 0
    length = 1
    field, err = Deposit32(value, start, length, 0x7)
    if err != nil || field != 0xF0F0F0F1 {
        t.Fail()
        t.Log("LSB")
    }

    start = 31
    length = 1
    field, err = Deposit32(value, start, length, 0x0)
    if err != nil || field != 0x70F0F0F0 {
        t.Fail()
        t.Log("MSB")
    }

    start = 4
    length = 16
    field, err = Deposit32(value, start, length, 0x1234)
    if err != nil || field != 0xF0F12340 {
        t.Fail()
        t.Log("get valid field")
    }
}

func TestDeposit64(t *testing.T) {
    var value uint64 = 0xF0F0F0F0F0F0F0F0
    var length uint
    var start uint
    var field uint64

    //check error
    start = 64
    length = 0
    field, err := Deposit64(value, start, length, 0x1234)
    if err == nil {
        t.Fail()
        t.Log("invalid start")
    }

    if value != field {
        t.Fail()
        t.Logf("invalid return %x", field)
    }

    start = 0
    length = 65
    field, err = Deposit64(value, start, length, 0x1234)
    if err == nil {
        t.Fail()
        t.Log("invalid length")
    }

    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    start = 63
    length = 2
    field, err = Deposit64(value, start, length, 0x1234)
    if err == nil {
        t.Fail()
        t.Log("invalid length from valid start")
    }
    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    //check pass case
    start = 0
    length = 1
    field, err = Deposit64(value, start, length, 0x7)
    if err != nil || field != 0xF0F0F0F0F0F0F0F1 {
        t.Fail()
        t.Log("LSB")
    }

    start = 63
    length = 1
    field, err = Deposit64(value, start, length, 0x0)
    if err != nil || field != 0x70F0F0F0F0F0F0F0 {
        t.Fail()
        t.Log("MSB")
    }

    start = 4
    length = 16
    field, err = Deposit64(value, start, length, 0x1234)
    if err != nil || field != 0xF0F0F0F0F0F12340 {
        t.Fail()
        t.Log("get valid field")
    }
}

func TestSetField32(t *testing.T) {
    var value uint32 = 0xF0F0F0F0
    var high uint
    var low uint
    var field uint32

    //check error
    high = 32
    low = 0
    field, err := SetField32(value, high, low, 0)
    if err == nil {
        t.Fail()
        t.Log("invalid high")
    }
    if value != field {
        t.Fail()
        t.Logf("invalid return %x", field)
    }

    high = 0
    low = 32
    field, err = SetField32(value, high, low, 0)
    if err == nil {
        t.Fail()
        t.Log("invalid low")
    }
    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    high = 10
    low = 20
    field, err = SetField32(value, high, low, 0)
    if err == nil {
        t.Fail()
        t.Log("high < low")
    }
    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    //check pass case
    high = 0
    low = 0
    field, err = SetField32(value, high, low, 0x7)
    if err != nil || field != 0xF0F0F0F1 {
        t.Fail()
        t.Log("LSB")
    }

    high = 31
    low = 31
    field, err = SetField32(value, high, low, 0x0)
    if err != nil || field != 0x70F0F0F0 {
        t.Fail()
        t.Log("MSB")
    }

    high = 19
    low = 4
    field, err = SetField32(value, high, low, 0x1234)
    if err != nil || field != 0xF0F12340 {
        t.Fail()
        t.Log("get valid field")
    }
}

func TestSetField64(t *testing.T) {
    var value uint64 = 0xF0F0F0F0F0F0F0F0
    var high uint
    var low uint
    var field uint64

    //check error
    high = 64
    low = 0
    field, err := SetField64(value, high, low, 0)
    if err == nil {
        t.Fail()
        t.Log("invalid high")
    }
    if value != field {
        t.Fail()
        t.Logf("invalid return %x", field)
    }

    high = 0
    low = 64
    field, err = SetField64(value, high, low, 0)
    if err == nil {
        t.Fail()
        t.Log("invalid low")
    }
    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    high = 10
    low = 20
    field, err = SetField64(value, high, low, 0)
    if err == nil {
        t.Fail()
        t.Log("high < low")
    }
    if value != field {
        t.Fail()
        t.Log("invalid return")
    }

    //check pass case
    high = 0
    low = 0
    field, err = SetField64(value, high, low, 0x7)
    if err != nil || field != 0xF0F0F0F0F0F0F0F1 {
        t.Fail()
        t.Log("LSB")
    }

    high = 63
    low = 63
    field, err = SetField64(value, high, low, 0x0)
    if err != nil || field != 0x70F0F0F0F0F0F0F0 {
        t.Fail()
        t.Log("MSB")
    }

    high = 19
    low = 4
    field, err = SetField64(value, high, low, 0x1234)
    if err != nil || field != 0xF0F0F0F0F0F12340 {
        t.Fail()
        t.Log("get valid field")
    }
}

func TestCountOne8(t *testing.T) {
    var value uint8 = 0xA5
    count := CountOne8(value);
    if count != 4 {
        t.Fail()
        t.Logf("expect %d for %x but get %x", 4, count, value)
    }
}

func TestCountOne16(t *testing.T) {
    var value uint16 = 0xA5A6
    count := CountOne16(value)
    if count != 8 {
        t.Fail()
        t.Logf("expect %d for %x but get %x", 8, count, value)
    }
}

func TestCountOne32(t *testing.T) {
    var value uint32 = 0xA5A5A5A5
    count := CountOne32(value)
    if count != 16 {
        t.Fail()
        t.Logf("expect %d for %x but get %x", 16, count, value)
    }
}

func TestCountOne64(t *testing.T) {
    var value uint64 = 0xA5A5A5A5A5A5A5A5
    count := CountOne64(value)
    if count != 32 {
        t.Fail()
        t.Logf("expect %d for %x but get %x", 32, count, value)
    }
}

func TestCountZero8(t *testing.T) {
    var value uint8 = 0xA5
    count := CountZero8(value);
    if count != 4 {
        t.Fail()
        t.Logf("expect %d for %x but get %x", 4, count, value)
    }
}

func TestCountZero16(t *testing.T) {
    var value uint16 = 0xA5A6
    count := CountZero16(value)
    if count != 8 {
        t.Fail()
        t.Logf("expect %d for %x but get %x", 8, count, value)
    }
}

func TestCountZero32(t *testing.T) {
    var value uint32 = 0xA5A5A5A5
    count := CountZero32(value)
    if count != 16 {
        t.Fail()
        t.Logf("expect %d for %x but get %x", 16, count, value)
    }
}

func TestCountZero64(t *testing.T) {
    var value uint64 = 0xA5A5A5A5A5A5A5A5
    count := CountZero64(value)
    if count != 32 {
        t.Fail()
        t.Logf("expect %d for %x but get %x", 32, count, value)
    }
}

func TestCountTrailZero32(t *testing.T) {
    var value uint32
    var count, expect_cnt, i uint

    value = 0xFFFFFFF0
    for i = 0; i < 4; i++ {
        count = CountTrailZero32(value)
        if expect_cnt = i *  8 + 4; count != expect_cnt {
            t.Fail()
            t.Logf("expect %d for %d but get %x", expect_cnt, count, value)
        }
        value <<= 8;
    }
}

func TestCountTrailZero64(t *testing.T) {
    var value uint64
    var count, expect_cnt, i uint

    value = 0xFFFFFFFFFFFFFFF0
    for i = 0; i < 8; i++ {
        count = CountTrailZero64(value)
        if expect_cnt = i *  8 + 4; count != expect_cnt {
            t.Fail()
            t.Logf("expect %d for %d but get %x", expect_cnt, count, value)
        }

        value <<= 8;
    }
}

func TestCountTrailOne32(t *testing.T) {
    var value uint32
    var count, expect_cnt, i uint

    value = 0x0FFFFFFF
    for i = 4; i > 0; i-- {
        count = CountTrailOne32(value)
        if expect_cnt = i * 8 - 4; count != expect_cnt {
            t.Fail()
            t.Logf("expect %d for %d but get %x", expect_cnt, count, value)
        }

        value >>= 8
    }
}

func TestCountTrailOne64(t *testing.T) {
    var value uint64
    var count, expect_cnt, i uint

    value = 0x0FFFFFFFFFFFFFFF
    for i = 8; i > 0; i-- {
        count = CountTrailOne64(value)
        if expect_cnt = i * 8 - 4; count != expect_cnt {
            t.Fail()
            t.Logf("expect %d for %d but get %x", expect_cnt, count, value)
        }

        value >>= 8
    }
}

func TestCountLeadZero32(t *testing.T) {
    var value uint32
    var count, expect_cnt, i uint

    value = 0x0FFFFFFF
    for i = 0; i < 4; i++ {
        count = CountLeadZero32(value)
        if expect_cnt = i * 8 + 4; count != expect_cnt {
            t.Fail()
            t.Logf("expect %2d for %2d but get %.8x", expect_cnt, count, value)
        }

        value >>= 8
    }
}

func TestCountLeadZero64(t *testing.T) {
    var value uint64
    var count, expect_cnt, i uint

    value = 0x0FFFFFFFFFFFFFFF
    for i = 0; i < 8; i++ {
        count = CountLeadZero64(value)
        if expect_cnt = i * 8 + 4; count != expect_cnt {
            t.Fail()
            t.Logf("expect %2d for %2d but get %.8x", expect_cnt, count, value)
        }

        value >>= 8
    }
}

func TestCountLeadOne32(t *testing.T) {
    var value uint32
    var count, expect_cnt, i uint

    value = 0xFFFFFFF0
    for i = 4; i > 0; i-- {
        count = CountLeadOne32(value)
        if expect_cnt = i * 8 - 4; count != expect_cnt {
            t.Fail()
            t.Logf("expect %2d for %2d but get %.8x", expect_cnt, count, value)
        }

        value <<= 8
    }
}

func TestCountLeadOne64(t *testing.T) {
    var value uint64
    var count, expect_cnt, i uint

    value = 0xFFFFFFFFFFFFFFF0
    for i = 8; i > 0; i-- {
        count = CountLeadOne64(value)
        if expect_cnt = i * 8 - 4; count != expect_cnt {
            t.Fail()
            t.Logf("expect %2d for %2d but get %.8x", expect_cnt, count, value)
        }

        value <<= 8
    }
}

func TestSetBit32(t *testing.T) {
    var value, ret uint32
    var pos uint
    var err error

    value = 0x0

    pos = 32 
    ret, err = SetBit32(value, pos)
    if err == nil {
        t.Logf("expect error")
        t.Fail()
    }

    pos = 0
    ret, err = SetBit32(value, pos)
    if err != nil || ret != 0x1{
        t.Logf("LSB error")
        t.Fail()
    }

    pos = 31
    ret, err = SetBit32(value, pos)
    if err != nil || ret != 0x80000000{
        t.Logf("MSB error")
        t.Fail()
    }

}

func TestSetBit64(t *testing.T) {
    var value, ret uint64
    var pos uint
    var err error

    value = 0x0

    pos = 64 
    ret, err = SetBit64(value, pos)
    if err == nil {
        t.Logf("expect error")
        t.Fail()
    }

    pos = 0
    ret, err = SetBit64(value, pos)
    if err != nil || ret != 0x1{
        t.Logf("LSB error")
        t.Fail()
    }

    pos = 63
    ret, err = SetBit64(value, pos)
    if err != nil || ret != uint64(1) << 63{
        t.Logf("MSB error")
        t.Fail()
    }

}

func TestToggleBit32(t *testing.T) {
    var value, ret uint32
    var pos uint
    var err error

    value = 0x0

    pos = 32 
    ret, err = ToggleBit32(value, pos)
    if err == nil {
        t.Logf("expect error")
        t.Fail()
    }

    pos = 0
    ret, err = ToggleBit32(value, pos)
    if err != nil || ret != 0x1{
        t.Logf("LSB error")
        t.Fail()
    }

    pos = 31
    ret, err = ToggleBit32(value, pos)
    if err != nil || ret != 0x80000000{
        t.Logf("MSB error")
        t.Fail()
    }

    value = 0x80000001
    pos = 31
    ret, err = ToggleBit32(value, pos)
    if err != nil || ret != 0x1{
        t.Logf("LSB error")
        t.Fail()
    }

    pos = 0
    ret, err = ToggleBit32(value, pos)
    if err != nil || ret != 0x80000000{
        t.Logf("MSB error")
        t.Fail()
    }

}

func TestToggleBit64(t *testing.T) {
    var value, ret uint64
    var pos uint
    var err error

    value = 0x0

    pos = 64 
    ret, err = ToggleBit64(value, pos)
    if err == nil {
        t.Logf("expect error")
        t.Fail()
    }

    pos = 0
    ret, err = ToggleBit64(value, pos)
    if err != nil || ret != 0x1{
        t.Logf("LSB error")
        t.Fail()
    }

    pos = 63
    ret, err = ToggleBit64(value, pos)
    if err != nil || ret != uint64(1) << 63{
        t.Logf("MSB error")
        t.Fail()
    }

    value = 0x8000000000000001
    pos = 63
    ret, err = ToggleBit64(value, pos)
    if err != nil || ret != 0x1{
        t.Logf("LSB error")
        t.Fail()
    }

    pos = 0
    ret, err = ToggleBit64(value, pos)
    if err != nil || ret != uint64(1) << 63{
        t.Logf("MSB error")
        t.Fail()
    }
}

func TestClearBit32(t *testing.T) {
    var value, ret uint32
    var pos uint
    var err error

    value = 0x80000001

    pos = 32 
    ret, err = ClearBit32(value, pos)
    if err == nil {
        t.Logf("expect error")
        t.Fail()
    }

    pos = 31
    ret, err = ClearBit32(value, pos)
    if err != nil || ret != 0x1{
        t.Logf("LSB error")
        t.Fail()
    }

    pos = 0
    ret, err = ClearBit32(value, pos)
    if err != nil || ret != 0x80000000{
        t.Logf("MSB error")
        t.Fail()
    }

}

func TestClearBit64(t *testing.T) {
    var value, ret uint64
    var pos uint
    var err error

    value = 0x8000000000000001

    pos = 64 
    ret, err = ClearBit64(value, pos)
    if err == nil {
        t.Logf("expect error")
        t.Fail()
    }

    pos = 63
    ret, err = ClearBit64(value, pos)
    if err != nil || ret != 0x1{
        t.Logf("LSB error")
        t.Fail()
    }

    pos = 0
    ret, err = ClearBit64(value, pos)
    if err != nil || ret != uint64(1) << 63{
        t.Logf("MSB error")
        t.Fail()
    }

}

func TestTestBit32(t *testing.T) {
    var value uint32
    var ret bool
    var pos uint
    var err error

    value = 0x80000001

    pos = 32
    ret, err = TestBit32(value, pos)
    if err == nil {
        t.Logf("expect error")
        t.Fail()
    }

    pos = 31
    ret, err = TestBit32(value, pos)
    if err != nil || !ret {
        t.Logf("LSB error")
        t.Fail()
    }

    pos = 0
    ret, err = TestBit32(value, pos)
    if err != nil || !ret {
        t.Logf("MSB error")
        t.Fail()
    }

    pos = 16
    ret, err = TestBit32(value, pos)
    if err != nil || ret {
        t.Logf("16th bit error")
        t.Fail()
    }

}

func TestTestBit64(t *testing.T) {
    var value uint64
    var ret bool
    var pos uint
    var err error

    value = 0x8000000000000001

    pos = 64
    ret, err = TestBit64(value, pos)
    if err == nil {
        t.Logf("expect error")
        t.Fail()
    }

    pos = 63
    ret, err = TestBit64(value, pos)
    if err != nil || !ret {
        t.Logf("LSB error")
        t.Fail()
    }

    pos = 0
    ret, err = TestBit64(value, pos)
    if err != nil || !ret{
        t.Logf("MSB error")
        t.Fail()
    }

    pos = 32
    ret, err = TestBit64(value, pos)
    if err != nil || ret {
        t.Logf("32nd bit error")
        t.Fail()
    }

}

func TestReverse32(t *testing.T) {
    var value, ret, expect uint32

    value = 0x01234567
    expect = 0xE6A2C480

    ret = Reverse32(value)
    if ret != expect {
        t.Fail()
        t.Logf("%x expect %x but get %x", value, expect, ret)
    }

    value = 0x89ABCDEF
    expect = 0xF7B3D591
    ret = Reverse32(value)
    if ret != expect {
        t.Fail()
        t.Logf("%x expect %x but get %x", value, expect, ret)
    }
}

func TestReverse64(t *testing.T) {
    var value, ret, expect uint64

    value = 0x0123456789ABCDEF
    expect = 0xF7B3D591E6A2C480

    ret = Reverse64(value)
    if ret != expect {
        t.Fail()
        t.Logf("%x expect %x but get %x", value, expect, ret)
    }
}

func TestRotateRight32(t *testing.T) {
    var value, ret, expect uint32
    var shift uint

    value = 0x01234567
    expect = 0x45670123
    shift = 16

    ret = RotateRight32(value, shift)
    if ret != expect {
        t.Fail()
        t.Logf("%x expect %x but get %x", value, expect, ret)
    }

    value = 0x01234567
    expect = 0x45670123
    shift = 48

    ret = RotateRight32(value, shift)
    if ret != expect {
        t.Fail()
        t.Logf("%x expect %x but get %x", value, expect, ret)
    }

    value = 0x01234567
    expect = 0x70123456
    shift = 4

    ret = RotateRight32(value, shift)
    if ret != expect {
        t.Fail()
        t.Logf("%x expect %x but get %x", value, expect, ret)
    }
}

func TestRotateLeft32(t *testing.T) {
    var value, ret, expect uint32
    var shift uint

    value = 0x01234567
    expect = 0x45670123
    shift = 16

    ret = RotateLeft32(value, shift)
    if ret != expect {
        t.Fail()
        t.Logf("%x expect %x but get %x", value, expect, ret)
    }

    value = 0x01234567
    expect = 0x45670123
    shift = 48

    ret = RotateLeft32(value, shift)
    if ret != expect {
        t.Fail()
        t.Logf("%x expect %x but get %x", value, expect, ret)
    }

    value = 0x01234567
    expect = 0x12345670
    shift = 4

    ret = RotateLeft32(value, shift)
    if ret != expect {
        t.Fail()
        t.Logf("%x expect %x but get %x", value, expect, ret)
    }
}

func TestRotateRight64(t *testing.T) {
    var value, ret, expect uint64
    var shift uint

    value = 0x0123456789ABCDEF
    expect = 0x89ABCDEF01234567
    shift = 32

    ret = RotateRight64(value, shift)
    if ret != expect {
        t.Fail()
        t.Logf("%x expect %x but get %x", value, expect, ret)
    }

    value = 0x0123456789ABCDEF
    expect = 0x89ABCDEF01234567
    shift = 96

    ret = RotateRight64(value, shift)
    if ret != expect {
        t.Fail()
        t.Logf("%x expect %x but get %x", value, expect, ret)
    }

    value = 0x0123456789ABCDEF
    expect = 0xF0123456789ABCDE
    shift = 4

    ret = RotateRight64(value, shift)
    if ret != expect {
        t.Fail()
        t.Logf("%x expect %x but get %x", value, expect, ret)
    }
}

func TestRotateLeft64(t *testing.T) {
    var value, ret, expect uint64
    var shift uint

    value = 0x0123456789ABCDEF
    expect = 0x89ABCDEF01234567
    shift = 32

    ret = RotateLeft64(value, shift)
    if ret != expect {
        t.Fail()
        t.Logf("%x expect %x but get %x", value, expect, ret)
    }

    value = 0x0123456789ABCDEF
    expect = 0x89ABCDEF01234567
    shift = 96

    ret = RotateLeft64(value, shift)
    if ret != expect {
        t.Fail()
        t.Logf("%x expect %x but get %x", value, expect, ret)
    }

    value = 0x0123456789ABCDEF
    expect = 0x123456789ABCDEF0
    shift = 4

    ret = RotateLeft64(value, shift)
    if ret != expect {
        t.Fail()
        t.Logf("%x expect %x but get %x", value, expect, ret)
    }
}
