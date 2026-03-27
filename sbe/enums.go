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

/*
* CmtaGiveUpCD, charNULL
* Line 132
*/

type CmtaGiveUpCD CHAR

const(
	CmtaGiveUpCDGiveUp CmtaGiveUpCD = 'G'			// give up
	CmtaGiveUpCDSGXoffset CmtaGiveUpCD = 'S'	// SGX offset
	CmtaGiveUpCDNULL = 0
)

func (c CmtaGiveUpCD) IsNULL() bool{
	return c == CmtaGiveUpCDNULL
}

/*
* CrossTypeEnum, uInt8
* Line 136
*/

type CrossTypeEnum UInt8

const(
	CrossTypeCross CrossTypeEnum = 3 		// default product configuration
	CrossTypeRCross CrossTypeEnum = 20		// RFQ + RFC Cross
	CrossTypeCCross CrossTypeEnum = 21		// Committed Cross
)

/*
* CustOrdHandInst, charNULL
* Line 141
*/

type CustOrdHandInst CHAR

const(
	CustOrdHandInstFCMProvidedScreen CustOrdHandInst = 'C'			// FCm provided screen
	CustOrdHandInstOtherProvidedScreen CustOrdHandInst = 'D'		// Other provided screen
	CustOrdHandInstFCMAPIorFix CustOrdHandInst = 'G'						// FCM API or FIX
	CustOrdHandInstAlgoEngine CustOrdHandInst = 'H'							// Algo Engine
	CustOrdHandInstDeskElectronic CustOrdHandInst = 'W'					// Desk Electronic
	CustOrdHandInstClientElectronic CustOrdHandInst = 'Y'				// Client Electronic
	CustOrdHandInstNULL CustOrdHandInst = 0
)

func (cohi CustOrdHandInst) IsNULL() bool {
	return cohi == CustOrdHandInstNULL
}

/*
* CustOrdHandInst, charNULL
* Line 149
*/

type CustOrderCapacity EnumNULL

const(
	// Member trading for their own account
	CustOrderCapacityMemberTradingForTheirOwnAccount CustOrderCapacity = 1
	// Member firm trading for its proprietary account
	CustOrderCapacityMemberFirmTradingForItsProprietaryAccount CustOrderCapacity = 2
	// Member trading for another member or non member
	CustOrderCapacityMemberTradingForAnotherMemberOrNonmember = 3
	// All other
	CustOrderCapacityAllOther = 4
	// CustOrderCapacity NULL value
	CustOrderCapacityNULL CustOrderCapacity = 255
)

func (c CustOrderCapacity) IsNULL() bool{
	return c == CustOrderCapacityNULL
}

/*
* CustOrdHandInst, charNULL
* Line 155
*/

