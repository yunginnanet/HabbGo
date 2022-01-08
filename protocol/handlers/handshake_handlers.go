package handlers

import (
	player2 "github.com/jtieri/HabbGo/game/player"
	"github.com/jtieri/HabbGo/protocol/composers"
	"github.com/jtieri/HabbGo/protocol/packets"
)

func InitCrypto(player *player2.Player, packet *packets.IncomingPacket) {
	player.Session.Send(composers.ComposeCryptoParams())
}

func GenerateKey(player *player2.Player, packet *packets.IncomingPacket) {
	player.Session.Send(composers.ComposeAvailableSets())
	player.Session.Send(composers.ComposeEndCrypto())
	// player.Session.Send(composers.ComposeSecretKey())
}

func GetSessionParams(player *player2.Player, packet *packets.IncomingPacket) {
	player.Session.Send(composers.ComposeSessionParams())
}

func VersionCheck(player *player2.Player, packet *packets.IncomingPacket) {

}

func UniqueID(player *player2.Player, packet *packets.IncomingPacket) {

}

func SECRETKEY(player *player2.Player, packets *packets.IncomingPacket) {
	player.Session.Send(composers.ComposeEndCrypto())
}

func SSO(p *player2.Player, packet *packets.IncomingPacket) {
	token := packet.ReadString()

	// TODO if p login with token is success login, otherwise send LOCALISED ERROR & disconnect from server
	if token == "" {
		player2.Login(p)
	} else {

	}
}

func TRY_LOGIN(p *player2.Player, packet *packets.IncomingPacket) {
	username := packet.ReadString()
	password := packet.ReadString()

	if player2.LoginDB(p, username, password) {
		player2.Login(p)
		p.Session.Send(composers.ComposeLoginOk())
	} else {
		// TODO send LOCALISED ERROR
	}
}
