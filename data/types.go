package data

type Sex uint8

const (
	Male Sex = iota
	Female
	Neither
)

// Key/Value: Key is ID
type Player struct {
	ID           uint64
	Username     string
	Password     string // bcrypt hash
	Sex          Sex
	Figure       string //  1000118001270012900121001
	PoolFigure   string
	Motto        string
	ConsoleMotto string
	Birthday     time.Time
	Email        string
	SoundEnabled bool
	CreatedOn    time.Time
	LastOnline   time.Time
	Credits      int
	Film         int
	Tickets      int
}

// Key/Value: Key is ID
type Badges struct {
	ID   uint64
	Code string
}

// Key/Value: Key is PlayerID
type PlayerBadge struct {
	PlayerID uint64 // Directly tied to ID from Player struct
	BadgeID  uint64 // Directly tied to ID from Badges struct
	Display  bool
}

// Key/Value: Key is ID
type RoomCategorie struct {
	ID                uint64
	ParentID          uint64 // ?
	OrderID           uint64 // ?
	Name              string
	IsNode            bool
	IsPublic          bool
	IsTrading         bool
	MinRankAccess     int
	MinRankSetFlatCat int
}

// Key/Value: Key is ID
type RoomModel struct {
	ID        uint64
	name      string
	door_x    int
	door_y    int
	door_z    int64 // 10 = 1.0 - 225 = 2.25 - etc
	door_dir  int
	heightmap string
}

type RoomAccess uint8

const (
	Open RoomAccess = iota
	Password
	Closed
)

// Key/Value: Key is ID
type Room struct {
	id               uint64
	category_id      int
	name             string
	description      string
	owner_id         int
	model_id         int
	ccts             string
	wallpaper        int
	floor            int
	show_name        bool
	password         string
	access           RoomAccess
	sudo_users       bool
	current_visitors int
	max_visitors     int
	rating           int
	hidden           bool
	created_at       time.Time
	updated_at       time.Time
}
