package player

import "github.com/jtieri/HabbGo/models"

type Player struct {
	session models.Session
	details models.PlayerDetails
}

func (p *Player) Session() models.Session {
	return p.session
}

func (p *Player) Details() models.PlayerDetails {
	return p.details
}

// Details encapsulates a player's details.
type Details struct {
	id int

	// TODO: Why is LastOnline a string?
	username, figure, sex, motto, consoleMotto,
	poolFigure, lastOnline, currentBadge string

	tickets, film, credits, soundEnabled int

	badges       []string
	displayBadge bool
}

func (d Details) ID() int {
	return d.id
}

func (d Details) Username() string {
	return d.username
}

func (d Details) Figure() string {
	return d.figure
}

func (d Details) Sex() string {
	return d.sex
}

func (d Details) Motto() string {
	return d.motto
}

func (d Details) ConsoleMotto() string {
	return d.consoleMotto
}

func (d Details) PoolFigure() string {
	return d.poolFigure
}

func (d Details) LastOnline() string {
	return d.lastOnline
}

func (d Details) CurrentBadge() string {
	return d.currentBadge
}

func (d Details) Tickets() int {
	return d.tickets
}

func (d Details) Film() int {
	return d.film
}

func (d Details) Credits() int {
	return d.credits
}

func (d Details) SoundEnabled() int {
	return d.soundEnabled
}

func (d Details) Badges() []string {
	return d.badges
}

func (d Details) DisplayBadge() bool {
	return d.displayBadge
}

func New(session models.Session) models.Player {
	return &Player{
		session: session,
		details: &Details{},
	}
}
