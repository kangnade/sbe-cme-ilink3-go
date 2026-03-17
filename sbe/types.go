package sbe

/*
Golang primitive Types Reference:
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
 * --- Types Definition (from CME Group ilinkbinary.xml) ---
 * "line" refers to line # in ilinkbinary.xml file
 */

// Additional Types For composites.go (from composites part)
// line 79
// const is Int8NULL
type Int8 int8

type Int8NULL int8

// line 83
type Int64 int64

type Int64NULL int64

// ilinkbinary.xml line 67 to 74 basic types, check Constants (basic types) for NULL values
// line 67
type UInt16 uint16
// line 68
type UInt16NULL uint16
// line 69
type UInt32 uint32
// line 70
type UInt32NULL uint32
// line 71
type UInt64 uint64
// line 72
type UInt64NULL uint64
// line 73
type UInt8 uint8
// line 74
type UInt8NULL uint8

// ilinkbinary.xml line 22 to 23 basic types, check Constants (basic types) for NULL values
// line 22
type Int32 int32

// line 23
type Int32NULL int32

// line 25 description: LocalMktDate
// const is LocalMktDateNULL
type LocalMktDate uint16

// ilinkbinary.xml line 66 type
type EnumNULL uint8

// line 28, description: NoPtyUpd
// const is NoPtyUpdValue
type NoPtyUpd uint8

// ilinkbinary.xml line 4 to line 20
// line 4, description: char
type CHAR byte 

// line 5, description: ClientFlowType
// const is ClientFlowTypeValue
type ClientFlowType [10]byte

// line 6, description: Cross order type supports only limit order
// const is CrossOrderLimitOrder
type CrossOrderType byte

// line 7, description: Cross Prioritization
type CrossPrioritization byte

// line 8, description: CrossType
type CrossType byte

// line 64, description UDI
// const is UDIValue
type UDI byte

/*
 * The below 2 Cancel Types are
 * FIX 4.4 : CxlRejResponseTo <434> field
 */

// line 9, description: Cancel Reject
// const is OrderCancelRequestRejected
type CxlRejRsp byte

// line 10, description: Cancel Replace Reject
// const is OrderCancelReplaceRequestRejected
type CxrRejRsp byte

// --- END CxlRejResponseTo Types ---


// line 11, description: ExchangeFlowType
// const is ExchangeFlowType
type ExchFlowTyp [11]byte


/*
 * Below Execution Types correspond to FIX 4.4 : ExecType <150> field
 */

// line 12, description: ExecTypNew
// const is ExecutionNew
type ExecTypNew byte

// line 13, description: Exec Type Reject
// const is ExecutionRejected
type ExecTypRej byte

// line 14, description: ExecTypStatus
// const is ExecutionOrderStatus
type ExecTypStatus byte

// line 15, description: Exec Type Cxl
// const is ExecutionCanceled
type ExecTypeCxl byte

// line 16, description: Exec Type Exp (Exp for Expired)
// const is ExecutionExpired
type ExecTypeExp byte

// line 17, description: ExecType Modify
// const is ExecutionReplaced
type ExecTypeModify byte

// line 18, description: Pending Cancel Type, result of FIX 4.4: Order Cancel Request <F> message
// const is ExecutionPendingCanceled
type ExecTypePendCxl byte

// line 19, description: Pending Replace, result of FIX 4.4: Order Cancel/Replace Request <G>)
// const is ExecutionPendingReplace
type ExecTypePendModify byte

// line 20, description: Execution type for trade, fill or partial fill
// const is ExecutionTrade
type ExecTypeTrade byte

// line 27, description: ModifyStatus, old OrdStatus(39) - 5, ExecType(150) - 5 Replace
// https://www.onixs.biz/fix-dictionary/5.0/app_6_f.html
// const is ExecutionReplaceStatus
type ModifyStatus byte

// --- END Execution Types ---


// line 21, description: HMACVersion
// const is HMACVersionValue
type HMACVersion [13]byte

// line 24, description: Leg security ID source in UDS creation, check SecurityIDSource <22> field
// const is LegSecIDSourceExchangeSymbol
type LegSecIDSource byte

// line 36, description: PartyIDSource
// const is GenAcceptedMktParID
// C = Generally accepted market participant identifier (e.g. NASD mnemonic)
type PartyIDSource byte

// line 41, description SecurityIDSource
// const is SecurityIDSourceValue
type SecurityIDSource byte

// line 42, description: ecurity request type for UDS creation
// const is SecurityReqTypeUSDCreation
type SecurityReqType byte

/*
 * Mass Action FIX 5.0 SP2 : MassActionType <1373> field
 */

 // line 26, description: Mass action type to represent mass cancel
type MassAction byte

// --- END Mass Action Types ---


/*
 * --- Order Status Types / Partially Matched FIX 4.4 : OrdStatus <39> field ---
 * https://www.onixs.biz/fix-dictionary/4.4/tagnum_39.html
 */

// line 29, description: Ord Status Cxl
// const is OrderStatusCanceled
type OrdStatusCxl byte

// line 30, description: OrdStatusCxlRej
// const is OrderStatusCancelRejected
type OrdStatusCxlRej byte

// line 31, description: OrdStatusCxrRej
// const is OrderStatusCancelReplaceRejected
type OrdStatusCxrRej byte

// line 32, description: Order Status Exp
// const is OrderStatusExpired
type OrdStatusExp byte

// line 33, description: Order status of New
// const is OrderStatusNew
type OrdStatusNew byte

// line 34, description: Pending Cancel Status - e.g. result of Order Cancel Request <F>
// const is OrderStatusPendingCancel
type OrdStatusPendCxl byte

// line 35, description: Order Status Reject
// const is OrderStatusRejected
type OrdStatusRej byte

// line 37, description: Pending Replace Status
// const is: OrderStatusPendingReplace
type PendModStatus byte

// --- END Order Status Types ---

/*
 * --- Quote Cancel Types ---
 */

// line 38, description: Quote Cancel by a list of Security Groups
// const is QuoteCancelBySecurityGroup
type QuoteCxTypeByGroup uint8

// line 39, description: Quote Cancel Type by Instrument
// const is QuoteCancelByInstrument
type QuoteCxTypeByInstr uint8

// line 40, description: Quote Cancel Type by Set
// const is QuoteCancelBySet
type QuoteCxTypeBySet uint8

// --- END Quote Cancel Types ---

/*
 * --- String Types ---
 */

 // line 43, description: String with length of 10 required
type String10Req [10]byte

// line 44, description: Optional string with length of 17
type String17 [17]byte

// line 45, description: StringLength2
type String2 [2]byte

// line 46, description: String Length 20
type String20 [20]byte

// line 47, description: String With Length 20 (required)
type String20Req [20]byte

// line 48, description: String with length of 256
type String256 [256]byte

// line 49, description: String Length of 3 characters
type String3 [3]byte

// line 50, description: String Length 30
type String30 [30]byte

// line 51, description: String with length of 30 required
type String30Req [30]byte

// line 52, description: String with length of 32 required
type String32Req [32]byte

// line 53, description: String with length 3 required
type String3Req [3]byte

// line 54, description: String length 40 char
type String40 [40]byte

// line 55, description: String with length 48
type String48 [48]byte

// line 56, description: Optional string with length of 5
type String5 [5]byte

// lin 57, description: String with length of 5 required
type String5Req [5]byte

// line 58, description: String with length of 60
type String60 [60]byte

// line 59, description: String with length of 75
type String75 [75]byte

// line 60, description: String field length 8
type String8 [8]byte

// line 61, description: String with length of 8 required
type String8Req [8]byte

// line 62, description: String with length of 35
type StringLength35 [35]byte

// line 63, description: StringLength6
type StringLength6 [6]byte

 // --- END String Types ---

/*
 * --- Constants (basic types) ---
 */

const(
     UInt16NULLValue UInt16NULL = 65535
     UInt32NULLValue UInt32NULL = 4294967295
     UInt64NULLValue UInt64NULL = 18446744073709551615
     UInt8NULLValue UInt8NULL = 255
     CharNULL CHAR = 0 // line 65
     EnumNULLValue EnumNULL = 255
     UDIValue UDI = 'Y'
     Int8NULLValue Int8NULL = 127
     Int32NULLValue Int32NULL = 2147483647
     Int64NULLValue Int64NULL = 9223372036854775807
     LocalMktDateNULL LocalMktDate = 65535
     NoPtyUpdValue NoPtyUpd = 1
     GenAcceptedMktParID PartyIDSource = 'C'
     SecurityIDSourceValue SecurityIDSource = '8'
     SecurityReqTypeUSDCreation SecurityReqType = '1'
)

/*
 * --- Cross Order Type Value ---
 */
const(
     CrossOrderLimitOrder CrossOrderType = '2'
     CrossPrioritizationValue CrossPrioritization = '0'
     CrossTypeValue CrossType = '3'
)

/*
 * --- Cancel Type Values / FIX 4.4 : CxlRejResponseTo <434> field ---
 */
const(
     OrderCancelRequestRejected CxlRejRsp = '1'
     OrderCancelReplaceRequestRejected CxrRejRsp = '2'
)

/*
 * --- Execution Type Values / FIX 4.4 : ExecType <150> field ---
 */
const(
     ExecutionNew ExecTypNew = '0'
     ExecutionRejected ExecTypRej = '8'
     ExecutionOrderStatus ExecTypStatus = 'I'
     ExecutionCanceled ExecTypeCxl = '4'
     ExecutionExpired ExecTypeExp = 'C'
     ExecutionReplaced ExecTypeModify = '5'
     ExecutionPendingCanceled ExecTypePendCxl = '6'
     ExecutionPendingReplace ExecTypePendModify = 'E'
     ExecutionTrade ExecTypeTrade = 'F'
     ExecutionReplaceStatus ModifyStatus = '5'
)

/*
 * --- Leg Security ID Source Type Values / FIX 4.4 : FIX 4.4 : LegSecurityIDSource <603> field ---
 */
 const(
     LegSecIDSourceExchangeSymbol LegSecIDSource = '8'
 )

/*
 * --- Mass Action Type Values / FIX 5.0 SP2 : MassActionType <1373> field ---
 */
const(
     MassActionCancelOrders MassAction = '3'
)

/*
 * --- Order Status Type Values / Partially Matched FIX 4.4 : OrdStatus <39> field ---
 */
const(
     OrderStatusCanceled OrdStatusCxl = '4'
     OrderStatusCancelRejected OrdStatusCxlRej = 'U'
     OrderStatusCancelReplaceRejected OrdStatusCxrRej = 'U'
     OrderStatusExpired OrdStatusExp = 'C'
     OrderStatusNew OrdStatusNew = '0'
     OrderStatusPendingCancel OrdStatusPendCxl = '6'
     OrderStatusRejected OrdStatusRej = '8'
     OrderStatusPendingReplace PendModStatus = 'E'
)

/*
 * --- Quote Cancel Type Values ---
 */

const(
     QuoteCancelBySecurityGroup QuoteCxTypeByGroup = 3
     QuoteCancelByInstrument QuoteCxTypeByInstr = 1
     QuoteCancelBySet QuoteCxTypeBySet = 100
)

/*
 * --- Package Level Variables ---
 */
var ClientFlowTypeValue = ClientFlowType{'I', 'D', 'E', 'M', 'P', 'O', 'T', 'E', 'N', 'T'}
var ExchangeFlowType = ExchFlowTyp{'R', 'E', 'C', 'O', 'V', 'E', 'R', 'A', 'B', 'L', 'E'}
var HMACVersionValue = HMACVersion{'C', 'M', 'E', '-', '1', '-', 'S', 'H', 'A', '-', '2', '5', '6'}

/*
 * --- IsNULL() methods for NULL types ---
 */

func(v Int8NULL) IsNULL() bool{ return v == Int8NULLValue }

func(v Int32NULL) IsNULL() bool { return v == Int32NULLValue }

func(v Int64NULL) IsNULL() bool { return v == Int64NULLValue }

func(v UInt8NULL) IsNULL() bool { return v == UInt8NULLValue }

func(v UInt16NULL) IsNULL() bool { return v == UInt16NULLValue }

func(v UInt32NULL) IsNULL() bool {return v == UInt32NULLValue }

func(v UInt64NULL) IsNULL() bool { return v == UInt64NULLValue }

func(v EnumNULL) IsNULL() bool { return v == EnumNULLValue }

func(v LocalMktDate) IsNULL() bool { return v == LocalMktDateNULL }

func(v CHAR) IsNULL() bool { return v == CharNULL }