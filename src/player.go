package gameserver

import (
	"fmt"
	"net"
)

//------------------------------------------------------------------
//-------------------TYPES------------------------------------------
//------------------------------------------------------------------

//------------------------------------------------------------------
//-------------------INTERFACES-------------------------------------
//------------------------------------------------------------------

//Player : the player interface
type Player interface {
	//GetTransform() Transform
	SendMessage(Message) error
	BroadcastRoom(Message) error
	SetRoom(Room) error
}

//Transform : interface of the players transform
// type Transform interface {
// 	GetPosition() [3]int
// 	SetPosition([3]int)
// }

//------------------------------------------------------------------
//-------------------FUNCTIONS--------------------------------------
//------------------------------------------------------------------
// func NewPlayer(name string, conn net.Conn) Player{
// 	var p player
// 	p.name = name
// 	p.conn = conn

// 	return &p
// }
//------------------------------------------------------------------
//-------------------STRUCTS----------------------------------------
//------------------------------------------------------------------
type player struct {
	name string
	conn net.Conn
	room Room
	//trnansform Transform
}

func (p *player) SendMessage(m Message) error {
	m.Send(nil, p.conn)
	fmt.Println("sending", m.GetType())
	return nil
}
func (p *player) SetRoom(r Room) error {
	p.room = r
	return nil
}
func (p *player) BroadcastRoom(m Message) error {
	p.room.BroadcastMessage(m)
	return nil
}
