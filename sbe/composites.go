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
	return d.Mantissa.IsNULL() || d.Exponent.IsNULL()
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
 * -- Decimal64NULL Composite Type ---
 * line 83
 */

type Decimal64NULL struct{
	Mantissa Int64NULL // description="mantissa" presence="optional" nullValue="9223372036854775807" primitiveType="int64"
	Exponent Int8NULL
}

// IsNULL method to return if Decimal64NULL is null
func (d Decimal64NULL) IsNULL() bool{
	return d.Mantissa.IsNULL() || d.Exponent.IsNULL()
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

// function isNULL returns true only when all attributes are null
func (m MaturityMonthYear) IsNULL() bool{
	return m.Year.IsNULL() && m.Month.IsNULL() && m.Day.IsNULL() && m.Week.IsNULL()
}

// ToTime method returns time.Time and bool
func(m MaturityMonthYear) ToTime() (time.Time, bool){
	// Maturity Month Year requires at least Year and Month to be present
	if(m.Year.IsNULL() || m.Month.IsNULL()){
		// return zero value of time.Time and boolean false
		return time.Time{}, false
	}

	day := 1
	// else, default day to 1 and set day
	if !m.Day.IsNULL() {
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
	if(m.Year.IsNULL() || m.Month.IsNULL()){
		return ""
	}
	if(!m.Day.IsNULL()){
		return fmt.Sprintf("%04d-%02d-%02d", m.Year, m.Month, m.Day)
	}
	if(!m.Week.IsNULL()){
		return fmt.Sprintf("%04d-%02d-W%02d-D%02d", m.Year, m.Month, m.Week, m.Day)
	}
	return fmt.Sprintf("%04d-%02d", m.Year, m.Month)
}