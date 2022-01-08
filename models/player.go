package models

type PlayerDetails interface {
	ID() int
	Username() string
	Figure() string
	Sex() string
	Motto() string
	ConsoleMotto() string
	PoolFigure() string
	LastOnline() string
	CurrentBadge() string
	Tickets() int
	Film() int
	Credits() int
	SoundEnabled() int
	Badges() []string
	DisplayBadge() bool
}

type Player interface {
	Session() Session
	Details() PlayerDetails
}
