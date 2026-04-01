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
	AvgPxIndNoAveragePricing 																					AvgPxInd = 0
	// Trade is part of an Average Price Group Identified by the AvgPxGrp ID
	AvgPxIndTradeispartofanAveragePriceGroupIdentifiedbytheAvgPxGrpID AvgPxInd = 1
	// Trade is part of a Notional Value Average Price Group
	AvgPxIndTradeispartofaNotionalValueAveragePriceGroup 							AvgPxInd = 3
	// UInt8NULL has value 255, used for checking IsNull()
	AvgPxIndNULL 																							AvgPxInd = 255
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
	BooleanFlagFalse 	BooleanFlag = 0
	BooleanFlagTrue		BooleanFlag = 1
)

/*
* BooleanNULL, uInt8NULL
* Line 124
*/

type BooleanNULL UInt8NULL

const(
	BooleanNULLFalse 	BooleanNULL = 0
	BooleanNULLTrue 	BooleanNULL = 1
	BooleanNULLValue 	BooleanNULL = 255 // UInt8NULL is 255
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
	ClearingAcctTypeCustomer 	ClearingAcctType = 0
	ClearingAcctTypeFirm 			ClearingAcctType = 1
	ClearingAcctTypeNULL 			ClearingAcctType = 255 // from EnumNULLValue 255
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
	CmtaGiveUpCDGiveUp 		CmtaGiveUpCD = 'G'			// give up
	CmtaGiveUpCDSGXoffset CmtaGiveUpCD = 'S'	// SGX offset
	CmtaGiveUpCDNULL 			CmtaGiveUpCD= 0
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
	CrossTypeCross 	CrossTypeEnum = 3 		// default product configuration
	CrossTypeRCross CrossTypeEnum = 20		// RFQ + RFC Cross
	CrossTypeCCross CrossTypeEnum = 21		// Committed Cross
)

/*
* CustOrdHandInst, charNULL
* Line 141
*/

type CustOrdHandlInst CHAR

const(
	CustOrdHandlInstFCMProvidedScreen 		CustOrdHandlInst = 'C'			// FCm provided screen
	CustOrdHandlInstOtherProvidedScreen 	CustOrdHandlInst = 'D'		// Other provided screen
	CustOrdHandlInstFCMAPIorFix 					CustOrdHandlInst = 'G'						// FCM API or FIX
	CustOrdHandlInstAlgoEngine 					CustOrdHandlInst = 'H'							// Algo Engine
	CustOrdHandlInstDeskElectronic 			CustOrdHandlInst = 'W'					// Desk Electronic
	CustOrdHandlInstClientElectronic 		CustOrdHandlInst = 'Y'				// Client Electronic
	CustOrdHandlInstNULL 								CustOrdHandlInst = 0
)

func (cohi CustOrdHandlInst) IsNull() bool {
	return cohi == CustOrdHandlInstNULL
}

/*
* CustOrdCapacity, enumNULL
* Line 149
*/

type CustOrderCapacity EnumNULL

const(
	// Member trading for their own account
	CustOrderCapacityMemberTradingForTheirOwnAccount 						CustOrderCapacity = 1
	// Member firm trading for its proprietary account
	CustOrderCapacityMemberFirmTradingForItsProprietaryAccount 	CustOrderCapacity = 2
	// Member trading for another member or non member
	CustOrderCapacityMemberTradingForAnotherMemberOrNonmember 	CustOrderCapacity= 3
	// All other
	CustOrderCapacityAllOther 																	CustOrderCapacity = 4
	// CustOrderCapacity NULL value
	CustOrderCapacityNULL 																			CustOrderCapacity = 255
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
	ExecModeAggressive 	ExecMode = 'A'
	// Passive
	ExecModePassive 		ExecMode = 'P'
	// null value because of charNULL type
	ExecModeNULL 				ExecMode = 0
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
	ExpCycleExpirationAtGivenDate				ExpCycle = 2
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
	FTINULL			FTI = 255
)

// IsNull returns true if the FTI value is Null
func (f FTI) IsNull() bool{
	return f == FTINULL
}

/*
* KeepAliveLapsed, uInt8
* Line 199
*/

type KeepAliveLapsed UInt8

const(
	// Not Lapsed
	KeepAliveLapsedNotLapsed KeepAliveLapsed = 0
	// Lapsed
	KeepAliveLapsedLapsed 	KeepAliveLapsed = 1
)

/*
* ListUpdAct, CHAR
* Line 203
*/

type ListUpdAct CHAR

const(
	// Add
	ListUpdActAdd 		ListUpdAct = 'A'
	// Delete
	ListUpdActDelete	ListUpdAct = 'D'
)

/*
* ManualOrdInd, enumNULL
* Line 207
*/

type ManualOrdInd EnumNULL

const(
	// Automated
	ManualOrdIndAutomated			ManualOrdInd = 0
	// Manual
	ManualOrdIndManual				ManualOrdInd = 1
	// ManualOrdInd Null val
	ManualOrdIndNULL					ManualOrdInd = 255
)

// Returns true if the ManualOrdInd value is NULL
func (m ManualOrdInd) IsNull() bool{
	return m == ManualOrdIndNULL
}

/*
* ManualOrdIndReq, uInt8
* Line 211
*/

type ManualOrdIndReq UInt8

const(
	// Automated
	ManualOrdIndReqAutomated 	ManualOrdIndReq = 0
	// Manual
	ManualOrdIndReqManual			ManualOrdIndReq = 1
)

/*
* MassActionOrdTyp, charNULL
* Line 215
*/

type MassActionOrdTyp CHAR

const(
	// Limit
	MassActionOrdTypLimit				MassActionOrdTyp = '2'
	// Stop Limit
	MassActionOrdTypStopLimit		MassActionOrdTyp = '4'
	// MassActionOrdTyp NULL val
	MassActionOrdTypNULL				MassActionOrdTyp = 0
)

// Returns true if the MassActionOrdTyp is NULL
func (m MassActionOrdTyp) IsNull() bool{
	return m == MassActionOrdTypNULL
}

/*
* MassActionResponse, uInt8
* Line 219
*/

type MassActionResponse UInt8

const(
	// Rejected
	MassActionResponseRejected		MassActionResponse = 0
	// Accepted
	MassActionResponseAccepted		MassActionResponse = 1
)

/*
* MassActionScope, uInt8
* Line 223
*/

type MassActionScope UInt8

const(
	// Instrument
	MassActionScopeInstrument					MassActionScope = 1
	// All
	MassActionScopeAll								MassActionScope = 7
	// Market Segment ID
	MassActionScopeMarketSegmentID		MassActionScope = 9
	// Instrument Group
	MassActionScopeInstrumentGroup		MassActionScope = 10
	// Quote Set ID
	MassActionScopeQuoteSetID					MassActionScope = 100
)

/*
* MassCancelTIF, uInt8NULL
* Line 230
*/

type MassCancelTIF UInt8NULL

const(
	// Day
	MassCancelTIFDay							MassCancelTIF = 0
	// Good Till Cancel
	MassCancelTIFGoodTillCancel		MassCancelTIF = 1
	// Good Till Date
	MassCancelTIFGoodTillDate			MassCancelTIF = 6
	// MassCancelTIF NULL val
	MassCancelTIFNULL							MassCancelTIF = 255
)

// Returns true if the Mass Cancel TIF value is NULL
func (m MassCancelTIF) IsNull() bool{
	return m == MassCancelTIFNULL
}

/*
* MassCxlReqTyp, uInt8NULL
* Line 235
*/

type MassCxlReqTyp UInt8NULL

const(
	// Mass Cancel Request Type Sender Sub ID
	MassCxlReqTypSenderSubID			MassCxlReqTyp = 100
	// Mass Cancel Request TypeAccount
	MassCxlReqTypAccount					MassCxlReqTyp = 101
	// Mass Cancel Request Type Null value
	MassCxlReqTypNULL							MassCxlReqTyp = 255
)

// Returns true if the Mass Cancel Request Type is NULL
func (m MassCxlReqTyp) IsNull() bool{
	return m == MassCxlReqTypNULL
}

/*
* MassStatusOrdTyp, uInt8NULL
* Line 239
*/

type MassStatusOrdTyp UInt8NULL

const(
	// Mass Status Order Type Sender SubID
	MassStatusOrdTypSenderSubID			MassStatusOrdTyp = 100
	// Mass Status Order Type 
	MassStatusOrdTypAccount					MassStatusOrdTyp = 101
	// Mass Status Order Type Null Value
	MassStatusOrdTypNULL						MassStatusOrdTyp = 255
)

// Returns true if the Mass Status Order Type is NULL
func (m MassStatusOrdTyp) IsNull() bool{
	return m == MassStatusOrdTypNULL
}

/*
* MassStatusReqTyp, uInt8
* Line 243
*/

type MassStatusReqTyp UInt8

const(
	// Mass Status Request Type Instrument
	MassStatusReqTypInstrument				MassStatusReqTyp = 1
	// Mass Status Request Type Instrument Group
	MassStatusReqTypInstrumentGroup		MassStatusReqTyp = 3
	// Mass Status Request Type All Orders
	MassStatusReqTypAllOrders					MassStatusReqTyp = 7
	// Mass Status Request Type Market Segment
	MassStatusReqTypMarketSegment			MassStatusReqTyp = 100
)

/*
* MassStatusTIF, uInt8NULL
* Line 249
*/

type MassStatusTIF UInt8NULL

const(
	// Mass Status TIF Day
	MassStatusTIFDay				MassStatusTIF = 0
	// Mass Status TIF GTC Good Till Canceled
	MassStatusTIFGTC				MassStatusTIF = 1
	// Mass Status TIF GTD Good Till Date
	MassStatusTIFGTD				MassStatusTIF = 6
	// Mass Status TIF GFS Good for Session
	MassStatusTIFGFS				MassStatusTIF = 99
	// Mass Status TIF Day NULL value
	MassStatusTIFNULL				MassStatusTIF = 255
)

// Returns true if the MassStatusTIF value is Null
func (m MassStatusTIF) IsNull() bool{
	return m == MassStatusTIFNULL
}


/*
* OFMOverrideReq, uInt8
* Line 255
*/

type OFMOverrideReq UInt8

const(
	// OFM Override Request Disabled
	OFMOverrideReqDisabled		OFMOverrideReq = 0
	// OFM Override Request Enabled
	OFMOverrideReqEnabled			OFMOverrideReq = 1
)

/*
* OrdStatusTrd, uInt8
* Line 
*/

type OrdStatusTrd UInt8

const(
	// Order Status Trade Partially Filled
	OrdStatusTrdPartiallyFilled		OrdStatusTrd = 1
	// Order Status Trade Filled
	OrdStatusTrdFilled						OrdStatusTrd = 2
)

/*
* OrdStatusTrdCxl, CHAR
* Line 263
*/

type OrdStatusTrdCxl CHAR

const(
	// Order Status Trade Cancel Trade Correction
	OrdStatusTrdCxlTradeCorrection			OrdStatusTrdCxl = 'G'
	// Order Status Trade Cancel Trade Cancel
	OrdStatusTrdCxlTradeCancel					OrdStatusTrdCxl = 'H'
)

/*
* OrderEventType, enumNULL
* Line 267
*/

type OrderEventType EnumNULL

const(
	// Order Event Type Partially Filled
	OrderEventTypePartiallyFilled		OrderEventType = 4
	// Order Event Type Filled
	OrderEventTypeFilled						OrderEventType = 5
	// Order Event Type NULL value
	OrderEventTypeNULL							OrderEventType = 255
)

// Returns true if the Order Event Type is Null
func (o OrderEventType) IsNull() bool{
	return o == OrderEventTypeNULL
}

/*
* OrderStatus, CHAR
* Line 271
*/

type OrderStatus CHAR

const(
	// Order Status New
	OrderStatusEnumNew								OrderStatus = '0'
	// Order Status Partially Filled
	OrderStatusEnumPartiallyFilled		OrderStatus = '1'
	// Order Status Filled
	OrderStatusEnumFilled		OrderStatus = '2' 
	// Order Status Cancelled
	OrderStatusEnumCancelled					OrderStatus = '4' 
	// Order Status Replaced
	OrderStatusEnumReplaced						OrderStatus = '5' 
	// Order Status Pending Cancel since version 6
	OrderStatusEnumPendingCancel			OrderStatus = '6' 
	// Order Status Rejected
	OrderStatusEnumRejected						OrderStatus = '8' 
	// Order Status Expired
	OrderStatusEnumExpired						OrderStatus = 'C'
	// Order Status Pending Replace since version 6
	OrderStatusEnumPendingReplace			OrderStatus = 'E'
	// Order Status Undefined
	OrderStatusEnumUndefined					OrderStatus = 'U'
)

/*
* OrderType, charNULL
* Line , 283
*/

type OrderType CHAR

const(
	// Order Type Market With Protection
	OrderTypeMarketWithProtection						OrderType = '1'
	// Order Type Limit
	OrderTypeLimit													OrderType = '2'
	// Order Type Stop Limit
	OrderTypeStopLimit											OrderType = '4'
	// Order Type Market With Leftover as Limit
	OrderTypeMarketWithLeftoverAsLimit			OrderType = 'K'
	// Order Type NULL value
	OrderTypeNULL														OrderType = 0
)

// Returns true if the Order Type is NULL
func (o OrderType) IsNull() bool{
	return o == OrderTypeNULL
}


/*
* PartyDetailRole, uInt16
* Line 296
*/

type PartyDetailRole UInt16

const(
	// Party Detail Role Executing Firm
	PartyDetailRoleExecutingFirm				PartyDetailRole = 1
		// Party Detail Role Customer Account
	PartyDetailRoleCustomerAccount			PartyDetailRole = 24
		// Party Detail Role Take Up Firm
	PartyDetailRoleTakeUpFirm						PartyDetailRole = 96
		// Party Detail Role Operator
	PartyDetailRoleOperator							PartyDetailRole = 118
		// Party Detail Role Take Up Account
	PartyDetailRoleTakeUpAccount				PartyDetailRole = 1000
)

/*
* QuoteAckStatus, uInt8
* Line 303
*/

type QuoteAckStatus UInt8

const(
	// Quote Acknowledge Status Accepted
	QuoteAckStatusAccepted			QuoteAckStatus = 0
	// Quote Acknowledge Status Rejected
	QuoteAckStatusRejected			QuoteAckStatus = 5
)

/*
* QuoteCxlStatus, uInt8
* Line 307
*/

type QuoteCxlStatus UInt8

const(
	// Quote Cancel Status Cancel per Instrument
	QuoteCxlStatusCancelperInstrument						QuoteCxlStatus = 1
		// Quote Cancel Status Cancel per Instrument Group
	QuoteCxlStatusCancelperInstrumentGroup			QuoteCxlStatus = 3
		// Quote Cancel Status Cancel All Quotes
	QuoteCxlStatusCancelAllQuotes								QuoteCxlStatus = 4
		// Quote Cancel Status Rejected
	QuoteCxlStatusRejected											QuoteCxlStatus = 5
		// Quote Cancel Status Cancel per Quote Set
	QuoteCxlStatusCancelperQuoteSet							QuoteCxlStatus = 100
)

/*
* QuoteCxlTyp, uInt8
* Line 314
*/

type QuoteCxlTyp UInt8

const(
	// Quote Cancel Type Cancel per Instrument
	QuoteCxlTypCancelperInstrument					QuoteCxlTyp = 1
	// Quote Cancel Type Cancel per Instrument Group
	QuoteCxlTypCancelperInstrumentGroup			QuoteCxlTyp = 3
	// Quote Cancel Type Cancel All Quotes
	QuoteCxlTypCancelAllQuotes							QuoteCxlTyp = 4
	// Quote Cancel Type Cancel per Quote Set
	QuoteCxlTypCancelperQuoteSet						QuoteCxlTyp = 100
)

/*
* QuoteTyp, EnumNULL
* Line 320
*/

type QuoteTyp EnumNULL

const(
	// Quote Type Tradeable
	QuoteTypTradeable			QuoteTyp = 1
	// Quote Type NULL value
	QuoteTypNULL					QuoteTyp = 255
)

// Returns true if QuoteTyp is NULL
func (q QuoteTyp) IsNull() bool{
	return q == QuoteTypNULL
}

/*
* RFQSide, uInt8NULL
* Line 323
*/

type RFQSide UInt8NULL

const(
	// RFQSide Buy
	RFQSideBuy				RFQSide = 1
		// RFQSide Sell
	RFQSideSell				RFQSide = 2
		// RFQSide Cross
	RFQSideCross			RFQSide = 8
		// RFQSide NULL value
	RFQSideNULL				RFQSide = 255
)

// Returns true if RFQSide is NULL
func (r RFQSide) IsNull() bool{
	return r == RFQSideNULL
}

/*
* ReqResult, uInt8
* Line 328
*/

type ReqResult UInt8

const(
	// ReqRequest Request Result Valid Request
	ReqResultValidRequest																	ReqResult = 0
	// ReqRequest Request Result No Data Found That Matched Selection Crteria
	ReqResultNoDataFoundThatMatchedSelectionCriteria			ReqResult = 2
	// ReqRequest Request Result Not Authorized to Retrieve Data
	ReqResultNotAuthorizedtoRetrieveData									ReqResult = 3
	// ReqRequest Request Result Data Temporarily Unavailable
	ReqResultDataTemporarilyUnavailable										ReqResult = 4
)

/*
* SLEDS, uInt8NULL
* Line 334
*/

type SLEDS UInt8NULL

const(
	// SLEDS Trade Clearing at Execution Price
	SLEDSTradeClearingatExecutionPrice							SLEDS = 0
	// SLEDS Trade Clearing at Alternate Clearing Price
	SLEDSTradeClearingatAlternateClearingPrice			SLEDS = 1
	// SLEDS NULL value
	SLEDSNULL																				SLEDS = 255
)

// Return true if the SLEDS is NULL
func (s SLEDS) IsNull() bool{
	return s == SLEDSNULL
}

/*
* SMPI, charNULL
* Line 338
*/

type SMPI CHAR

const(
	// SMPI Cancel Newest
	SMPICancelNewest			SMPI = 'N'
	// SMPI Cancel Oldest
	SMPICancelOldest			SMPI = 'O'
	// SMPI NULL
	SMPINULL							SMPI = 0
)

// Returns true if the SMPI value is NULL
func (s SMPI) IsNull() bool{
	return s == SMPINULL
}

/*
* SecRspTyp, UInt8
* Line 342
*/

type SecRspTyp UInt8

const(
	// SecRspTyp Accept Security Proposal as is
	SecRspTypAcceptSecurityProposalasis																			SecRspTyp = 1
		// SecRspTyp Accept Security proposal with revisions as indicated in the message
	SecRspTypAcceptSecurityProposalWithRevisionsAsIndicatedInTheMessage			SecRspTyp = 2
		// SecRspTyp Reject Security Proposal
	SecRspTypRejectSecurityProposal																					SecRspTyp = 5
)

/*
* ShortSaleType, enumNULL
* Line 347
*/

type ShortSaleType EnumNULL

const(
	// ShortSaleType Long Sell
	ShortSaleTypeLongSell																				ShortSaleType = 0
	// ShortSaleType Short Sale With No Exemption SESH
	ShortSaleTypeShortSaleWithNoExemptionSESH										ShortSaleType = 1 
	// ShortSaleType Short Sale With Exemption SSEX
	ShortSaleTypeShortSaleWithExemptionSSEX											ShortSaleType = 2
	// ShortSaleType Undisclosed Sell Information Not Available UNDI
	ShortSaleTypeUndisclosedSellInformationNotAvailableUNDI			ShortSaleType = 3
	// ShortSaleType NULL Value
	ShortSaleTypeNULL																						ShortSaleType = 255
)

// Returns true if Short Sale Type is NULL
func (s ShortSaleType) IsNull() bool{
	return s == ShortSaleTypeNULL
}

/*
* Side, uInt8
* Line 353
*/

type Side UInt8

const(
	// Side Buy
	SideBuy			Side = 1
	// Side Sell
	SideSell		Side = 2
)

/*
* SideNULL, enumNULL 
* Line 357
*/

type SideNULL EnumNULL

const(
	// SideNULL Buy
	SideNULLBuy			SideNULL = 1
	// SideNULL Sell
	SideNULLSell		SideNULL = 2
	// SideNULL Null Value
	SideNULLNULLVal SideNULL = 255
)

// Returns true if the SideNULL value is NULL
func (s SideNULL) IsNull() bool{
	return s == SideNULLNULLVal
}

/*
* SideReq, uInt8 
* Line 361
*/

type SideReq UInt8

const(
	// SideReq Buy
	SideReqBuy						SideReq = 1 
	// SideReq Sell
	SideReqSell						SideReq = 2
	// SideReq  Undisclosed since version 6
	SideReqUndisclosed		SideReq = 7
)

/*
* SideTimeInForce, uInt8 
* Line 
*/

type SideTimeInForce UInt8

const(
	// SideTimeInForce Day
	SideTimeInForceDay			SideTimeInForce = 0
	// SideTimeInForce FAK
	SideTimeInForceFAK			SideTimeInForce = 3
)

/*
* SplitMsg, uInt8NULL 
* Line 370
*/

type SplitMsg UInt8NULL

const(
	// SplitMsg Split Message Delayed
	SplitMsgSplitMessageDelayed					SplitMsg = 0
	// SplitMsg Out of Order Message Delayed
	SplitMsgOutofOrderMessageDelayed		SplitMsg = 1
	// SplitMsg Complete Message Delayed
	SplitMsgCompleteMessageDelayed			SplitMsg = 2
	// SplitMsg NULL value
	SplitMsgNULL												SplitMsg = 255
)

// Returns true if the SplitMsg value is Null
func (s SplitMsg) IsNull() bool{
	return s == SplitMsgNULL
}

/*
* TimeInForce, uInt8NULL 
* Line 375
*/

type TimeInForce UInt8NULL

const(
	// TimeInForce Day
	TimeInForceDay		TimeInForce = 0
	// TimeInForce Good Till Cancel
	TimeInForceGoodTillCancel		TimeInForce = 1
	// TimeInForce Fill and Kill
	TimeInForceFillAndKill		TimeInForce = 3 
	// TimeInForce Fill or Kill
	TimeInForceFillOrKill		TimeInForce = 4
	// TimeInForce Good Till Date
	TimeInForceGoodTillDate		TimeInForce = 6
	// TimeInForce Good For Session since version 6
	TimeInForceGoodForSession		TimeInForce = 99
	// TimeInForce NULL value
	TimeInForceNULL		TimeInForce = 255
)

// Returns true if the TimeInForce is NULL
func (t TimeInForce) IsNull() bool{
	return t == TimeInForceNULL
}

/*
* TradeAddendum, uInt8
* Line 383
*/

type TradeAddendum UInt8

const(
	// TradeAddendum Partially Filled
	TradeAddendumPartiallyFilled		TradeAddendum = 4
	// TradeAddendum Filled
	TradeAddendumFilled							TradeAddendum = 5
	// TradeAddendum Trade Cancel
	TradeAddendumTradeCancel				TradeAddendum = 100
	// TradeAddendum Trade Correction
	TradeAddendumTradeCorrection		TradeAddendum = 101
)

/*
* Set type
* ExecInst, uInt8
* Line 389
*/

type ExecInst UInt8

const(
	// ExecInst 1 = All or None, 0 = Not All or None
	ExecInstAON					ExecInst = 1 << 0 			// (bit 0 set) 00000001 = 1
	// ExecInst 1 = Only Best, 0 = Not Only Best
	ExecInstOB					ExecInst = 1 << 1				// (bit 1 set) 00000010 = 2
	// ExecInst 1 = Not Held, 0 = Not Not Held
	ExecInstNH					ExecInst = 1 << 2				// (bit 2 set) 00000100 = 4
	// ExecInst 0 = Reserved For Future Use
	ExecInstReserved1		ExecInst = 1 << 3				// (bit 3 set) 00001000 = 8
	// ExecInst 0 = Reserved For Future Use
	ExecInstReserved2		ExecInst = 1 << 4				// (bit 4 set) 00010000 = 16
	// ExecInst 0 = Reserved For Future Use
	ExecInstReserved3		ExecInst = 1 << 5				// (bit 5 set) 00100000 = 32
	// ExecInst 0 = Reserved For Future Use
	ExecInstReserved4		ExecInst = 1 << 6				// (bit 6 set) 01000000 = 64
	// ExecInst 0 = Reserved For Future Use
	ExecInstReserved5		ExecInst = 1 << 7				// (bit 7 set) 10000000 = 128
)

// HasFlag checks if a specific bit is set
func (e ExecInst) HasFlag(flag ExecInst) bool {
  // Bit-wise AND operator  
	return e&flag != 0
}