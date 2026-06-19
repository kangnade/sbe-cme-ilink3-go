package sbe

import (
	"bytes"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

/*
message_negotiate500_test provides tests for the Negotiate500 message type
*/

// Add mutex for protecting global stdout while capturing testing output
var mutex sync.Mutex

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

// capturePrettyPrint safely captures the stdout string output by locking global access
func capturePrettyPrint(m *Negotiate500) string{
	mutex.Lock()
	defer mutex.Unlock() // ensures the lock is released when this function finishes

	// take a snapshot of the original stdout
	old := os.Stdout

	// creates a pipeline with reader and writer
	r, w, _ := os.Pipe()

	// change the target of stdout to w - the writer (stored in buffer w)
	os.Stdout = w

	// call prettyPrint() from the message object
	m.PrettyPrint()

	// close the writer after printing the message
	w.Close()
	// restore the original os stdout
	os.Stdout = old

	// creates a buffer that to read from reader r from os.Pipe()
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	
	// return the buf
	return buf.String()
}

/*------- Tests -------*/

func TestNewNegoatiate500_ConstantFields(t *testing.T){
	// allows tests to be executed concurrently
	t.Parallel()

	m := NewNegotiate500()

	if m.CustomerFlow != ClientFlowTypeValue{
		t.Errorf("CustomerFlow = %v, expect = %v", m.CustomerFlow, ClientFlowTypeValue)
	}

	if m.HMACVersion != HMACVersionValue{
		t.Errorf("HMACVersion = %v, expect = %v", m.HMACVersion, HMACVersionValue)
	}
}

func TestNewNegotiate500_ZeroValueFields(t *testing.T){
	t.Parallel()

	m := NewNegotiate500()

	if m.UUID != 0{
		t.Errorf("UUID = %v, expect = 0", m.UUID)
	}

	if m.RequestTimestamp != 0{
		t.Errorf("RequestTimestamp = %v, expect = 0", m.RequestTimestamp)
	}

  if m.Credentials.Length != 0{
		t.Errorf("Credentials.Length = %v, expect = 0", m.Credentials.Length)
	}
}

func TestGetMessageTypeID(t *testing.T){
	t.Parallel()

	m := NewNegotiate500()

	if typeId := m.GetMessageTypeID(); typeId != 500{
		t.Errorf("MessageTypeID = %v, expect = 500", typeId)
	}
}

func TestBlockLength(t *testing.T){
	t.Parallel()

	m := NewNegotiate500()

	if blockLen := m.BlockLength(); blockLen != 76{
		t.Errorf("BlockLength = %v, expect = 76", blockLen)
	}
}

func TestWireSize(t *testing.T){
	t.Parallel()

	msgNoCred := makeNegotiate500()
	msgWithCred := makeNegotiate500WithCredentials()

	tests := []struct{
		name 			string
		msg  			*Negotiate500
		expected 	int
	}{
		{"No Credentials", msgNoCred, 8 + 76 + 2 + 0},
		{"With Credentials", msgWithCred, 8 + 76 + 2 + len(msgWithCred.Credentials.VarData)},
	}

	for _, st := range tests{
		t.Run(st.name, func(t *testing.T){
			if size := st.msg.WireSize(); size != st.expected{
				t.Errorf("WireSize() = %v, expected = %v", size, st.expected)
			}
		})
	}
}

func TestEncode_Offsets(t *testing.T){
	t.Parallel()

	t.Run("Fixed Length Fields", func(t *testing.T){
		m := makeNegotiate500()
		bufSize := int(m.BlockLength()) + 2 // Negotiate500 Block Length 76 + 2 Credentials <data> length prefix
		c := Encoder(bufSize)
		m.Encode(c)

		expectedOffset := 32 + 20 + 8 + 8 + 3 + 5 + 2
		if received := c.GetOffset(); received != expectedOffset{
			t.Errorf("Encode offset = %d, Expected = %d", received, expectedOffset)
		}
	})

	t.Run("Variable Length Fields", func(t *testing.T){
		m := makeNegotiate500WithCredentials()
		bufSize := int(m.BlockLength()) + 2 + int(m.Credentials.Length)
		c := Encoder(bufSize)
		m.Encode(c)

		expectedOffset := bufSize
		if received := c.GetOffset(); received != expectedOffset{
			t.Errorf("Encode offset = %d, Expected = %d", received, expectedOffset)
		}
	})
}

func TestPrettyPrint_Output_Validation(t *testing.T){
	m := makeNegotiate500WithCredentials()
	output := capturePrettyPrint(m)

	expectedFields := []string{
		"=== Negotiate500 (id = 500) ===",
		"CustomerFlow:",
		"HMACVersion:",
		"HMACSignature:",
		"AccessKeyID:",
		"UUID:",
		"RequestTimestamp:",
		"Session:",
		"Firm:",
		"Credentials.Length:",
		"Credentials.Data:",
	}

	for _, field := range expectedFields{
		if !strings.Contains(output, field){
			t.Errorf("PrettyPrint() output missing field: %q", field)
		}
	}

	expectedFlow := clean(ClientFlowTypeValue[:])
	expectedHMAC := clean(HMACVersionValue[:])

	if !strings.Contains(output, expectedFlow){
		t.Errorf("PrettyPrint() failed to accurately resolve raw flow constant text %q", expectedFlow)
	}
	if !strings.Contains(output, expectedHMAC) {
		t.Errorf("PrettyPrint() failed to accurately resolve raw HMAC scheme constant text %q", expectedHMAC)
	}
}

func TestPrettyPrint_OmitEmptyCredentials(t *testing.T) {
	m := makeNegotiate500() // Length = 0
	output := capturePrettyPrint(m)

	if strings.Contains(output, "Credentials.Data:") {
		t.Error("PrettyPrint() structural leak: Printed empty Credentials.Data track even when byte array length is zero")
	}
}