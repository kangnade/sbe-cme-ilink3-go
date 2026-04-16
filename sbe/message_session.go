package sbe

import (
	"encoding/binary"
)

/*
* message_session.go contains the message types in ilinkbinary.xml
* that handle the FIX Protocol session layer.
* More specifically, here includes:
* Message id = 500, 501, 502, 503, 504, 505, 506, 507
* 508, 509, 510, 513
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

// Need Encode and Decode methods for Negotiate500
func (m *Negotiate500) Encode() []byte{
	// the total length of the buffer should be
	// blockLength 0-75 bytes
	// Credentials (DATA is UInt16) 2 bytes
	// len(Credentials.VarData)
	bufferSize := 76 + 2 + len(m.Credentials.VarData)

	buffer := make([]byte, bufferSize)

	// offset 0, 0-31, HMACSignature String32Req
	copy(buffer[0:32], m.HMACSignature[:])

	// offset 32, 32-52 AccessKeyID
	copy(buffer[32:52], m.AccessKeyID[:])

	// offset 52, 52-60 UUID
	binary.LittleEndian.PutUint64(buffer[52:60], uint64(m.UUID))

	// offet 60, 60-68 RequestTimestamp
	binary.LittleEndian.PutUint64(buffer[60:68], uint64(m.RequestTimestamp))

	// offset 68, 68-71, Session
	copy(buffer[68:71], m.Session[:])

	// offset 71, 71-76, Firm
	copy(buffer[71:76], m.Firm[:])

	// offset 76, Credential length uint16 == 2 bytes
	binary.LittleEndian.PutUint16(buffer[76:78], uint16(m.Credentials.Length))

	// If Credentials.length > 0, append data to the message
	if m.Credentials.Length > 0{
		copy(buffer[78:], m.Credentials.VarData)
	}

	return buffer
}

func(m *Negotiate500) Decode(buffer []byte){
	// offset 0, 0-31, HMACSignature String32Req
	copy(m.HMACSignature[:], buffer[0:32])

	// offset 32, 32-52 AccessKeyID
	copy(m.AccessKeyID[:], buffer[32:52])

	// offset 52, 52-60 UUID
	m.UUID = UInt64(binary.LittleEndian.Uint64(buffer[52:60]))

	// offet 60, 60-68 RequestTimestamp
	m.RequestTimestamp = UInt64(binary.LittleEndian.Uint64(buffer[60:68]))

	// offset 68, 68-71, Session
	copy(m.Session[:], buffer[68:71])

	// offset 71, 71-76, Firm
	copy(m.Firm[:], buffer[71:76])

	// offset 76, Credential length uint16 == 2 bytes
	m.Credentials.Length = UInt16(binary.LittleEndian.Uint16(buffer[76:78]))

	// If Credentials.length > 0, append data to the message
	if m.Credentials.Length > 0{
		m.Credentials.VarData = make([]byte, m.Credentials.Length)
		copy(m.Credentials.VarData, buffer[78:])
	}
} 