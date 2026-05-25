package sbe

import (
	"fmt"
)

/*
* Defines the MessageType Negotiate500
 */

// For copy and paste
// id = "", description = "", blockLength = "", semantcType=""

// ilinkbinary.xml file line 400
// id = "500", description = "Negotiate", blockLength = "76", semantcType="Negotiate"
type Negotiate500 struct{
	// id = "39000" type = "ClientFlowType" description = "Constant value representing type of flow from customer to CME" semanticType = "String"
	CustomerFlow 				ClientFlowType
	// id = "39003" type = "HMACVersion" description = "Constant value representing CME HMAC version" semanticType = "String"
	HMACVersion					HMACVersion
	// id = "39005" type = "String32Req" description = "Contains the HMAC signature" offset="0" semanticType = "String"
	HMACSignature				String32Req
	// id = "39004" type = "String20Req" description = "Contains the AccessKeyID assigned to this session on this port" offset = "32" semanticType = "String"
	AccessKeyID					String20Req
	// id = "39001" type = "uInt64" description = "Session Identifier defined as type long (uInt64); recommend to use timestamp as 
	// number of microseconds since epoc (Jan 1, 1970)" offset = "52" semanticType = "int"
	UUID								UInt64	
	// id = "39002" type = "uInt64" description = "Time of request; recommend to use timestamp as number of nanoseconds since 
	// epoch (Jan 1, 1970)" offset = "60" semanticType = "int"
	RequestTimestamp		UInt64
	// id = "39006" type = "String3Req" description = "Session ID" offset = "68" semanticType = "String"
	Session							String3Req
	// id = "39007" type = "String5Req" description = "Firm ID" offset = "71" semanticType = "String"
	Firm								String5Req
	// id = "39008" type = "DATA" description = "Not used and will be set to 0" semanticType = "data"
	Credentials					DATA
}

// NewNegotitate500 returns the pointer to a new Negotiate500 message with constant values filled
func NewNegotiate500() *Negotiate500{
	return &Negotiate500{
		// Initializes with the CustomerFlow and HMACVersion constant fields
		CustomerFlow: ClientFlowTypeValue,
		HMACVersion: HMACVersionValue,
	}
}

// Encodes the Negotitate500 Message into SBE bytes to be sent to CME
func (m *Negotiate500) Encode(c *Coder){
	c.Encode(&m.HMACSignature) 								// Offset = 0, 32 bytes
	c.Encode(&m.AccessKeyID)									// Offset = 32, 20 bytes
	c.Encode(&m.UUID)													// Offset = 52, 8 bytes
	c.Encode(&m.RequestTimestamp)							// Offset = 60, 8 bytes
	c.Encode(&m.Session)											// Offset = 68, 3 bytes
	c.Encode(&m.Firm)													// Offset = 71, 5 bytes
	// Field blockLength=76 bytes complete
	// Data field:
	c.Encode(&m.Credentials.Length)
	if m.Credentials.Length > 0{
		c.EncodeRawVarLen(m.Credentials.VarData)
	}
}

// GetMessageTypeID return the message TemplateId as UInt16, explicitly cast it
// to uint16 to be used in parser.go
func (m *Negotiate500) GetMessageTypeID() UInt16{
	return 500
}

// WireSize returns the total bytes needed for the Negotitate500 message
// MessageHeader (8 bytes) + BlockLength (76 bytes) + Credential.Length (2 bytes) + Credential.VarData
func (m *Negotiate500) WireSize() int{
	// 8 bytes MessageHeader
	// 76 bytes BlockLength for Negotiate500
	// 2 bytes for Credentials.Length
	// int(Credentials.Length) bytes for Credentials.VarData
	return 8 + int(m.BlockLength()) + 2 + int(m.Credentials.Length)
}

// BlockLength returns the block length of Negotiate 500 as 76
func (m *Negotiate500) BlockLength() UInt16{
	return 76
}

// PrettyPrint prints the Negotiate500 Message into readable output for debugging purpose
func (m *Negotiate500) PrettyPrint(){
	fmt.Printf("=== Negotiate500 (id = 500) ===\n")
	// For each field, calls the clean() method from utils.go to remove null byte
	// using strings.TrimRight(string(m.field), "\x00") can be slow
	fmt.Printf("CustomerFlow:					%s (Constant)\n", clean(m.CustomerFlow[:]))
	fmt.Printf("HMACVersion:					%s (Constant)\n", clean(m.HMACVersion[:]))
	fmt.Printf("HMACSignature:				%s\n", clean(m.HMACSignature[:]))
	fmt.Printf("AccessKeyID:					%s\n", clean(m.AccessKeyID[:]))
	fmt.Printf("UUID:									%d\n", m.UUID)
	fmt.Printf("RequestTimestamp:			%d\n", m.RequestTimestamp)
	fmt.Printf("Session:							%s\n", clean(m.Session[:]))
	fmt.Printf("Firm:									%s\n", clean(m.Firm[:]))
	fmt.Printf("Credentials.Length:		%d\n", m.Credentials.Length)
	if m.Credentials.Length > 0{
		fmt.Printf("Credentials.Data:		%x\n", m.Credentials.VarData)
	}
	fmt.Printf("===============================\n")
}