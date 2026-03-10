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
 */

// ilinkbinary.xml line 67 to 74 basic types, check Constants (basic types) for NULL values
type UInt16 uint16

type UInt32 uint32

type UInt64 uint64

type UInt8 uint8

// ilinkbinary.xml line 22 to 23 basic types, check Constants (basic types) for NULL values
type Int32 int32

// line 25 description: LocalMktDate
// const is LocalMktDateNULL
type LocalMktDate uint16

// ilinkbinary.xml line 66 type
type EnumNULL uint8

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
// --- END Execution Types ---

// line 21, description: HMACVersion
// const is HMACVersionValue
type HMACVersion [13]byte

// line 24, description: Leg security ID source in UDS creation, check SecurityIDSource <22> field
// const is LegSecIDSourceExchangeSymbol
type LegSecIDSource byte


/*
 * --- Constants (basic types) ---
 */

const(
     UInt16NULL UInt16 = 65535
     UInt32NULL UInt32 = 4294967295
     UInt64NULL UInt64 = 18446744073709551615
     UInt8NULL UInt8 = 255
     EnumNULLValue UInt8 = 255
     Int32NULL Int32 = 2147483647
     LocalMktDateNULL LocalMktDate = 65535
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
)

/*
 * --- Leg Security ID Source TYpes / FIX 4.4 : FIX 4.4 : LegSecurityIDSource <603> field ---
 */
 const(
     LegSecIDSourceExchangeSymbol LegSecIDSource = '8'
 )

/*
 * --- Package Level Variables ---
 */
var ClientFlowTypeValue = ClientFlowType{'I', 'D', 'E', 'M', 'P', 'O', 'T', 'E', 'N', 'T'}
var ExchangeFlowType = ExchFlowTyp{'R', 'E', 'C', 'O', 'V', 'E', 'R', 'A', 'B', 'L', 'E'}
var HMACVersionValue = HMACVersion{'C', 'M', 'E', '-', '1', '-', 'S', 'H', 'A', '-', '2', '5', '6'}