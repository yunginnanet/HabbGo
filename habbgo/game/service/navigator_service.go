package service

import (
	"github.com/jtieri/HabbGo/habbgo/database"
	"github.com/jtieri/HabbGo/habbgo/game/model"
	"gorm.io/gorm"
	"sync"
)

var ns *navService
var nonce sync.Once

type navService struct {
	repo *database.NavRepo
	nav  *model.Navigator
	mux  *sync.Mutex
}

// NavigatorService will initialize the single instance of navService if it is the first time it is called and then it
// will return the instance.
func NavigatorService() *navService {
	nonce.Do(func() {
		ns = &navService{
			repo: nil,
			nav:  new(model.Navigator),
			mux:  &sync.Mutex{},
		}
	})

	return ns
}

// SetDBCon is called when a NavService struct is allocated initially so that it has access to the applications db.
func (ns *navService) SetDBCon(db gorm.DB) {
	ns.repo = database.NewNavRepo(db)
}

// BuildNavigator retrieves the room categories from the database and builds the in-game Navigator with them.
//func (ns *navService) BuildNavigator() {
//	ns.nav.Categories = ns.repo.Categories()
//}

// CategoryById retrieves a navigator category given the int parameter id and returns it if there is a match.
func (ns *navService) CategoryById(id int) *model.Category {
	for _, cat := range ns.nav.Categories {
		if cat.Id == id {
			return &cat
		}
	}

	return nil
}

// CategoriesByParentId retrieves a slice of sub-categories given the int parameter pid and returns it if there is a match.
func (ns *navService) CategoriesByParentId(pid int) []*model.Category {
	var categories []*model.Category

	for _, cat := range ns.nav.Categories {
		if cat.Pid == pid {
			categories = append(categories, &cat)
		}
	}

	return categories
}

func CurrentVisitors(cat *model.Category) int {
	visitors := 0

	for _, room := range RoomService().Rooms() {
		if room.Details.CatId == cat.Id {
			visitors += room.Details.CurrentVisitors
		}
	}

	return visitors
}

func MaxVisitors(cat *model.Category) int {
	visitors := 0

	for _, room := range RoomService().Rooms() {
		if room.Details.CatId == cat.Id {
			visitors += room.Details.MaxVisitors
		}
	}

	return visitors
}
