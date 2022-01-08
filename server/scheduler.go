package server

import (
	"github.com/jtieri/HabbGo/game/player"
	logger "github.com/jtieri/HabbGo/log"
	"github.com/jtieri/HabbGo/protocol/packets"
)

func Handle(p *player.Player, packet *packets.IncomingPacket) {
	handler, found := p.Session.GetPacketHandler(packet.HeaderID)

	if found {
		logger.LogIncomingPacket(p.Session.Address(), handler, packet)
		handler.Run(p, packet)
	} else {
		logger.LogUnknownPacket(p.Session.Address(), packet)
	}

}
