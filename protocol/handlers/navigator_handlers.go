package handlers

import (
	"github.com/jtieri/HabbGo/game/navigator"
	room2 "github.com/jtieri/HabbGo/game/room"
	"github.com/jtieri/HabbGo/protocol/composers"
	"github.com/jtieri/HabbGo/models"
)

func Navigate(player models.Player, packet models.IncomingPacket) {
	roomService := room2.RoomService()

	nodeMask := packet.ReadInt() == 1
	catId := packet.ReadInt()

	if catId >= room2.PublicRoomOffset {
		r := roomService.RoomById(catId - room2.PublicRoomOffset)
		if r != nil {
			catId = r.Details.CatId
		}
	}

	category := navigator.NavigatorService().CategoryById(catId)

	// TODO also check that access rank isnt higher than players rank
	if category == nil {
		return
	}

	subCategories := navigator.NavigatorService().CategoriesByParentId(category.Id)
	// sort categories by player count

	currentVisitors := navigator.CurrentVisitors(category)
	maxVisitors := navigator.MaxVisitors(category)

	var rooms []*room2.Room
	if category.Public {
		for _, room := range roomService.ReplaceRooms(roomService.RoomsByPlayerId(0)) {
			if room.Details.CatId == category.Id && (!nodeMask) && room.Details.CurrentVisitors < room.Details.MaxVisitors {
				rooms = append(rooms, room)
			}
		}
	} else {
		// TODO finish private room logic
	}

	// TODO sort rooms by player count before sending NavNodeInfo

	player.Session().Send(composers.ComposeNavNodeInfo(player, category, nodeMask, subCategories, rooms, currentVisitors, maxVisitors))
}
