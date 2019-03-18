package gameserver

import "errors"

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
	AddPlayer(string) error
	//TODO : RemovePlayer
}

//------------------------------------------------------------------
//-------------------FUNCTIONS--------------------------------------
//------------------------------------------------------------------

//NewRoom : returns a new room given the key and max players
func NewRoom(s Server, key string, maxP int) Room {
	var r room
	r.key = key
	r.maxPlayers = maxP
	r.numConnected = 0
	r.players = make([]string, 0)
	r.server = s
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
	server       Server
}

func (r *room) FreeSpots() int {
	s := r.maxPlayers - r.numConnected
	return s
}
func (r *room) GetKey() string {
	return r.key
}
func (r *room) BroadcastMessage(m Message) error {
	for _, v := range r.players {
		r.server.SendMessageToAddr(m, v)
	}
	return nil
}
func (r *room) AddPlayer(key string) error {
	for _, v := range r.players {
		if v == key {
			return errors.New("player already in the room")
		}

	}
	r.players = append(r.players, key)
	return nil
}
