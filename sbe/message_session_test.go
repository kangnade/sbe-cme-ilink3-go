package sbe

import (
	"bytes"
	"testing"
	"time"
)

func TestNegotiate500_RoundTrip(t *testing.T){
	// 1. Create test Negotiate500 message
	original := &Negotiate500{
		UUID: 123456789,
		RequestTimestamp: UInt64(time.Now().UnixNano()),
		Credentials: DATA{
			Length: 4,
			VarData: []byte{'t', 'e', 's', 't'},
		},
	}
	// Fill in the fixed-size byte array
	copy(original.HMACSignature[:], "hmac-signature") // must be within 32 bytes
	copy(original.AccessKeyID[:], "ABCD123") // must be within 20 bytes
	copy(original.Session[:], "122") // must be within 3 bytes
	copy(original.Firm[:], "SAXOB") // must be within 5 bytes

	// 2. Encode the Negotiate500 message to bytes
	encoded := original.Encode()

	// 3. Decode the previously encoded message
	decoded := &Negotiate500{}
	decoded.Decode(encoded)

	// 4. Validate the result
	if decoded.UUID != original.UUID{
		t.Errorf("Unmatched UUID. Expected: %d, Result: %d.", original.UUID, decoded.UUID)
	}
	if !bytes.Equal(decoded.HMACSignature[:], original.HMACSignature[:]){
		t.Errorf("Unmatched HMACSignature. Expected: %s, Result: %s", original.HMACSignature[:], decoded.HMACSignature[:])
	}
	if !bytes.Equal(decoded.Credentials.VarData, original.Credentials.VarData){
		t.Errorf("Unmatched Credential VarData. Expected: %s, Result: %s", original.Credentials.VarData, decoded.Credentials.VarData)
	}
}