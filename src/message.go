package gameserver

import (
	"fmt"
	"reflect"
)

//------------------------------------------------------------------
//-------------------TYPES------------------------------------------
//------------------------------------------------------------------

//MessageType : the type of message to send
type MessageType byte

//MessageContent : the content of the message to send
type MessageContent interface{}

//------------------------------------------------------------------
//-------------------CONSTANTS--------------------------------------
//------------------------------------------------------------------
const (
	StrangeMessage MessageType = 1

	VoidMessage MessageType = 3

	ForceTransform MessageType = 5

	SimpleTransform MessageType = 10

	CompleteTransform MessageType = 12

	VarString MessageType = 70

	VarByte MessageType = 80

	Vector2byte MessageType = 82

	Vector3byte MessageType = 84

	VarInt MessageType = 90

	Vector2int MessageType = 92

	Vector3int MessageType = 94

	VarFloat MessageType = 96

	Vector2float MessageType = 98

	Vector3float MessageType = 100

	VarInt64 MessageType = 102

	Vector2int64 MessageType = 104

	Vector3int64 MessageType = 106
)

//------------------------------------------------------------------
//-------------------INTERFACES-------------------------------------
//------------------------------------------------------------------

//Message : the interface of the message
type Message interface {
	GetType() MessageType
	GetContent() MessageContent
}

//------------------------------------------------------------------
//-------------------FUNCTIONS--------------------------------------
//------------------------------------------------------------------

//NewMessage : creates a new message
func NewMessage(mt MessageType, mc MessageContent) Message {
	var m message
	m.mType = mt
	m.mContent = mc

	fmt.Println("type ", mt, "content", reflect.TypeOf(mc))
	return &m
}

//------------------------------------------------------------------
//-------------------STRUCTS----------------------------------------
//------------------------------------------------------------------
type message struct {
	mType    MessageType
	mContent MessageContent
}

//-----------------------------------------------------------------
func (m *message) GetType() MessageType {
	t := m.mType
	return t
}
func (m *message) GetContent() MessageContent {
	c := m.mContent
	return c
}
