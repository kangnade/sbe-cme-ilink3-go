package sbe

/*
message_negotiation_response_5001_test provides tests for the NegotiationResponse501 message type
*/

/*------- Helpers -------*/

// makeNegotiationResponse501 returns a Negotitation Response 501 message with no credentials filled
func makeNegotiationResponse501() *NegotiationResponse501{
	m := NewNegotiationResponse501()

	m.UUID = UInt64(123456789)
	m.RequestTimestamp = UInt64(1779630405)
	m.SecretKeySecureIDExpiration = UInt16NULL(UInt16NULLValue)
	m.FaultToleranceIndicator = FTIPrimary
	// SplitMsg types are defined in enums.go
	m.SplitMsg = SplitMsgCompleteMessageDelayed
	m.PreviousSeqNo = UInt32(123456)
	m.PreviousUUID = UInt64(987654321)
	m.EnvironmentIndicator = UInt8NULL(UInt8NULLValue)
	m.Credentials = DATA{Length: 0}

	return m
}

// makeNegotiationResponse501WithCredentials returns a Negotitation Response 501 message with credentials filled
func makeNegotiationResponse501WithCredentials() *NegotiationResponse501{
	m := makeNegotiationResponse501()
	credentialVarData := []byte("session_token_test")
	m.Credentials = DATA{
		Length: UInt16(len(credentialVarData)),
		VarData: credentialVarData,
	}
	return m
}