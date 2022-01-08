package models

type Session interface {
	Listen()
	Send(packet OutgoingPacket)
	Queue(packet OutgoingPacket)
	Flush(packet OutgoingPacket)
	Address() string
	GetPacketHandler(headerID Packet) (Handler, bool)
	Close()
}
