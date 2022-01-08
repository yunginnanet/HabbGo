package player

import (
	"github.com/jtieri/HabbGo/habbgo/protocol/packets"
)

type Player struct {
	Session Session
	Details *Details
}

type Details struct {
	Id           int
	Username     string
	Figure       string
	Sex          string
	Motto        string
	ConsoleMotto string
	Tickets      int
	PoolFigure   string
	Film         int
	Credits      int
	LastOnline   string
	Badges       []string
	CurrentBadge string
	DisplayBadge bool
	SoundEnabled int
}

type Session interface {
	Listen()
	Send(packet *packets.OutgoingPacket)
	Queue(packet *packets.OutgoingPacket)
	Flush(packet *packets.OutgoingPacket)
	Address() string
	GetPacketHandler(headerID Packet) (Handler, bool)
	Close()
}

type Packet interface {
	Int() int
}

type Handler interface {
	Run(*Player, *packets.IncomingPacket)
}

func New(session Session) *Player {
	return &Player{
		Session: session,
		Details: &Details{},
	}
}
