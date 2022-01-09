package composers

import (
	"strconv"

	"github.com/yunginnanet/habbgo/models"
	"github.com/yunginnanet/habbgo/protocol/packets"
)

func ComposeUserObj(p models.Player) models.OutgoingPacket {
	packet := packets.NewOutgoing(headerID(5)) // Base64 Header @E

	packet.WriteString(strconv.Itoa(p.Details().ID()))
	packet.WriteString(p.Details().Username())
	packet.WriteString(p.Details().Figure())
	packet.WriteString(p.Details().Sex())
	packet.WriteString(p.Details().Motto())
	packet.WriteInt(p.Details().Tickets())
	packet.WriteString(p.Details().PoolFigure())
	packet.WriteInt(p.Details().Film())
	// packet.WriteInt(directMail)

	return packet
}

func ComposeCreditBalance(credits int) *packets.OutgoingPacket {
	p := packets.NewOutgoing(headerID(6)) // Base64 Header @F
	p.WriteString(strconv.Itoa(credits) + ".0")
	return p
}

func ComposeAvailableBadges(p models.Player) models.OutgoingPacket {
	packet := packets.NewOutgoing(headerID(229)) // Base64 Header

	packet.WriteInt(len(p.Details().Badges()))

	var bSlot int
	for i, b := range p.Details().Badges() {
		packet.WriteString(b)

		if b == p.Details().CurrentBadge() {
			bSlot = i
		}
	}

	packet.WriteInt(bSlot)
	packet.WriteBool(p.Details().DisplayBadge())
	return packet
}

func ComposeSoundSetting(ss int) *packets.OutgoingPacket {
	p := packets.NewOutgoing(headerID(308)) // Base 64 Header Dt
	p.WriteInt(ss)
	return p
}

func ComposeLatency(l int) *packets.OutgoingPacket {
	p := packets.NewOutgoing(headerID(354)) // Base 64 Header Eb
	p.WriteInt(l)
	return p
}
