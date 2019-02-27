package gameserver

import "net"

//------------------------------------------------------------------
//-------------------TYPES------------------------------------------
//------------------------------------------------------------------

//------------------------------------------------------------------
//-------------------INTERFACES-------------------------------------
//------------------------------------------------------------------

//Player : the player interface
type Player interface {
	//GetTransform() Transform
	SendMessage(MessageType, MessageContent)
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

func (p *player) SendMessage(mt MessageType, mc MessageContent) {

}
