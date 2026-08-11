package sbe

import (
	"bytes"
	"io"
	"os"
	"sync"
)

/*--- Define the global mutex for testing PrettyPrint() ---*/
var printMutex		sync.Mutex

/*--- Define PrettyPrinter interface for message types that implement the PrettyPrint() method ---*/
type PrettyPrinter interface{
	PrettyPrint()
}

// capturePrettyPrint safely captures the stdout string output by locking global access
func capturePrettyPrint(m *Negotiate500) string{
	// printMutex is defined globally in the message_test_helpers.go file
	printMutex.Lock()
	defer printMutex.Unlock() // ensures the lock is released when this function finishes

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