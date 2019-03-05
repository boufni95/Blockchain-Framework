package gameserver

import (
	"fmt"
	"net"
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
	//WelcomeMessage : message to send data to the client
	//STRUCTURE:
	//1byte -> message type
	//4byte -> client code
	WelcomeMessage MessageType = 0 // s-> c

	//StrangeMessage : message unknown to the clinet
	//STRUCTURE:
	//1byte -> message type
	//1byte -> (last peace)?0:10
	//1byte -> peace type
	//N1bytes -> in some cases, peace length (es string)
	//N2bytes ->peace
	StrangeMessage MessageType = 1 // s -> c

	//VoidMessage : message containing no information
	//STRUCTURE:
	//1byte -> message type
	VoidMessage MessageType = 3 // s <-> c

	//ForceTransform : forces the client to set the players
	//transform to the value sent
	//STRUCTURE:
	//1byte -> message type
	//12byte -> Vector3int position
	ForceTransform MessageType = 5 // s -> c

	//SimpleTransform : mesage containing a simple transform
	//STRUCTURE:
	//1byte -> message type
	//4byte -> owner
	//12byte -> Vector3int position
	SimpleTransform MessageType = 10 // s <-> c

	//CompleteTransform : message containing a coplete transform
	//STRUCTURE:
	//1byte -> message type
	//
	CompleteTransform MessageType = 12 // s <-> c

	//StringName : message containing the name of the player
	//STRUCTURE:
	//1byte -> message type
	//1byte -> name lenght
	//Nbyte -> name
	StringName MessageType = 18 // s <-> c

	//NewConnection : a new player connected
	//STRUCTURE:
	//1byte -> message type
	//4byte -> owner
	//1byte -> name length
	//Nbyte -> name
	NewConnection MessageType = 20 // s -> c

	//NewDisconnection : a player disconnected
	//STRUCTURE:
	//1byte -> message type
	//4byte -> owner
	NewDisconnection MessageType = 22 // s -> c

	//NewInRoom : a player just connected to the room
	//STRUCTURE:
	//1byte : message type
	//4byte : owner
	//1byte -> name length
	//Nbyte -> name
	NewInRoom MessageType = 24 // s-> c

	//NewOutRoom : a player exited the room
	//STRUCTURE:
	//1byte -> message type
	//4byte -> owner
	NewOutRoom MessageType = 26 // s-> c

	//ChatAll : send text message to all players
	//STRUCTURE:
	//1byte -> message type
	//1byte -> owner name lenght
	//N1byte -> owner name
	//1byte -> text message lenth
	//N2byte -> text message
	ChatAll MessageType = 30 // s <-> c better to s -> c

	//ChatRoom : send text message to all in the room
	//STRUCTURE:
	//1byte -> message type
	//4byte -> owner
	//1byte -> text message lenght
	//Nbyte -> text message
	ChatRoom MessageType = 32 // s <-> c

	//ChatTo : send text message to specific user
	//STRUCTURE:
	//1byte -> message type
	//4byte -> owner
	//4byte -> destination
	//1byte -> text message length
	//Nbyte -> text message
	ChatTo MessageType = 34 // s <-> c

	//These are used to build the message content

	//VarString : sends a string with a numbr of bytes
	//equal to string lenght +1
	VarString MessageType = 70 // s <-> c

	//VarByte : send a byte
	VarByte MessageType = 80 // s <-> c

	//Vector2byte : send a vector of 2 bytes
	Vector2byte MessageType = 82 // s <-> c

	//Vector3byte : send a vector of 3 bytes
	Vector3byte MessageType = 84 // s <-> c

	//VarInt : send 32bit int
	VarInt MessageType = 90 // s <-> c

	//Vector2int : send a vector of 2 32bit int
	Vector2int MessageType = 92 // s <-> c

	//Vector3int : send a vector of 3 32bit int
	Vector3int MessageType = 94 // s <-> c

	//the float messages sends a 32bit (float, but on the
	//net travels an) int, wich is the input float * 100
	//it then is reconverted on the client by doing int / 100.0
	//this is to prevent problems since different arciteture implement
	//floats differently

	//VarFloat : send a 32bit float
	VarFloat MessageType = 96 // s <-> c

	//Vector2float : send a vector of 2 32bit float
	Vector2float MessageType = 98 // s <-> c

	//Vector3float : send a vector of 3 32bit float
	Vector3float MessageType = 100 // s <-> c

	//the int 64 messages are useful to send big ints
	//or they can be used to send floar with a higher
	//precision, they are note used in the package
	//but they are implemented so they are usable

	//to send float using int64 make a moltiplication
	//of the float number with a number as high as the
	//precision you want
	//then divede with the same number on the client

	//VarInt64 : send 64bit int
	VarInt64 MessageType = 102 // s <-> c

	//Vector2int64 : send a vector of 2 64bit int
	Vector2int64 MessageType = 104 // s <-> c

	////Vector3int64 : send a vector of 3 64bit int
	Vector3int64 MessageType = 106 // s <-> c
)

//------------------------------------------------------------------
//-------------------INTERFACES-------------------------------------
//------------------------------------------------------------------

//Message : the interface of the message
type Message interface {
	GetType() MessageType
	GetContent() MessageContent
	Send(Server, net.Conn)
	Mutate(MessageType, MessageContent)
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

func reciveMess(m Message, conn net.Conn) {
	switch m.GetType() {
	case VoidMessage:
		{

		}
	case SimpleTransform:
		{

		}
	case CompleteTransform:
		{

		}
	case StringName:
		{

		}
	case ChatAll:
		{

		}
	case ChatRoom:
		{

		}
	case ChatTo:
		{

		}
	}
}
func extractContent(mt MessageType, b []byte) MessageContent {
	switch mt {
	case VarString:
		{

		}
	case VarByte:
		{

		}
	case Vector2byte:
		{

		}
	case Vector3byte:
		{

		}
	case VarInt:
		{

		}
	case Vector2int:
		{

		}
	case Vector3int:
		{

		}
	case VarFloat:
		{

		}
	case Vector2float:
		{

		}
	case Vector3float:
		{

		}
	case VarInt64:
		{

		}
	case Vector2int64:
		{

		}
	case Vector3int64:
		{

		}
	}

	return nil
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
func (m *message) Send(s Server, conn net.Conn) {
	switch m.mType {
	case WelcomeMessage:
		{
			Bytes := make([]byte, 4)

			intTo4Byte(&Bytes, m.mContent.(int), true)
			toSend := make([]byte, 1)
			toSend[0] = (byte)(m.mType)
			toSend = append(toSend, Bytes...)
			conn.Write(toSend)
			break
		}
	case NewConnection:
		{
			s := m.mContent.(struct {
				code int
				name []byte
			})
			toSend := make([]byte, 1)
			toSend[0] = (byte)(m.mType)
			Bytes := make([]byte, 4)
			intTo4Byte(&Bytes, s.code, true)
			toSend = append(toSend, Bytes...)
			toSend = append(toSend, (byte)(len(s.name)))
			toSend = append(toSend, s.name...)
			//spew.Dump(toSend)
			conn.Write(toSend)
			break
		}
	}
}
func (m *message) Mutate(mt MessageType, mc MessageContent) {
	m.mType = mt
	m.mContent = mc
}
