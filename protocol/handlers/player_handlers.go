package handlers

import (
	"github.com/yunginnanet/habbgo/models"
	"github.com/yunginnanet/habbgo/protocol/composers"
)

func GetInfo(player models.Player, packet models.IncomingPacket) {
	player.Session().Send(composers.ComposeUserObj(player))
}

func GetCredits(player models.Player, packet models.IncomingPacket) {
	player.Session().Send(composers.ComposeCreditBalance(player.Details().Credits()))
}

func GetAvailableBadges(player models.Player, packet models.IncomingPacket) {
	player.Session().Send(composers.ComposeAvailableBadges(player))
}

func GetSoundSetting(player models.Player, packet models.IncomingPacket) {
	player.Session().Send(composers.ComposeSoundSetting(player.Details().SoundEnabled()))
}

func TestLatency(player models.Player, packet models.IncomingPacket) {
	l := packet.ReadInt()
	player.Session().Send(composers.ComposeLatency(l))
}
