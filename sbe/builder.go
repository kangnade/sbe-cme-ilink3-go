package sbe

import "fmt"

/*
* builder.go handles outgoing message to the CME
* builde.go takes a filled ILinkOutbound message struct:
* writes the SBE MessageHeader, encodes the message body into
* raw bytes that will be sent to CME via TCP
*
* MessageHeader is defined in the sbe package's composites.go as follows:
* type MessageHeader struct {
*	BlockLength UInt16
*	TemplateId  UInt16
*	SchemaId    UInt16
*	Version     UInt16
* }
 */

// SBE schema constants defined in ilinkbinary.xml
const(
	ILinkSchemaID 			UInt16 = 1
	ILinkVersion				UInt16 = 9
)

// encodeHeader writes the 8 bytes SBE MessageHeader into the Coder
func encodeHeader(c *Coder, msg ILinkOutbound){
	blockLength := msg.BlockLength()
	templateID := msg.GetMessageTypeID()
	schemaID := ILinkSchemaID
	version := ILinkVersion
	c.Encode(&blockLength)
	c.Encode(&templateID)
	c.Encode(&schemaID)
	c.Encode(&version)
}

// Build encodes an ILinkOutbound message into bytes stream to send to CME
func Build(msg ILinkOutbound) ([]byte, error){
	
	// Check if the outbound message type has Credential DATA field
	// if yes, we need to call WireSize() to get the total buffer size
	var buffSize int
	if ws, ok := msg.(MessageWireSize); ok{
		buffSize = ws.WireSize()
	}else{
		// Fixed message type requires MessageHeader and BlockLength
		buffSize = 8 + int(msg.BlockLength())
	}

	if buffSize < 8 {
		return nil, fmt.Errorf("Build: invalid buffSize: %d for templateID: %d", buffSize, msg.GetMessageTypeID(),)
	}

	// from coder.go, initialize the Coder using Encoder() with the desired buffSize, Encoder returns a pointer to Coder c
	c := Encoder(buffSize)

	// Encode the 8 bytes MessageHeader by calling the internal encodeHeader()
	encodeHeader(c, msg)

	// encode the rest of the message into byte stream
	msg.Encode(c)

	return c.GetBytes(), nil
}