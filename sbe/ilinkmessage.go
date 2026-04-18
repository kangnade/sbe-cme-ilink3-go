package sbe

/*
* Defines the ilink3 message interface
* IlinkInbound, IlinkOutbound, and IlinkMessage
* IlinkInbound interface: implements Decode, GetMessageTypeID, and PrettyPrint methods
* InlinkOutbound interface: implements Encode, GetMessageTypeID, and PrettyPrint methods
* InlinkMessage interface: implements Encode, Decode, GetMessageTypeID and PrettyPrint methods
 */

type IlinkOutbound interface{
	Encode(c *Coder)
	GetMessageTypeID() UInt16
	PrettyPrint()
}

type IlinkInbound interface{
	Decode(c *Coder)
	GetMessageTypeID() UInt16
	PrettyPrint()
}

type IlinkMessage interface{
	// interface embedding
	IlinkInbound
	IlinkOutbound
}

