package main

import (
	"github.com/rs/zerolog"

	"github.com/yunginnanet/habbgo/config"
	"github.com/yunginnanet/habbgo/server"
)

var log zerolog.Logger

func init() {
	config.LoadConfig()
	log = config.StartLogger()
}

func main() {
	log.Info().Msg("Booting up BobbaGo...")
	log.Info().Str("caller", config.Filename).Msg("Loaded config file...")

	// TODO: replace with bitcask database
	//	log.Println("Attempting to make connection with the database... ")
	//	host := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Name)


	log.Info().Str("caller", config.HabboBind).Int("port", config.HabboPort).
		Msg("Starting the game server... ")

	gameServer := server.New()
	gameServer.Start()

	defer gameServer.Stop()
}
