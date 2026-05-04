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
	BlockLength() uint16
	PrettyPrint()
}

type ILinkInbound interface{
	Decode(c *Coder)
	GetMessageTypeID() UInt16
	BlockLength() uint16
	PrettyPrint()
}

type ILinkMessage interface{
	// interface embedding
	ILinkInbound
	ILinkOutbound
}

