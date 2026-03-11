package sbe

import (
	"math"
)

/*
 * This file contains composite types in CME Group ilinkbinary.xml file
 * Composite types from line 75 to 114
 */

/*
 * -- DATA Composite Type ---
 * line 75
 */

type DATA struct{
	Length UInt16
	VarData []byte
}

// function to create a new DATA via 
func NewDATA(data []byte) DATA{
	return DATA{
		Length: UInt16(len(data)),
		VarData: data,
	}
}

// function to convert data to a string in Go
func (d DATA) ToString() string{
	return string(d.VarData)
}

// function to check if data is empty, for slice []byte, check len() == 0
func (d DATA) isNULL() bool{
	return d.Length == 0 || len(d.VarData) == 0
}

/*
 * -- Decimal32NULL Composite Type ---
 * line 79
 */

type Decimal32NULL struct{
	Mantissa Int32NULL			// presence="optional" nullValue="2147483647" primitiveType="int32" gives Int32NULL
	Exponent Int8NULL				// presence="optional" nullValue="127" primitiveType="int8" gives Int8NULL
}

// IsNULL method to check if a Dcimal32NULL is null
func (d Decimal32NULL) IsNULL() bool{
	return d.Mantissa == Int32NULLValue || d.Exponent == Int8NULLValue
}

// SBE decimal encoding
// value = mantissa × 10^exponent
// e.g. 12345.67 
// mantissa = 1234567
// exponent = -2
// 12345.67 = mantisa * 10^(exponent)
// float64 binary is not reliable for transmission
// instead transmit mantissa and exponent
// note func math.Pow10() float64

// ToFloat64 returns the result as float64
func (d Decimal32NULL) ToFloat64() float64{
	return float64(d.Mantissa) * math.Pow10(int(d.Exponent));
}

/*
 * -- Decimal32NULL Composite Type ---
 * line 79
 */