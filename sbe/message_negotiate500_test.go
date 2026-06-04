package sbe

import (
	"bytes"
	"io"
	"os"
	"sync"
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