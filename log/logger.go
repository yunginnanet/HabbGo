package log

import (
	"log"
	"reflect"
	"runtime"
	"strings"

	"github.com/jtieri/HabbGo/game/player"
	packets2 "github.com/jtieri/HabbGo/protocol/packets"
)

func LogOutgoingPacket(playerAddr string, p *packets2.OutgoingPacket) {
	log.Printf("[%v] [OUTGOING] [%v - %v]: %v ", playerAddr, p.Header, p.HeaderId, p.Payload.String())
}

func LogIncomingPacket(playerAddr string, handler player.Handler, p *packets2.IncomingPacket) {
	hName := getHandlerName(runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name())
	log.Printf("[%v] [INCOMING] [%v - %v|%v]: %v ", playerAddr, hName, p.Header, p.HeaderID, p.Payload.String())
}

func LogUnknownPacket(playerAddr string, p *packets2.IncomingPacket) {
	log.Printf("[%v] [UNK] [%v - %v]: %v ", playerAddr, p.Header, p.HeaderID, p.Payload.String())
}

func getHandlerName(handler string) string {
	sp := strings.Split(handler, "/") // e.g. github.com/jtieri/HabbGo/habbgo/protocol/handlers.GenerateKey
	s2 := sp[len(sp)-1]               // e.g. handlers.GenerateKey
	return strings.Split(s2, ".")[1]  // e.g. GenerateKey
}
