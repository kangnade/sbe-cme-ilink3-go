package sbe

import (
	"encoding/binary"
	"encoding/hex"
	"log"
)

/*
* The ilinkbinary.xml message types are defined in the message_tyoeOfMessae.go files.
* The sbe used by CME adopts the LittleEndian byte order, and message comes in and out
* as byte stream, which is []byte in Go.
*
* However, if we define each Encode/Decode methods for the message type manually, it
* would involve manual encoding/decoding with hardcoded offsets, this is not ideal.
* The reason is that manual and hardcoded encoding and decoding is:
* 1. Error-prone (mistake on hardcoded offset)
* 2. Non-scalable given the fact we have many message types
*
* According to the binaryo package https://pkg.go.dev/encoding/binary#Decode,
* https://pkg.go.dev/encoding/binary#Encode and https://pkg.go.dev/encoding/binary#ByteOrder
* we could use the ByteOrder interface, which is implemented by binary.LittleEndian.
* With the defined Decode:
* func Decode(buf []byte, order ByteOrder, data any) (int, error)
* and Encode:
* func Encode(buf []byte, order ByteOrder, data any) (int, error)
*
* We could avoid manual encoding/decoding, and still encode/decode with desired LittleEndian
* byte order.
*
* The coder.go serves this purpose.
 */

// Coder handles encoding and decoding with
// (1) buffer []byte: takes the stream of bytes,
// (2) order binary.ByteOrder: choosing the LittleEndian order,
// (3) offset int: tracks the current position (offset)
type Coder struct{
	buffer 		[]byte
	order 		binary.ByteOrder
	offset		int
}

// The Decoder initializes a Coder with the provided byte slice for decoding
// and returns a pointer to the it.
func Decoder(data []byte) *Coder{
	return &Coder{
		buffer: data,
		order: binary.LittleEndian,
		offset: 0,
	}
}

// The Encoder initializes a Coder with the given size of a CME ilink3 message type
// specified in its blockLength and returns a pointer to it
func Encoder(size int) *Coder{
	return &Coder{
		buffer: make([]byte, size),
		order: binary.LittleEndian,
		offset: 0,
	}
}

// GetOffset returns the current offset/position of a Coder
func (c *Coder) GetOffset() int{
	return c.offset
}

// GetBytes returns the encoded buffer
func (c *Coder) GetBytes() []byte{
	return c.buffer[: c.offset]
}

// Decode decodes the given bytes and buffer them
func (c *Coder) Decode(data interface{}){
	n, err := binary.Decode(c.buffer[c.offset:], c.order, data)
	if err != nil{
		log.Printf("Decode error occured at offset: %d", c.offset)
		log.Println(hex.Dump(c.buffer))
		log.Fatal(err)
	}
	c.offset += n // move forward the offset exactly by how many bytes are consumed
}

// Encode encodes the given bytes and buffer them
func (c *Coder) Encode(data interface{}){
	n, err := binary.Encode(c.buffer[c.offset:], c.order, data)
	if err != nil{
		log.Printf("Encode error occured at: %d", c.offset)
		log.Println(hex.Dump(c.buffer))
		log.Fatal(err)
	}
	c.offset += n
}