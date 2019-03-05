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
	SendMessage(Message)
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
	//trnansform Transform
}

func (p *player) SendMessage(m Message) {
	m.Send(nil, p.conn)
	fmt.Println("sending")
}
