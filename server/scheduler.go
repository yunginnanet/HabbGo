package server

import (
	logger "github.com/jtieri/HabbGo/log"
	"github.com/jtieri/HabbGo/models"
)

func Handle(p models.Player, packet models.IncomingPacket) {
	handler, found := p.Session().GetPacketHandler(models.Packet(packet))

	if found {
		logger.LogIncomingPacket(p.Session().Address(), handler, packet)
		handler.Run(p, packet)
	} else {
		logger.LogUnknownPacket(p.Session().Address(), packet)
	}

}
