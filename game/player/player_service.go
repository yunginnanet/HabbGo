package player

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/jtieri/HabbGo/models"
)

func Login(player models.Player) {
	// Set player logged in & ping ready for latency test
	// Possibly add player to a list of online players? Health endpoint with server stats?
	// Save current time to Conn for players last online time

	// Check if player is banned & if so send USER_BANNED
	// Log IP address to Conn

	LoadBadges(player)

	// If Config has alerts enabled, send player ALERT

	// Check if player gets club gift & update club status
}

func register(username, figure, gender, email, birthday, createdAt, password string, salt []byte) {
	err := Register(username, figure, gender, email, birthday, createdAt, password, salt)
	if err != nil {
		log.Error().Err(err)
	}
}

func (p *Player) LogErr(err error) {
	fmt.Printf("[%d-%s] Player encountered error: %e \n", p.Details().ID(), p.Details().Username(), err)
}
