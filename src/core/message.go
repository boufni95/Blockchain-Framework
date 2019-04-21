package core

import (
	"fmt"
	"net"

	"github.com/davecgh/go-spew/spew"
)

//------------------------------------------------------------------
//-------------------TYPES------------------------------------------
//------------------------------------------------------------------

//MessageType : the type of message to send
type MessageType byte

//TODO : implement DEFAULT MESSAGES
//TODO : implement CUSTOM MESSAGES

//MessageContent : the content of the message to send
type MessageContent interface{}

//------------------------------------------------------------------
//-------------------CONSTANTS--------------------------------------
//------------------------------------------------------------------
const (
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
	Send(Server, net.Conn) error
	SendInGame(Server, net.Conn) error
	Mutate(MessageType, MessageContent)
	GenerateMessage() []byte
}

//------------------------------------------------------------------
//-------------------FUNCTIONS--------------------------------------
//------------------------------------------------------------------

//NewMessage : creates a new message
func NewMessage(mt MessageType, mc MessageContent) Message {
	var m message
	m.mType = mt
	m.mContent = mc

	//fmt.Println("type ", mt, "content", reflect.TypeOf(mc))
	return &m
}

//FIXME : deprecated
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
func (m *message) Send(s Server, conn net.Conn) error {
	switch m.mType {
	case WelcomeMessage:
		{

			Bytes := m.mContent.([]byte)

			//intTo4Byte(&Bytes, m.mContent.(int), true)
			toSend := make([]byte, 1)
			toSend[0] = (byte)(m.mType)
			toSend = append(toSend, Bytes...)
			conn.Write(toSend)
			fmt.Println("sendig welcome")
		}
	case StrangeMessage:
		{
			//TODO : implement StrangeMessage
		}
	case VoidMessage:
		{
			//TODO : implement VoidMessage
		}
	case InGameMessage:
		{
			Bytes := make([]byte, 1) //create the byte slice

			Bytes[0] = (byte)(m.mType) //put message type

			c := m.mContent.(Message) // extract content

			b := c.GenerateMessage() //gnerate messag

			Bytes = append(Bytes, b...) //append and send
			spew.Dump(Bytes)
			conn.Write(Bytes)
		}
		/*
			case ForceTransform:
				{
					//TODO : implement ForceTransform
				}
			case SimpleTransform: //TODO update simple transform
				{
					c := m.mContent.(struct {
						code []byte
						pos  []byte
					})
					toSend := make([]byte, 1)
					toSend[0] = (byte)(m.mType)
					toSend = append(toSend, c.code...)
					toSend = append(toSend, c.pos...)
					conn.Write(toSend)
					spew.Dump(toSend)

				}
			case CompleteTransform:
				{
					//TODO : implement CompleteTransform
				}
		*/
	case NameString:
		{
			//TODO : implement NameString
		}
	case NewConnection:
		{
			s := m.mContent.(struct {
				code []byte
				name []byte
			})
			toSend := make([]byte, 1)
			toSend[0] = (byte)(m.mType)
			Bytes := s.code
			//intTo4Byte(&Bytes, s.code, true)
			toSend = append(toSend, Bytes...)
			toSend = append(toSend, (byte)(len(s.name)))
			toSend = append(toSend, s.name...)
			//spew.Dump(toSend)
			conn.Write(toSend)

		}
	case NewDisconnection:
		{
			//TODO : implement NewDisconnection
		}
	case NewInRoom:
		{
			c := m.mContent.(struct {
				owner []byte
				pType byte
				nLen  []byte
				name  []byte
			})
			//TODO : add player type!!
			toSend := make([]byte, 1)
			toSend[0] = (byte)(m.mType)
			toSend = append(toSend, c.owner...)
			toSend = append(toSend, c.pType)
			toSend = append(toSend, c.nLen...)
			toSend = append(toSend, c.name...)
			conn.Write(toSend)
		}
	case NewOutRoom:
		{
			c := m.mContent.([]byte)
			toSend := make([]byte, 1)
			toSend[0] = (byte)(m.mType)
			toSend = append(toSend, c...)
			conn.Write(toSend)
		}
	case ChatAll:
		{
			c := m.mContent.(struct {
				nLen []byte
				name []byte
				mLen []byte
				mess []byte
			})
			toSend := make([]byte, 1)
			toSend[0] = (byte)(m.mType)
			toSend = append(toSend, c.nLen...)
			toSend = append(toSend, c.name...)
			toSend = append(toSend, c.mLen...)
			toSend = append(toSend, c.mess...)
			conn.Write(toSend)
		}
	case ChatRoom:
		{
			//TODO : implement ChatRoom
		}
	case ChatTo:
		{
			//TODO : implement ChatTo
		}
	default:
		{
			//throw error here
		}
	}
	return nil
}

//FIXME : this might be useless
func (m *message) SendInGame(s Server, conn net.Conn) error {
	switch m.mType {
	case ForceTransform:
		{

		}
	case SimpleTransform:
		{

		}
	case CompleteTransform:
		{

		}
	}
	return nil
}

//GenerateMessage : extract content from message and gives a slice of bytes
func (m *message) GenerateMessage() []byte {
	var b []byte
	switch m.mType {
	//TODO : add al cases
	case ForceTransform:
		{

		}
	case SimpleTransform:
		{
			c := m.mContent.(struct {
				code  []byte
				pType byte
				pos   []byte
			})
			b := make([]byte, 1)
			b[0] = (byte)(m.mType)
			b = append(b, c.code...)
			b = append(b, c.pType)
			b = append(b, c.pos...)
			return b
		}
	case CompleteTransform:
		{

		}
	case NewObject:
		{
			c := m.mContent.([]byte)
			b := make([]byte, 1)
			b[0] = (byte)(m.mType)
			b = append(b, c...)
			return b
		}
	case UpdateObject:
		{
			c := m.mContent.([]byte)
			b := make([]byte, 1)
			b[0] = (byte)(m.mType)
			b = append(b, c...)
			return b
		}
	case DestroyObject:
		{
			c := m.mContent.([]byte)
			b := make([]byte, 1)
			b[0] = (byte)(m.mType)
			b = append(b, c...)
			return b
		}
	default:
		{

		}
	}
	return b
}
func (m *message) Mutate(mt MessageType, mc MessageContent) {
	m.mType = mt
	m.mContent = mc
}
