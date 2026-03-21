package sbe

/*
* Defining enum types in CME Group ilinkbinary.xml file
* line 115 to 388
 */

/*
* AvgPxInd enum: uInt8NULL
* Line 115
 */

type AvgPxInd UInt8NULL

const(
	// No Average Pricing
	NoAveragePricing AvgPxInd = 0
	// Trade is part of an Average Price Group Identified by the AvgPxGrp ID
	TradeispartofanAveragePriceGroupIdentifiedbytheAvgPxGrpID AvgPxInd = 1
	// Trade is part of a Notional Value Average Price Group
	TradeispartofaNotionalValueAveragePriceGroup AvgPxInd = 3
	// UInt8NULL has value 255, used for checking IsNULL()
	AvgPxIndNULL AvgPxInd = 255
)

// Returns if an AvgPxInd is NULL
func (p AvgPxInd) IsNULL() bool{
	return p == AvgPxIndNULL
}

/*
* BooleanFlag, uInt8
* Line 120
*/

type BooleanFlag UInt8

const(
	// Use Go's iota
	BooleanFlagFalse BooleanFlag = iota
	BooleanFlagTrue
)

/*
* BooleanNULL, uInt8NULL
* Line 124
*/

type BooleanNULL UInt8NULL

const(
	BooleanNULLFalse BooleanNULL = 0
	BooleanNULLTrue BooleanNULL = 1
	BooleanNULLValue BooleanNULL = 255 // UInt8NULL is 255
)

// Returns true if BooleanNULL is BooleanNULLValue / null
func (b BooleanNULL) IsNULL() bool{
	return b == BooleanNULLValue
}

/*
* ClearingAcctType, enumNULL
* Line 128
*/

type ClearingAcctType EnumNULL

const(
	ClearingAcctTypeCustomer ClearingAcctType = 0
	ClearingAcctTypeFirm ClearingAcctType = 1
	ClearingAcctTypeNULL ClearingAcctType = 255 // from EnumNULLValue 255
)

// Returns true if ClearingAcctType is null
func (acctType ClearingAcctType) IsNULL() bool{
	return acctType == ClearingAcctTypeNULL
}

