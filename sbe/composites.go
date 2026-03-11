package sbe

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

func (d Decimal32NULL) isNULL() bool{
	return d.Mantissa == Int32NULLValue || d.Exponent == Int8NULLValue
}
