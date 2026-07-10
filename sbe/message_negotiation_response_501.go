package sbe

// Negotiation Response 501 is the negotiation response message from CME to customer.
// The ilink 3 binary specification has: blockLength = 33, id = 501
// Negotiation Response 501 is an inbound message type that implements the ILinkInbound interface
type NegotiationResponse501 struct{
	// ServerFlow is the constant value representing type of flow from CME to customer
	ServerFlow					ExchFlowTyp

	// Fixed payload fields in the wire:
	// Matches Negotiate.UUID
	UUID 															UInt64				// offset = 0, 8 bytes
	// Matches Negotiate.RequestTimestamp
	RequestTimestamp									UInt64				// offset = 8, 8 bytes
	// This indicates in how many days the HMAC secret key will expire
	SecretKeySecureIDExpiration				UInt16NULL		// offset = 16, 2 bytes
	// Indicates whether the connection is primary or backup
	FaultToleranceIndicator						FTI						// offset = 18, 1 bytes
	// Indicates whether a message was delayed as a result of being split among multiple 
	// packets (0) or if a message was delayed as a result of TCP re-transmission (1) or 
	// if a complete message was delayed due to a previously submitted split or out of order 
	// message (2). If absent then the message was not delayed and was neither split nor received out of order
	SplitMsg													SplitMsg			// offset = 19, 1 bytes
	// Refers to the SeqNum sent in the previous message before this one from CME
	PreviousSeqNo											UInt32				// offset = 20, 4 bytes
	// Refers to the UUID sent in the previous message before this one from CME
	PreviousUUID											UInt64				// offset = 24, 8 bytes
	// Provides customers with the information of the type of environment they are connecting to
	EnvironmentIndicator							UInt8NULL			// offset = 32, 1 byte
	
	// Variable length data field:
	// Not used and will be set to 0
	Credentials												DATA
}

// NewNegotiationResponse501 creates a new instance of NegotiationResponse501 message
// with pre-populated constants fields
func NewNegotiationResponse501() *NegotiationResponse501{
	return &NegotiationResponse501{
		ServerFlow: ExchangeFlowType, // The constant value for ExchFlowTyp is ExchangeFlowType
	}
}

// GetMessageTypeID returns the message type id for the message type, for NegotiationResponse501
// it is: 501
func (m *NegotiationResponse501) GetMessageTypeID() UInt16{
	return 501
}

// BlockLength returns the block length of the message type, for Negotiation Response 501, this is: 33
func (m *NegotiationResponse501) BlockLength() UInt16{
	return 33
}

// Decode method decodes the incoming byte stream into the fields of the NegotiationResponse501 struct
func (m *NegotiationResponse501) Decode(c *Coder){
	// Populate the constant value ServerFlow with ExchangeFlowType
	m.ServerFlow = ExchangeFlowType

	// Use the Coder to decode the incoming byte stream into the fields of NegotiationResponse501 struct
	c.Decode(&m.UUID)
	c.Decode(&m.RequestTimestamp)
	c.Decode(&m.SecretKeySecureIDExpiration)
	c.Decode(&m.FaultToleranceIndicator)
	c.Decode(&m.SplitMsg)
	c.Decode(&m.PreviousSeqNo)
	c.Decode(&m.PreviousUUID)
	c.Decode(&m.EnvironmentIndicator)
	c.Decode(&m.Credentials.Length)

	// Only if the credentials is not 0, we'll decode the variable length data
	if m.Credentials.Length > 0{
		m.Credentials.VarData = c.DecodeRawVarLen(int(m.Credentials.Length))
	}
}