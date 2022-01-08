package player

import (
	packets2 "github.com/jtieri/HabbGo/protocol/packets"
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
	Send(packet *packets2.OutgoingPacket)
	Queue(packet *packets2.OutgoingPacket)
	Flush(packet *packets2.OutgoingPacket)
	Address() string
	GetPacketHandler(headerID Packet) (Handler, bool)
	Close()
}

type Packet interface {
	Int() int
}

type Handler interface {
	Run(*Player, *packets2.IncomingPacket)
}

func New(session Session) *Player {
	return &Player{
		Session: session,
		Details: &Details{},
	}
}
