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

	//StrangeMessage : message unknown to the clinet
	StrangeMessage MessageType = 1

	//VoidMessage : message containing no information, useful
	//to ckeck timeout
	VoidMessage MessageType = 3

	//ForceTransform : forces the client to set the players
	//transform to the value sent
	ForceTransform MessageType = 5

	//SimpleTransform : mesage containing a simple transform
	SimpleTransform MessageType = 10

	//CompleteTransform : message containing a coplete transform
	CompleteTransform MessageType = 12

	//These constants are used by the strange message

	//VarString : sends a string with a numbr of bytes
	//equal to string lenght +1
	VarString MessageType = 70

	//VarByte : send a byte
	VarByte MessageType = 80

	//Vector2byte : send a vector of 2 bytes
	Vector2byte MessageType = 82

	//Vector3byte : send a vector of 3 bytes
	Vector3byte MessageType = 84

	//VarInt : send 32bit int
	VarInt MessageType = 90

	//Vector2int : send a vector of 2 32bit int
	Vector2int MessageType = 92

	//Vector3int : send a vector of 3 32bit int
	Vector3int MessageType = 94

	//the float messages sends a 32bit float, but on the
	//net travels an int, wich is the input float * 100
	//it then is reconverted on the client by doing int / 100.0
	//this is to prevent problems since different arciteture implement
	//floats differently

	//VarFloat : send a 32bit float
	VarFloat MessageType = 96

	//Vector2float : send a vector of 2 32bit float
	Vector2float MessageType = 98

	//Vector3float : send a vector of 3 32bit float
	Vector3float MessageType = 100

	//the int 64 messages are useful to send big ints
	//or they can be used to send floar with a higher
	//precision, they are note used in the package
	//but they are implemented so they are usable

	//to send float using int64 make a moltiplication
	//of the float number with a number as high as the
	//precision you want
	//then divede with the same number on the client

	//VarInt64 : send 64bit int
	VarInt64 MessageType = 102

	//Vector2int64 : send a vector of 2 64bit int
	Vector2int64 MessageType = 104

	////Vector3int64 : send a vector of 3 64bit int
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
