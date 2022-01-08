package composers

import (
	"github.com/jtieri/HabbGo/protocol/packets"
)

func DATE(date string) *packets.OutgoingPacket {
	p := packets.NewOutgoing(headerID(163)) // Base64 Header - Bc
	p.Write(date)
	return p
}

func APPROVENAMEREPLY(approveCode int) *packets.OutgoingPacket {
	p := packets.NewOutgoing(headerID(36)) // Base64 Header -
	p.WriteInt(approveCode)
	return p
}

func NAMEUNACCEPTABLE() *packets.OutgoingPacket {
	p := packets.NewOutgoing(headerID(37)) // Base64 Header -
	p.WriteInt(0)
	return p
}

func PASSWORD_APPROVED(errorCode int) *packets.OutgoingPacket {
	p := packets.NewOutgoing(headerID(282)) // Base64 Header - DZ
	p.WriteInt(errorCode)
	return p
}

func EMAIL_APPROVED() *packets.OutgoingPacket {
	p := packets.NewOutgoing(headerID(271)) // Base64 Header -
	return p
}

func EMAIL_REJECTED() *packets.OutgoingPacket {
	p := packets.NewOutgoing(headerID(272)) // Base64 Header -
	return p
}
