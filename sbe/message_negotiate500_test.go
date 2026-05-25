package sbe

/*
message_negotiate500_test provides tests for the Negotiate500 message type
*/

/*------- Helpers -------*/

// makeNegotiate500 returns a Negotitate 500 message with no credentials filled
func makeNegotiate500() *Negotiate500{
	m := NewNegotiate500()
	copy(m.HMACSignature[:], []byte("test-cme-sbe-hmca-signature-long"))
	copy(m.AccessKeyID[:], []byte("TESTACCESSKEYSAXO01"))
	m.UUID = UInt64(123456789)
	m.RequestTimestamp = UInt64(1779630405)
	copy(m.Session[:], []byte("S01"))
	copy(m.Firm[:], []byte("TESTSAXO01"))
	m.Credentials = DATA{Length: 0}
	return m
}

// makeNegotiate500WithCredentials returns a fully filled Negotiate 500 message
// with Credentials filled
func makeNegotiate500WithCredentials() *Negotiate500{
	m := makeNegotiate500()
	credVarData := []byte("test message")
	m.Credentials = DATA{
		Length: UInt16(len(credVarData)),
		VarData: credVarData,
	}
	return m
}