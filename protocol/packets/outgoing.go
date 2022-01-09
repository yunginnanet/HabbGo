package packets

import (
	"bytes"

	"github.com/yunginnanet/habbgo/models"
	encoding2 "github.com/yunginnanet/habbgo/protocol/encoding"
)

// OutgoingPacket represents a server->client packet.
type OutgoingPacket struct {
	header   string
	headerID int
	payload  *bytes.Buffer
}

func (packet *OutgoingPacket) Int() int {
	return packet.headerID
}

func (packet *OutgoingPacket) Header() string {
	return packet.header
}

func (packet *OutgoingPacket) HeaderID() int {
	return packet.headerID
}

func (packet *OutgoingPacket) Payload() *bytes.Buffer {
	return packet.payload
}

// NewOutgoing returns a pointer to a newly allocated OutgoingPacket struct.
// The two byte Base64 encoded header is written to the packets buffer on creation for quick composition of packets.
func NewOutgoing(headerID models.Packet) *OutgoingPacket {
	header := encoding2.EncodeB64(headerID.Int(), 2)
	return &OutgoingPacket{header: string(header), headerID: headerID.Int(), payload: bytes.NewBuffer(header)}
}

// Write will write the passed in object as a string
func (packet *OutgoingPacket) Write(s string) {
	packet.Payload().Write([]byte(s))
}

// WriteString writes a string to the packets buffer.
func (packet *OutgoingPacket) WriteString(s string) {
	packet.Payload().Write([]byte(s))
	packet.Payload().WriteByte(2) // FUSEv0.2.0 string parameter ending marker
}

// WriteInt writes a Vl64 encoded int to the packets buffer.
func (packet *OutgoingPacket) WriteInt(i int) {
	packet.Payload().Write(encoding2.EncodeVl64(i))
}

// WriteBool writes a Vl64 encoded int representing true or false to the packets buffer.
func (packet *OutgoingPacket) WriteBool(b bool) {
	if b {
		packet.WriteInt(1) // H
	} else {
		packet.WriteInt(0) // I
	}
}

// WriteValue writes a key-value entry, separated by '=',to the packets buffer.
func (packet *OutgoingPacket) WriteValue(key []byte, value []byte) {
	packet.Payload().Write(key)
	packet.Payload().Write([]byte("="))
	packet.Payload().Write(value)
	packet.Payload().WriteByte(13) // FUSEv0.2.0 key-value parameter ending marker
}

// WriteKeyValue writes a key-value pair, separated by ':', to the packets buffer.
func (packet *OutgoingPacket) WriteKeyValue(key []byte, value []byte) {
	packet.Payload().Write(key)
	packet.Payload().Write([]byte(":"))
	packet.Payload().Write(value)
	packet.Payload().WriteByte(13) // FUSEv0.2.0 key-value parameter ending marker
}

// WriteDelim writes a custom key-delimeter value to the packets buffer.
func (packet *OutgoingPacket) WriteDelim(key []byte, delim []byte) {
	packet.Payload().Write(key)
	packet.Payload().Write(delim)
}

// String returns the remaining bytes in the packets buffer as a string.
func (packet *OutgoingPacket) String() string {
	return string(packet.Payload().Bytes())
}

// Finish writes byte 0x01 to the packets buffer to signal the ending of a packet.
func (packet *OutgoingPacket) Finish() {
	packet.Payload().WriteByte(1) // FUSEv0.2.0 server->client packet ending marker
}
