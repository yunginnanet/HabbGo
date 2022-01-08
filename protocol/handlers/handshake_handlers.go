package handlers

import (
	player2 "github.com/jtieri/HabbGo/game/player"
	"github.com/jtieri/HabbGo/models"
	"github.com/jtieri/HabbGo/protocol/composers"
)

func InitCrypto(player models.Player, packet models.IncomingPacket) {
	player.Session().Send(composers.ComposeCryptoParams())
}

func GenerateKey(player models.Player, packet models.IncomingPacket) {
	player.Session().Send(composers.ComposeAvailableSets())
	player.Session().Send(composers.ComposeEndCrypto())
	// player.Session().Send(composers.ComposeSecretKey())
}

func GetSessionParams(player models.Player, packet models.IncomingPacket) {
	player.Session().Send(composers.ComposeSessionParams())
}

func VersionCheck(player models.Player, packet models.IncomingPacket) {

}

func UniqueID(player models.Player, packet models.IncomingPacket) {

}

func SECRETKEY(player models.Player, packets models.IncomingPacket) {
	player.Session().Send(composers.ComposeEndCrypto())
}

func SSO(p models.Player, packet models.IncomingPacket) {
	token := packet.ReadString()

	// TODO if p login with token is success login, otherwise send LOCALISED ERROR & disconnect from server
	if token == "" {
		player2.Login(p)
	} else {

	}
}

func TRY_LOGIN(p models.Player, packet models.IncomingPacket) {
	username := packet.ReadString()
	password := packet.ReadString()

	if player2.LoginDB(p, username, password) {
		player2.Login(p)
		p.Session().Send(composers.ComposeLoginOk())
	} else {
		// TODO send LOCALISED ERROR
	}
}
