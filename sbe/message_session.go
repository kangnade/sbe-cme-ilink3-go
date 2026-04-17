package sbe

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

func(m *Negotiate500) Decode(c *Coder){
	c.Decode(&m.HMACSignature) 								// Offset = 0, 32 bytes
	c.Decode(&m.AccessKeyID)									// Offset = 32, 20 bytes
	c.Decode(&m.UUID)													// Offset = 52, 8 bytes
	c.Decode(&m.RequestTimestamp)							// Offset = 60, 8 bytes
	c.Decode(&m.Session)											// Offset = 68, 3 bytes
	c.Decode(&m.Firm)													// Offset = 71, 5 bytes
	// Field blockLength=76 bytes complete
	// Data field:
	c.Decode(&m.Credentials.Length)
	if m.Credentials.Length > 0{
		m.Credentials.VarData = c.DecodeRawVarLen(int(m.Credentials.Length))
	}
} 