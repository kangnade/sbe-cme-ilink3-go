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
	encoded := Encoder(76 + 2 + int(original.Credentials.Length))
	original.Encode(encoded)
	encodedBytes := encoded.GetBytes()

	// 3. Decode the previously encoded message
	decoder := Decoder(encodedBytes)
	decoded := &Negotiate500{}
	decoded.Decode(decoder)

	// 4. Validate the result
	if decoded.UUID != original.UUID {
		t.Errorf("Unmatched UUID. Expected: %d, Got: %d", original.UUID, decoded.UUID)
	}
	if decoded.RequestTimestamp != original.RequestTimestamp {
		t.Errorf("Unmatched RequestTimestamp. Expected: %d, Got: %d", original.RequestTimestamp, decoded.RequestTimestamp)
	}
	if !bytes.Equal(decoded.HMACSignature[:], original.HMACSignature[:]) {
		t.Errorf("Unmatched HMACSignature. Expected: %s, Got: %s", original.HMACSignature[:], decoded.HMACSignature[:])
	}
	if !bytes.Equal(decoded.AccessKeyID[:], original.AccessKeyID[:]) {
		t.Errorf("Unmatched AccessKeyID. Expected: %s, Got: %s", original.AccessKeyID[:], decoded.AccessKeyID[:])
	}
	if !bytes.Equal(decoded.Session[:], original.Session[:]) {
		t.Errorf("Unmatched Session. Expected: %s, Got: %s", original.Session[:], decoded.Session[:])
	}
	if !bytes.Equal(decoded.Firm[:], original.Firm[:]) {
		t.Errorf("Unmatched Firm. Expected: %s, Got: %s", original.Firm[:], decoded.Firm[:])
	}
	if decoded.Credentials.Length != original.Credentials.Length {
		t.Errorf("Unmatched Credentials.Length. Expected: %d, Got: %d", original.Credentials.Length, decoded.Credentials.Length)
	}
	if !bytes.Equal(decoded.Credentials.VarData, original.Credentials.VarData) {
		t.Errorf("Unmatched Credentials.VarData. Expected: %s, Got: %s", original.Credentials.VarData, decoded.Credentials.VarData)
	}
}