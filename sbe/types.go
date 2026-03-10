package sbe

/*
Golang primitive types:
=======================
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
=======================
*/

/*
* Types are marked with lines in ilinkbinary.xml file
* and defined in the following way:
* named types in Golang: type namedType primitiveType
*                        const namedTypeNULL namedType = nullValue
*                        function for const value optional
 */

// ilinkbinary.xml line 67 to 74 types in the following way:
type UInt16 uint16
const UInt16NULL UInt16 = 65535

type UInt32 uint32
const UInt32NULL UInt32 = 4294967295

type UInt64 uint64
const UInt64NULL UInt64 = 18446744073709551615

type UInt8 uint8
const UInt8NULL UInt8 = 255

// ilinkbinary.xml line 66 type
type EnumNULL uint8
const EnumNULLValue UInt8 = 255

// ilinkbinary.xml line 4 to line 20

// line 4
type CHAR byte

// line 5
type ClientFlowType [10]byte
// Use array composite literal
var clientFlowTypeValue = ClientFlowType{'I', 'D', 'E', 'M', 'P', 'O', 'T', 'E', 'N', 'T'}
// Function to get ClientFlowType const value
func GetClientFlowType() ClientFlowType{
     return clientFlowTypeValue
}

// line 6, description: Cross order type supports only limit order
type CrossOrderType byte
const CrossOrderTypeValue CrossOrderType = '2'
func GetCrossOrderType() CrossOrderType{
     return CrossOrderTypeValue
}

