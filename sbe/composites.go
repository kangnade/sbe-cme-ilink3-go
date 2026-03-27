package sbe

import (
	"fmt"
	"math"
	"time"
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
func (d DATA) IsNull() bool{
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

// IsNull method to check if a Dcimal32NULL is null
func (d Decimal32NULL) IsNull() bool{
	return d.Mantissa.IsNull() || d.Exponent.IsNull()
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
	return float64(d.Mantissa) * math.Pow10(int(d.Exponent))
}

/*
 * -- Decimal64NULL Composite Type ---
 * line 83
 */

type Decimal64NULL struct{
	Mantissa Int64NULL // description="mantissa" presence="optional" nullValue="9223372036854775807" primitiveType="int64"
	Exponent Int8NULL
}

// IsNull method to return if Decimal64NULL is null
func (d Decimal64NULL) IsNull() bool{
	return d.Mantissa.IsNull() || d.Exponent.IsNull()
}

// ToFloat64 returns the result as float64
func (d Decimal64NULL) ToFloat64() float64{
	return float64(d.Mantissa) * math.Pow10(int(d.Exponent))
}

/*
 * -- MaturityMonthYear Composite Type ---
 * line 87
 * description: Year, Month and Date
 */

type MaturityMonthYear struct{
	Year UInt16NULL
	Month UInt8NULL
	Day UInt8NULL
	Week UInt8NULL
}

// function IsNull returns true only when all attributes are null
func (m MaturityMonthYear) IsNull() bool{
	return m.Year.IsNull() && m.Month.IsNull() && m.Day.IsNull() && m.Week.IsNull()
}

// ToTime method returns time.Time and bool
func(m MaturityMonthYear) ToTime() (time.Time, bool){
	// Maturity Month Year requires at least Year and Month to be present
	if(m.Year.IsNull() || m.Month.IsNull()){
		// return zero value of time.Time and boolean false
		return time.Time{}, false
	}

	day := 1
	// else, default day to 1 and set day
	if !m.Day.IsNull() {
		day = int(m.Day)
	}

	return time.Date(
		int(m.Year),
		time.Month(m.Month),
		day,
		0, 0, 0, 0,
		time.UTC,
	), true
}

// ToString function to return human readable fotmat for logging
func (m MaturityMonthYear)  ToString() string{
	if(m.Year.IsNull() || m.Month.IsNull()){
		return ""
	}
	if(!m.Day.IsNull()){
		return fmt.Sprintf("%04d-%02d-%02d", m.Year, m.Month, m.Day)
	}
	if(!m.Week.IsNull()){
		return fmt.Sprintf("%04d-%02d-W%02d", m.Year, m.Month, m.Week)
	}
	return fmt.Sprintf("%04d-%02d", m.Year, m.Month)
}

/*
 * -- PRICE9 Composite Type ---
 * line 93
 * description: Price with constant exponent -9
 */

// The exponent is always constant -9 for both PRICE9 and PRICENULL9
const PRICE9Exponent Int8 = -9

// Price9 struct definition
type PRICE9 struct{
	Mantissa Int64
}

// ToFloat64 function for Price9 composite type
func (p PRICE9) ToFloat64() float64{
	return float64(p.Mantissa) * math.Pow10(int(PRICE9Exponent))
}

/*
 * -- PRICENULL9 Composite Type ---
 * line 97
 * description: Optional price with constant exponent -9
 * since the mantissa is optional, we need IsNull to check if it is null
 */

 // Exponent is -9, reuse above from PRICE9
 type PRICENULL9 struct{
	// The Mantissa is Int64NULL
	Mantissa Int64NULL
 }

// IsNull function for PRICENULL9
func (p PRICENULL9) IsNull() bool{
	return p.Mantissa.IsNull()
}

// ToFloat64 for PRICENULL9 type
func (p PRICENULL9) ToFloat64() float64{
	return float64(p.Mantissa) * math.Pow10(int(PRICE9Exponent))
}

/*
 * -- GroupSize Composite Type ---
 * line 101
 * description: Repeating group dimensions
 */

 type GroupSize struct{
	BlockLength UInt16		// length of each group
	NumInGroup UInt8		// number of entries in the group
 }

 /*
 * -- GroupSizeEncoding Composite Type ---
 * line 105
 * description: Repeating group dimensions, with uint16 for NumInGroup
 */
 
type GroupSizeEncoding struct{
	BlockLength UInt16
	NumInGroup UInt16
}

 /*
 * -- MessageHeader Composite Type ---
 * line 109
 * description: Template ID and length of message root
 */

 type MessageHeader struct{
	BlockLength UInt16
	TemplateId UInt16
	SchemaId UInt16
	Version UInt16
 }