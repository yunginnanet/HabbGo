package server

import (
	"sync"

	"github.com/jtieri/HabbGo/game/player"
	handlers2 "github.com/jtieri/HabbGo/protocol/handlers"
	"github.com/jtieri/HabbGo/protocol/packets"
)

type HabboHandler func(*player.Player, *packets.IncomingPacket)

func (h HabboHandler) Run(p *player.Player, pi *packets.IncomingPacket) {
	h(p, pi)
}

type HabboPacket int

func (p HabboPacket) Int() int {
	return int(p)
}

type Router struct {
	RegisteredPackets map[player.Packet]player.Handler
	mu                sync.RWMutex
}

func (r *Router) GetPacketHandler(headerID player.Packet) (h player.Handler, found bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	h, found = r.RegisteredPackets[headerID]
	return h, found
}

func newRouter() (r *Router) {
	r = &Router{RegisteredPackets: make(map[player.Packet]player.Handler)}
	r.handshakeHandlers()
	r.registrationHandlers()
	r.playerHandlers()
	r.navigatorHandlers()
	return
}

func (r *Router) registerHandler(p HabboPacket, h HabboHandler) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.RegisteredPackets[p] = h
}

func (r *Router) bulkRegisterHandlers(pairs map[HabboPacket]HabboHandler) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for p, h := range pairs {
		r.RegisteredPackets[p] = h
	}
}

func (r *Router) handshakeHandlers() {
	var handshakes = map[HabboPacket]HabboHandler{
		206:  handlers2.InitCrypto,
		202:  handlers2.GenerateKey,  // older clients
		2002: handlers2.GenerateKey,  // newer clients
		5:    handlers2.VersionCheck, // 1170 - VERSIONCHECK in later clients? v26+? // TODO figure out exact client revisions when these HabboPacket headers change
		6:    handlers2.UniqueID,
		181:  handlers2.GetSessionParams,
		204:  handlers2.SSO,
		4:    handlers2.TRY_LOGIN,
		207:  handlers2.SECRETKEY,
	}
	r.bulkRegisterHandlers(handshakes)
}

func (r *Router) registrationHandlers() {
	var registration = map[HabboPacket]HabboHandler{
		9:   handlers2.GETAVAILABLESETS,
		49:  handlers2.GDATE,
		42:  handlers2.APPROVENAME,
		203: handlers2.APPROVE_PASSWORD,
		197: handlers2.APPROVEEMAIL,
		43:  handlers2.REGISTER,
	}
	r.bulkRegisterHandlers(registration)
}

func (r *Router) playerHandlers() {
	var playerHandlers = map[HabboPacket]HabboHandler{
		7:   handlers2.GetInfo,
		8:   handlers2.GetCredits,
		157: handlers2.GetAvailableBadges,
		228: handlers2.GetSoundSetting,
		315: handlers2.TestLatency,
	}
	r.bulkRegisterHandlers(playerHandlers)
}

func (r *Router) navigatorHandlers() {
	var navigator = map[HabboPacket]HabboHandler{150: handlers2.Navigate}
	// 151: GETUSERFLATCATS
	// 21: GETFLATINFO
	// 23: DELETEFLAT
	// 24: UPDATEFLAT
	// 25: SETFLATINFO
	// 13: SBUSYF
	// 152: GETFLATCAT
	// 153: SETFLATCAT
	// 155: REMOVEALLRIGHTS
	// 156: GETPARENTCHAIN
	// 16: SUSERF
	// 264: GET_RECOMMENDED_ROOMS
	// 17: SRCHF
	// 154: GETSPACENODEUSERS
	// 18: GETFVRF
	// 19: ADD_FAVORITE_ROOM
	// 20: DEL_FAVORITE_ROOM
	r.bulkRegisterHandlers(navigator)
}
