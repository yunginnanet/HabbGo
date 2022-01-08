package composers

import (
	"strconv"
	"strings"

	navigator2 "github.com/jtieri/HabbGo/game/navigator"
	room2 "github.com/jtieri/HabbGo/game/room"
	"github.com/jtieri/HabbGo/models"
	"github.com/jtieri/HabbGo/protocol/packets"
)

func ComposeNavNodeInfo(player models.Player, cat *navigator2.Category, nodeMask bool, subcats []*navigator2.Category,
	rooms []*room2.Room, currentVisitors int, maxVisitors int) *packets.OutgoingPacket {
	p := packets.NewOutgoing(headerID(220)) // Base64 Header C\

	p.WriteBool(nodeMask) // hideCategory
	p.WriteInt(cat.Id)

	if cat.Public {
		p.WriteInt(0)
	} else {
		p.WriteInt(2)
	}

	p.WriteString(cat.Name)
	p.WriteInt(currentVisitors)
	p.WriteInt(maxVisitors)
	p.WriteInt(cat.Pid)

	if !cat.Public {
		p.WriteInt(len(rooms))
	}

	for _, r := range rooms {
		if r.Details.Owner_Id == 0 { // if r is public
			desc := r.Details.Desc

			var door int
			if strings.Contains(desc, "/") {
				data := strings.Split(desc, "/")
				desc = data[0]
				door, _ = strconv.Atoi(data[1])
			}

			p.WriteInt(r.Details.Id + room2.PublicRoomOffset) // writeInt roomId
			p.WriteInt(1)                                     // writeInt 1
			p.WriteString(r.Details.Name)                     // writeString roomName
			p.WriteInt(r.Details.CurrentVisitors)             // writeInt currentVisitors
			p.WriteInt(r.Details.MaxVisitors)                 // writeInt maxVisitors
			p.WriteInt(r.Details.CatId)                       // writeInt catId
			p.WriteString(desc)                               // writeString roomDesc
			p.WriteInt(r.Details.Id)                          // writeInt roomId
			p.WriteInt(door)                                  // writeInt door
			p.WriteString(r.Details.CCTs)                     // writeString roomCCTs
			p.WriteInt(0)                                     // writeInt 0
			p.WriteInt(1)                                     // writeInt 1
		} else {
			p.WriteInt(r.Details.Id)
			p.WriteString(r.Details.Name)

			// TODO check that player is owner of r, that r is showing owner name, or that player has right SEE_ALL_ROOMOWNERS
			if player.Details().Username() == r.Details.Owner_Name {
				p.WriteString(r.Details.Owner_Name)
			} else {
				p.WriteString("-")
			}

			p.WriteString(room2.AccessType(r.Details.AccessType))
			p.WriteInt(r.Details.CurrentVisitors)
			p.WriteInt(r.Details.MaxVisitors)
			p.WriteString(r.Details.Desc)
		}
	}

	// iterate over sub-categories
	for _, subcat := range subcats {
		if subcat.MinRankAccess > 1 {
			continue
		}

		p.WriteInt(subcat.Id)
		p.WriteInt(0)
		p.WriteString(subcat.Name)
		p.WriteInt(navigator2.CurrentVisitors(subcat)) // writeInt currentVisitors
		p.WriteInt(navigator2.MaxVisitors(subcat))     // writeInt maxVisitors
		p.WriteInt(cat.Id)
	}

	return p
}
