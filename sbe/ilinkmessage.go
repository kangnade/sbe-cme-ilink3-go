package sbe

/*
* Defines the ilink3 message interface
* IlinkInbound, IlinkOutbound, and IlinkMessage
* IlinkInbound interface: implements Decode, GetMessageTypeID, and PrettyPrint methods
* InlinkOutbound interface: implements Encode, GetMessageTypeID, and PrettyPrint methods
* InlinkMessage interface: implements Encode, Decode, GetMessageTypeID and PrettyPrint methods
 */

type ILinkOutbound interface{
	Encode(c *Coder)
	GetMessageTypeID() UInt16
	BlockLength() UInt16
	PrettyPrint()
}

type ILinkInbound interface{
	Decode(c *Coder)
	GetMessageTypeID() UInt16
	BlockLength() UInt16
	PrettyPrint()
}

type ILinkMessage interface{
	// interface embedding
	ILinkInbound
	ILinkOutbound
}

// The MessageWireSize interface is implemented by message types that have variable length DAT fields
// e.g. the Credentials DATA type field in Negotitate500 message
// builder.go checks for this interface to determine the correct overall buffer size.
// Message types in SBE that do not have the Credentials DATA fields do not need to implement this interface.
type MessageWireSize interface{
	WireSize() int
}