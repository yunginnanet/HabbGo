package server

import (
	"sync"

	"github.com/jtieri/HabbGo/habbgo/game/player"
	"github.com/jtieri/HabbGo/habbgo/protocol/handlers"
	"github.com/jtieri/HabbGo/habbgo/protocol/packets"
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
		206: handlers.InitCrypto,
		202: handlers.GenerateKey,  // older clients
		2002: handlers.GenerateKey, // newer clients
		5: handlers.VersionCheck,   // 1170 - VERSIONCHECK in later clients? v26+? // TODO figure out exact client revisions when these HabboPacket headers change
		6: handlers.UniqueID,
		181: handlers.GetSessionParams,
		204: handlers.SSO,
		4: handlers.TRY_LOGIN,
		207: handlers.SECRETKEY,
	}
	r.bulkRegisterHandlers(handshakes)
}

func (r *Router) registrationHandlers() {
	var registration = map[HabboPacket]HabboHandler{
		9: handlers.GETAVAILABLESETS,
		49: handlers.GDATE,
		42: handlers.APPROVENAME,
		203: handlers.APPROVE_PASSWORD,
		197: handlers.APPROVEEMAIL,
		43: handlers.REGISTER,
	}
	r.bulkRegisterHandlers(registration)
}

func (r *Router) playerHandlers() {
	var playerHandlers = map[HabboPacket]HabboHandler{
		7:   handlers.GetInfo,
		8:   handlers.GetCredits,
		157: handlers.GetAvailableBadges,
		228: handlers.GetSoundSetting,
		315: handlers.TestLatency,
	}
	r.bulkRegisterHandlers(playerHandlers)
}

func (r *Router) navigatorHandlers() {
	var navigator = map[HabboPacket]HabboHandler{150: handlers.Navigate}
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
