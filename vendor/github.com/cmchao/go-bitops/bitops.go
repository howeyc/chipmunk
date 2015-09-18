/*
    Package bitops provide a set of bit operations for uint32/uint64. They are similar to 
    what you can find in conventional C library.
*/
package bitops

import "fmt"

// real implementation for Extract32 and GetField32
func extract32(value uint32, start uint, length uint) (uint32, error) {
    return (value >> start) & ( ^uint32(0) >> (32 - length)), nil
}

// real implementation for Extract64 and GetField64
func extract64(value uint64, start uint, length uint) (uint64, error) {
    return (value >> start) & ( ^uint64(0) >> (64 - length)), nil
}

// Extract32 specify field from uint32 by staring position and length
// LSB/MSB are 0/31 and return original value if error occurs
func Extract32(value uint32, start uint, length uint) (uint32, error) {
    if (start > 31  ||  length >  32 - start ) {
        return value, fmt.Errorf("invalid start(%v) or length(%v)", start, length);
    }

    return extract32(value, start, length)
}

// Extract64 specify field from uint64 by starting position and length
// LSB/MSB are 0/63 and return original value if error occurs
func Extract64(value uint64, start uint, length uint) (uint64, error) {
    if (start > 63  ||  length >  64 - start ) {
        return value, fmt.Errorf("invalid start(%v) or length(%v)", start, length);
    }

    return extract64(value, start, length)
}

// GetField32 specify field between high and low bit from uint32 
// LSB/MSB are 0/31 and return original value if error occurs
func GetField32(value uint32, high uint, low uint) (uint32, error) {
    if (high > 31  || low > 31 || high < low) {
        return value, fmt.Errorf("invalid high(%v) or low(%v)", high, low);
    }

    return extract32(value, low, high - low + 1)
}

// GetField64 specify field between high and low bit from uint64 
// LSB/MSB are 0/63 and return original value if error occurs
func GetField64(value uint64, high uint, low uint) (uint64, error) {
    if (high > 63  || low > 63 || high < low) {
        return value, fmt.Errorf("invalid high(%v) or low(%v)", high, low);
    }

    return extract64(value, low, high - low + 1)
}

// real implementation for Depoit32 and SetField32
func deposit32(value uint32, start uint, length uint, field uint32) (uint32, error) {
    mask := (^uint32(0) >> (32 - length)) << start;
    return (value & ^mask) | ((field << start) & mask), nil;
}

// real implementation for Depoit64 and SetField64
func deposit64(value uint64, start uint, length uint, field uint64) (uint64, error) {
    mask := (^uint64(0) >> (64 - length)) << start;
    return (value & ^mask) | ((field << start) & mask), nil;
}

// Deposit32 specified field to uint32  variable by staring position and length
// LSB/MSB are 0/31 and return original value if error occurs
func Deposit32(value uint32, start uint, length uint, field uint32) (uint32, error) {
    if start > 31 || length > 32 - start {
        return value, fmt.Errorf("invalid start(%v) or length(%v)", start, length);
    }

    return deposit32(value, start, length, field)
}

// Deposit64 specified field to uint64  variable by staring position and length
// LSB/MSB are 0/63 and return original value if error occurs
func Deposit64(value uint64, start uint, length uint, field uint64) (uint64, error) {
    if start > 63 || length > 64 - start {
        return value, fmt.Errorf("invalid start(%v) or length(%v)", start, length);
    }

    return deposit64(value, start, length, field)
}

// SetField32 specified field to uint32 variable by staring position and length
// LSB/MSB are 0/31 and return original value if error occurs
func SetField32(value uint32, high uint, low uint, field uint32) (uint32, error) {
    if high > 31  || low > 31 || high < low {
        return value, fmt.Errorf("invalid high(%v) or low(%v)", high, low);
    }

    return deposit32(value, low, high - low + 1, field)
}

// SetField64 specified field to uint64 variable by staring position and length
// LSB/MSB are 0/63 and return original value if error occurs
func SetField64(value uint64, high uint, low uint, field uint64) (uint64, error) {
    if high > 63  || low > 63 || high < low {
        return value, fmt.Errorf("invalid high(%v) or low(%v)", high, low);
    }

    return deposit64(value, low, high - low + 1, field)
}

// CountOne8 return number of 1 in uint8 variable
func CountOne8(value uint8) (uint) {
    value = (value & 0x55) + ((value >> 1) & 0x55);
    value = (value & 0x33) + ((value >> 2) & 0x33);
    value = (value & 0x0f) + ((value >> 4) & 0x0f);

    return uint(value);
}

// CountOne16 return number of 1 in uint16 variable
func CountOne16(value uint16) (uint) {
    value = (value & 0x5555) + ((value >> 1) & 0x5555);
    value = (value & 0x3333) + ((value >> 2) & 0x3333);
    value = (value & 0x0f0f) + ((value >> 4) & 0x0f0f);
    value = (value & 0x00ff) + ((value >> 8) & 0x00ff);

    return uint(value);
}

// CountOne32 return number of 1 in uint32 variable
func CountOne32(value uint32) (uint) {
    value = (value & 0x55555555) + ((value >>  1) & 0x55555555);
    value = (value & 0x33333333) + ((value >>  2) & 0x33333333);
    value = (value & 0x0f0f0f0f) + ((value >>  4) & 0x0f0f0f0f);
    value = (value & 0x00ff00ff) + ((value >>  8) & 0x00ff00ff);
    value = (value & 0x0000ffff) + ((value >> 16) & 0x0000ffff);

    return uint(value);
}

// CountOne64 return number of 1 in uint64 variable
func CountOne64(value uint64) (uint) {
    value = (value & 0x5555555555555555) + ((value >>  1) & 0x5555555555555555);
    value = (value & 0x3333333333333333) + ((value >>  2) & 0x3333333333333333);
    value = (value & 0x0f0f0f0f0f0f0f0f) + ((value >>  4) & 0x0f0f0f0f0f0f0f0f);
    value = (value & 0x00ff00ff00ff00ff) + ((value >>  8) & 0x00ff00ff00ff00ff);
    value = (value & 0x0000ffff0000ffff) + ((value >> 16) & 0x0000ffff0000ffff);
    value = (value & 0x00000000ffffffff) + ((value >> 32) & 0x00000000ffffffff);

    return uint(value);
}

// CountOne8 return number of 0 in uint8 variable
func CountZero8(value uint8) (uint) {
    return 8 - CountOne8(value)
}

// CountOne16 return number of 0 in uint16 variable
func CountZero16(value uint16) (uint) {
    return 16 - CountOne16(value)
}

// CountOne32 return number of 0 in uint32 variable
func CountZero32(value uint32) (uint) {
    return 32 - CountOne32(value)
}

// CountOne64 return number of 0 in uint64 variable
func CountZero64(value uint64) (uint) {
    return 64 - CountOne64(value)
}

// CountTrailZero32 return number of trailing zero in a 32-bit value
func CountTrailZero32(value uint32) (uint) {
    var count uint = 0

    if (value & 0x0000FFFF) == 0 {
        count += 16;
        value >>= 16;
    }
    if (value & 0x000000FF) == 0 {
        count += 8;
        value >>= 8;
    }
    if (value & 0x0000000F) == 0 {
        count += 4;
        value >>= 4;
    }
    if (value & 0x00000003) == 0 {
        count += 2;
        value >>= 2;
    }
    if (value & 0x00000001) == 0 {
        count++;
        value >>= 1;
    }
    if (value & 0x00000001) == 0 {
        count++;
    }

    return count
}

// CountTrailZero64 return number of trailing zero in a 32-bit value
func CountTrailZero64(value uint64) (uint) {
    var count uint = 0

    if  uint32(value) == 0 {
        count += 32;
        value >>= 32;
    }

    return count + CountTrailZero32(uint32(value))
}

// CountTrailOne32 return number of trailing 1 in a 32-bit value
func CountTrailOne32(value uint32) (uint) {
    return CountTrailZero32(^value)
}

// CountTrailOne64 return number of trailing 1 in a 32-bit value
func CountTrailOne64(value uint64) (uint) {
    return CountTrailZero64(^value)
}

// CountLeadZero32 return number of leading 0 in a 32-bit value
func CountLeadZero32(value uint32) (uint) {
    var count uint = 0

    if (value & 0xFFFF0000) == 0 {
        count += 16;
        value <<= 16;
    }
    if (value & 0xFF000000) == 0 {
        count += 8;
        value <<= 8;
    }
    if (value & 0xF0000000) == 0 {
        count += 4;
        value <<= 4;
    }
    if (value & 0x30000000) == 0 {
        count += 2;
        value <<= 2;
    }
    if (value & 0x10000000) == 0 {
        count++;
        value <<= 1;
    }
    if (value & 0x10000000) == 0 {
        count++;
    }

    return count
}

// CountLeadZero64 return number of leading 0 in a 32-bit value
func CountLeadZero64(value uint64) (uint) {
    var count uint = 0

    if (value >> 32) == 0 {
        count += 32
    } else {
        value >>= 32;
    }

    return count + CountLeadZero32(uint32(value))
}

// CountLeadOne32 return number of leading 1 in a 32-bit value
func CountLeadOne32(value uint32) (uint) {
    return CountLeadZero32(^value)
}

// CountLeadOne64 return number of leading 1 in a 32-bit value
func CountLeadOne64(value uint64) (uint) {
    return CountLeadZero64(^value)
}

// SetBit32 set the specified bit to 1 for 32-bit value and return the new value
func SetBit32(value uint32, pos uint) (uint32, error) {
    if pos >= 32 {
        return value, fmt.Errorf("invalid position(%v)", pos)
    }

    return (value | (uint32(1) << pos)), nil
}

// SetBit64 set the specified bit to 1 for 64-bit value and return the new value
func SetBit64(value uint64, pos uint) (uint64, error) {
    if pos >= 64 {
        return value, fmt.Errorf("invalid position(%v)", pos)
    }
    return (value | (uint64(1) << pos)), nil
}

// ToggleBit32 set the specified bit to 1 for 32-bit value and return the new value
func ToggleBit32(value uint32, pos uint) (uint32, error) {
    if pos >= 32 {
        return value, fmt.Errorf("invalid position(%v)", pos)
    }

    return (value ^ (uint32(1) << pos)), nil
}

// ToggleBit64 set the specified bit to 1 for 64-bit value and return the new value
func ToggleBit64(value uint64, pos uint) (uint64, error) {
    if pos >= 64 {
        return value, fmt.Errorf("invalid position(%v)", pos)
    }
    return (value ^ (uint64(1) << pos)), nil
}

// ClearBit32 set the specified bit to 1 for 32-bit value and return the new value
func ClearBit32(value uint32, pos uint) (uint32, error) {
    if pos >= 32 {
        return value, fmt.Errorf("invalid position(%v)", pos)
    }

    return (value &^ (uint32(1) << pos)), nil
}

// ClearBit64 set the specified bit to 1 for 64-bit value and return the new value
func ClearBit64(value uint64, pos uint) (uint64, error) {
    if pos >= 64 {
        return value, fmt.Errorf("invalid position(%v)", pos)
    }
    return (value &^ (uint64(1) << pos)), nil
}

// TestBit32 set the specified bit to 1 for 32-bit value and return the new value
func TestBit32(value uint32, pos uint) (bool, error) {
    if pos >= 32 {
        return false, fmt.Errorf("invalid position(%v)", pos)
    }

    return (value & (uint32(1) << pos)) != 0, nil
}

// TestBit64 set the specified bit to 1 for 64-bit value and return the new value
func TestBit64(value uint64, pos uint) (bool, error) {
    if pos >= 64 {
        return false, fmt.Errorf("invalid position(%v)", pos)
    }
    return (value & (uint64(1) << pos)) != 0, nil
}

// Reverse32 set reverse the bit order for 32-bit variable
func Reverse32(value uint32) (uint32) {
    value = ((value >> 1) & 0x55555555) | ((value & 0x55555555) << 1)
    value = ((value >> 2) & 0x33333333) | ((value & 0x33333333) << 2)
    value = ((value >> 4) & 0x0F0F0F0F) | ((value & 0x0F0F0F0F) << 4)
    value = ((value >> 8) & 0x00FF00FF) | ((value & 0x00FF00FF) << 8)
    value = ( value >> 16             ) | ( value               << 16)

    return value
}

// Reverse64 set reverse the bit order for 64-bit variable
func Reverse64(value uint64) (uint64) {
    high :=  Reverse32(uint32(value))
    low :=  Reverse32(uint32(value >> 32))

    return (uint64(high) << 32) | uint64(low)
}

// RotateRight32 rotate an 32-bit value right
func RotateRight32(value uint32, shift uint) (uint32) {
    shift = shift & 0x1F

    return (value >> shift) | (value << (32 - shift))
}

// RotateLeft32 rotate an 32-bit value left
func RotateLeft32(value uint32, shift uint) (uint32){
    shift = shift & 0x1F

    return (value << shift) | (value >> (32 - shift))
}

// RotateRight64 rotate an 64-bit value right
func RotateRight64(value uint64, shift uint) (uint64){
    shift = shift & 0x3F

    return (value >> shift) | (value << (64 - shift))
}

// RotateLeft64 rotate an 64-bit value left
func RotateLeft64(value uint64, shift uint) (uint64){
    shift = shift & 0x3F

    return (value << shift) | (value >> (64 - shift))
}
