package gameserver

//------------------------------------------------------------------
//-------------------TYPES------------------------------------------
//------------------------------------------------------------------

//RoomStatus : the type of the room status
type RoomStatus int

//------------------------------------------------------------------
//-------------------CONSTANTS--------------------------------------
//------------------------------------------------------------------

const (

	//NullRoomStatus : the room is created but not usable
	NullRoomStatus RoomStatus = 0

	//AvalRoomStatus : the room is avalaible
	//and there is space for players
	AvalRoomStatus RoomStatus = 2

	//FullRoomStatus : the room is full
	FullRoomStatus RoomStatus = 5

	//ErrorRoomStatus : the room is reporting an error
	ErrorRoomStatus RoomStatus = -1
)

//------------------------------------------------------------------
//-------------------INTERFACES-------------------------------------
//------------------------------------------------------------------

//Room : the interface of a room
type Room interface {
	FreeSpots() int
	GetKey() string
	BroadcastMessage(Message) error
}

//------------------------------------------------------------------
//-------------------FUNCTIONS--------------------------------------
//------------------------------------------------------------------

//NewRoom : returns a new room given the key and max players
func NewRoom(key string, maxP int) Room {
	var r room
	r.key = key
	r.maxPlayers = maxP
	r.numConnected = 0
	r.players = make([]string, 0)
	return &r
}

//------------------------------------------------------------------
//-------------------STRUCTS----------------------------------------
//------------------------------------------------------------------
type room struct {
	key          string
	maxPlayers   int
	numConnected int
	players      []string
}

func (r *room) FreeSpots() int {
	s := r.maxPlayers - r.numConnected
	return s
}
func (r *room) GetKey() string {
	return r.key
}
func (r *room) BroadcastMessage(m Message) error {
	return nil
}
