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
	// UInt8NULL has value 255, used for checking IsNull()
	AvgPxIndNULL AvgPxInd = 255
)

// Returns if an AvgPxInd is NULL
func (p AvgPxInd) IsNull() bool{
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
func (b BooleanNULL) IsNull() bool{
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
func (acctType ClearingAcctType) IsNull() bool{
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

func (c CmtaGiveUpCD) IsNull() bool{
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

func (cohi CustOrdHandInst) IsNull() bool {
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

func (c CustOrderCapacity) IsNull() bool{
	return c == CustOrderCapacityNULL
}

/*
* DKReason, charNULL
* Line 155
*/

type DKReason CHAR

const(
	// Unknown Security
	DKReasonUnknownSecurity							DKReason = 'A'
	// Wrong Side
	DKReasonWrongSide										DKReason = 'B'
	// Quantity Exceeds Order
	DKReasonQuantityExceedsOrder				DKReason = 'C'
	// No Matching Order
	DKReasonNoMatchingOrder							DKReason = 'D'
	// Price Exceeds Limit
	DKReasonPriceExceedsLimit						DKReason = 'E'
	// Calculation Difference
	DKReasonCalculationDifference				DKReason = 'F'
	// No Matching Execution Report
	DKReasonNoMatchingExecutionReport		DKReason = 'G'
	// Other
	DKReasonOther												DKReason = 'Z'
	// DKReason NULL value
	DKReasonNULL												DKReason = 0
)

func (d DKReason) IsNull() bool{
	return d == DKReasonNULL
}

/*
* ExecAckStatus, uInt8
* Line 165
* Both const values are required as uInt8, no null required
*/

type ExecAckStatus UInt8

const(
	// Accepted
	ExecAckStatusAccepted ExecAckStatus = 1
	// Rejected
	ExecAckStatusRejected ExecAckStatus = 2
)

/*
* ExecMode, charNULL
* Line 169
*/

type ExecMode CHAR

const(
	// Aggressive
	ExecModeAggressive = 'A'
	// Passive
	ExecModePassive = 'P'
	// null value because of charNULL type
	ExecModeNULL = 0
)

func (e ExecMode) IsNull() bool{
	return e == ExecModeNULL
}

/*
* ExecReason, uInt8NULL
* Line 173
* uInt8NULL is optional and requires IsNull func
*/

type ExecReason UInt8NULL

const(
	// Market Exchange Option
	ExecReasonMarketExchangeOption										ExecReason = 8
	// Cancelled Not Best
	ExecReasonCancelledNotBest												ExecReason = 9
	// Cancel on Disconnect
	ExecReasonCancelOnDisconnect											ExecReason = 100
	// Self Match Prevention Oldest Order Cancelled
	ExecReasonSelfMatchPreventionOldestOrderCancelled ExecReason = 103
	// Cancel On Globex Credit Controls Violation
	ExecReasonCancelOnGlobexCreditControlsViolation 	ExecReason = 104
	// Cancel From Firmsoft
	ExecReasonCancelFromFirmsoft										  ExecReason = 105
	// Cancel From Risk Management API 
	ExecReasonCancelFromRiskManagementAPI 						ExecReason = 106
	// Self Match Prevention Newest Order Cancelled
	ExecReasonSelfMatchPreventionNewestOrderCancelled ExecReason = 107
	// Cancel due to min qty (Cancelduetovolquotedoptionorderrestedqtylessthanminordersize)
	ExecReasonCancelDueToVolLessThanNinOrderSize		  ExecReason = 108
	// Cancel RFC Order
	ExecReasonCancelRFCOrder 													ExecReason = 109
	// Cancel Upon Contract Expiration
	ExecReasonCancelUponContractExpiration	 					ExecReason = 110
	// System Cancel
	ExecReasonSystemCancel 														ExecReason = 111
	// ExecReason NULL value (uInt8NUll)
	ExecReasonNULL																		ExecReason = 255
)

// IsNull returns true if the given value is Null
func (e ExecReason) IsNull() bool{
	return e == ExecReasonNULL
}

/*
* ExecTypTrdCxl, CHAR
* Line 187
*/

type ExecTypTrdCxl CHAR

const(
	// Trade Correction
	ExecTypTrdCxlTradeCorrection			ExecTypTrdCxl = 'G'
	// Trade Cancel
	ExecTypTrdCxlTradeCancel					ExecTypTrdCxl = 'H'
)

/*
* ExpCycle, uInt8NULL
* Line 191
*/

type ExpCycle UInt8NULL

const(
	// Expire On Trading Session Close
	ExpCycleExpireOnTradingSessionClose ExpCycle = 0
	// Expiration at given date
	ExpCycleExpirationatgivendate				ExpCycle = 2
	// ExpCycle Null value
	ExpCycleNULL												ExpCycle = 255
)

// IsNull returns true if the ExpCycle value is Null
func (e ExpCycle) IsNull() bool{
	return e == ExpCycleNULL
}

/*
* FTI, uInt8NULL
* Line 195
*/

type FTI UInt8NULL

const(
	// Backup
	FTIBackup		FTI = 0
	FTIPrimary	FTI = 1
	FTINull			FTI = 255
)

// IsNull returns true if the FTI value is Null
func (f FTI) IsNull() bool{
	return f == FTINull
}

