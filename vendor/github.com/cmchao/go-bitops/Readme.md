# Introduction
The package implement a set of common bit operations which are widely used in conventional C/C++. The some function of this library should be a little slower than native C implementation. It is because C language prefer to use assert to check invalid parameter (ex : clear 100th bit for a 32bit variable) but this implementation check all possible error and return them.

# Feature List
| Function Prefix  | uint64 | uint32 | uint16 | uint8 | return error |
| -----------------|--------|--------|--------|-------|--------------|
| ClearBit         |   x    |   x    |        |       |       x      |
| ToggleBit        |   x    |   x    |        |       |       x      |
| SetBit           |   x    |   x    |        |       |       x      |
| TestBit          |   x    |   x    |        |       |       x      |
| CountLeadOne     |   x    |   x    |        |       |              |
| CountLeadZero    |   x    |   x    |        |       |              |
| CountTrailOne    |   x    |   x    |        |       |              |
| CountTrailZero   |   x    |   x    |        |       |              |
| CountOne         |   x    |   x    |   x    |   x   |              |
| CountZero        |   x    |   x    |   x    |   x   |              |
| Deposit          |   x    |   x    |        |       |       x      |
| Extract          |   x    |   x    |        |       |       x      |
| GetField         |   x    |   x    |        |       |       x      |
| SetField         |   x    |   x    |        |       |       x      |
| Reverse          |   x    |   x    |        |       |              |
| Rotate           |   x    |   x    |        |       |              |

# API Reference
https://godoc.org/github.com/cmchao/go-bitops

# Auto-build
[go-bitops at Travis](https://travis-ci.org/cmchao/go-bitops)
