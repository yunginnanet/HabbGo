package models

import "bytes"

// OutgoingPacket represents a server->client packet.
type OutgoingPacket interface {
	Packet
	Header()   string
	HeaderID() int
	Payload()  *bytes.Buffer
	Finish()
}

// IncomingPacket represents a client->server packet.
type IncomingPacket interface {
	Packet
	Header()   string
	HeaderID() int
	Payload()  *bytes.Buffer
	ReadB64() int
	ReadBytes(int) []byte
	ReadInt() int
	ReadBool() bool
	ReadString() string
	String() string
	Bytes() []byte
}


type Packet interface {
	Int() int
}

type Handler interface {
	Run(Player, IncomingPacket)
}
