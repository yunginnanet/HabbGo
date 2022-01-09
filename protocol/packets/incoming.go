package packets

import (
	"bytes"

	encoding2 "github.com/yunginnanet/habbgo/protocol/encoding"
)

// IncomingPacket represents a client->server packet.
type IncomingPacket struct {
	header   string
	headerID PacketID
	payload  *bytes.Buffer
}

func (packet *IncomingPacket) Int() int {
	return packet.headerID.Int()
}

func (packet *IncomingPacket) Header() string {
	return packet.header
}

func (packet *IncomingPacket) HeaderID() int {
	return packet.headerID.Int()
}

func (packet *IncomingPacket) Payload() *bytes.Buffer {
	return packet.payload
}

type PacketID int

func (pid PacketID) Int() int {
	return int(pid)
}

// NewIncoming returns a pointer to a newly allocated IncomingPacket struct with its appropriate header information.
func NewIncoming(rawHeader []byte, payload *bytes.Buffer) *IncomingPacket {
	return &IncomingPacket{
		header:   string(rawHeader),
		headerID: PacketID(encoding2.DecodeB64(rawHeader)),
		payload:  payload,
	}
}

// ReadB64 reads two bytes from the packets buffer and returns their Base64 decoded value as an integer.
func (packet *IncomingPacket) ReadB64() int {
	data := make([]byte, 2)
	data[0], _ = packet.Payload().ReadByte()
	data[1], _ = packet.Payload().ReadByte()
	return encoding2.DecodeB64(data)
}

// ReadBytes advances the packets buffer i bytes and returns those i bytes in a slice.
func (packet *IncomingPacket) ReadBytes(i int) []byte {
	data := packet.Payload().Next(i)
	return data
}

// ReadInt reads one integer from the packets buffer by decoding a Vl64 encoded sequence of bytes.
func (packet *IncomingPacket) ReadInt() int {
	data := packet.Bytes()
	length := int(data[0] >> 3 & 7)
	value := encoding2.DecodeVl64(data)
	packet.ReadBytes(length)
	return value
}

// ReadBool reads one integer from the packets buffer and if it equals 1 returns true, otherwise returns false.
func (packet *IncomingPacket) ReadBool() bool {
	return packet.ReadInt() == 1
}

// ReadString reads two bytes from the packets buffer to get a length of n and then returns a string of n bytes.
func (packet *IncomingPacket) ReadString() string {
	length := packet.ReadB64()
	message := packet.ReadBytes(length)
	return string(message)
}

// String returns the remaining bytes in the packets buffer as a string.
func (packet *IncomingPacket) String() string {
	return string(packet.Bytes())
}

// Bytes returns a slice containing the remaining bytes in the packets buffer.
func (packet *IncomingPacket) Bytes() []byte {
	return packet.Payload().Bytes()
}
