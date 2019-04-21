package gameserver

// Review Remark: Consistently repeating switch statement is something to reconsider. Apply polymorphism and use polymorphic behaviour.
// Review Remark: Right now the flow is procedural. Make it OO flow.

import (
	"fmt"
	"net"
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
// Review Remark: Inconsisten enum for types. 0,1,3,5,18..
// Review Remark: Each if case could be a separate class.
// Review Remark: By moving each case into a class of its' own, you no longer need to comment those structure cases.
// Review Remark: You no longer need structure cases because it will be clear and short what is being done in those classes per case.
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

	//InGameMessage : defines a type of messages that are send diring game
	//STRUCTURE:
	//1byte -> message type
	//1byte -> in game message type
	//Nbyte -> message
	InGameMessage MessageType = 5
	/*------------------------------------------------------------------------------------------------------
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
	//-------------------------------------------------------------------------------------------------------*/
	//StringName : message containing the name of the player
	//STRUCTURE:
	//1byte -> message type
	//1byte -> name lenght
	//Nbyte -> name
	NameString MessageType = 18 // s <-> c

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
	NewInRoom MessageType = 24 // s <-> c

	//NewOutRoom : a player exited the room
	//STRUCTURE:
	//1byte -> message type
	//4byte -> owner
	NewOutRoom MessageType = 26 // s <-> c

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
const (
	//ForceTransform : forces the client to set the players
	//transform to the value sent
	//STRUCTURE:
	//1byte -> message type
	//12byte -> Vector3int position
	ForceTransform MessageType = 1 // s -> c

	//SimpleTransform : mesage containing a simple transform
	//STRUCTURE:
	//1byte -> message type
	//4byte -> owner
	//12byte -> Vector3int position
	SimpleTransform MessageType = 5 // s <-> c

	//CompleteTransform : message containing a coplete transform
	//STRUCTURE:
	//1byte -> message type
	//
	CompleteTransform MessageType = 7 // s <-> c
)

// Review Remark:  Don't seperate stuff in comment blocks!
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

// Review Remark: Don't use 3 lines for a single comment!
//------------------------------------------------------------------
//-------------------FUNCTIONS--------------------------------------
//------------------------------------------------------------------

// Review Remark: Redundant comment for obvious
//NewMessage : creates a new message
func NewMessage(mt MessageType, mc MessageContent) Message {
	var m message
	m.mType = mt
	m.mContent = mc
	// Review Remark: Commented out code!!
	//fmt.Println("type ", mt, "content", reflect.TypeOf(mc))
	return &m
}

//TODO : use extract content instead of doing it directly.
// Review Remark: Generics? (<Message>)
// Review Remark: Switch statement seems to prepare the content to extract for sent message send to server.
func extractContent(mt MessageType, b []byte) MessageContent {
	// Review Remark: Code smell: Many switch cases. Prefer polymorphic behaviour.
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

// Review Remark: Instead of message type, consider different object per it.
// Review Remark: Just message itself should be the content of message.go
//------------------------------------------------------------------
//-------------------STRUCTS----------------------------------------
//------------------------------------------------------------------
// Review Remark: If you extend this to go for polymorphism, you will no longer need any of the switch cases!!!
// Review Remark: OCP
type message struct {
	// Review Remark: m prefix is not needed.
	mType    MessageType
	// Review Remark: m prefix is not needed
	mContent MessageContent
}

// Review Remark: Don't separate code blocks with comments like that
//-----------------------------------------------------------------
func (m *message) GetType() MessageType {
	t := m.mType
	return t
}

func (m *message) GetContent() MessageContent {
	c := m.mContent
	return c
}
// Review Remark: End of message.go

// Review Remark: Sending a message is not the same as message itself. So it needs to be split
// Review Remark: MessageSender or Transmitter is a name for component to consider.
func (m *message) Send(s Server, conn net.Conn) error {

	// Review Remark: One is a magic number. Should be messageStartIndex = 1.
	toSend := make([]byte, 1)
	toSend[0] = (byte)(m.mType)

	// Move this to a function: PrepareByArray. Could be part of a message class (each type)
	switch m.mType {
	case WelcomeMessage:
		{
			Bytes := m.mContent.([]byte)
			//intTo4Byte(&Bytes, m.mContent.(int), true)
			toSend = append(toSend, Bytes...)
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
			// Review Remark: Clear naming is needed! don't use a single symbol for a name!!!!
			c := m.mContent.(Message)
			// Review Remark: Clear naming is needed! don't use a single symbol for a name!!!!
			b := c.GenerateMessage()
			Bytes = append(Bytes, b...)
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
			// Review Remark: Clear naming is needed! don't use a single symbol for a name!!!!
			s := m.mContent.(struct {
				code []byte
				name []byte
			})
			Bytes := s.code
			//intTo4Byte(&Bytes, s.code, true)
			toSend = append(toSend, Bytes...)
			toSend = append(toSend, (byte)(len(s.name)))
			toSend = append(toSend, s.name...)
			//spew.Dump(toSend)

		}
	case NewDisconnection:
		{
			//TODO : implement NewDisconnection
		}
	case NewInRoom:
		{
			// Review Remark: Clear naming is needed! don't use a single symbol for a name!!!!
			c := m.mContent.(struct {
				owner []byte
				nLen  []byte
				name  []byte
			})
			toSend = append(toSend, c.owner...)
			toSend = append(toSend, c.nLen...)
			toSend = append(toSend, c.name...)
		}
	case NewOutRoom:
		{
			//TODO : implement NewOutRoom
		}
	case ChatAll:
		{
			// Review Remark: Clear naming is needed! don't use a single symbol for a name!!!!
			c := m.mContent.(struct {
				nLen []byte
				name []byte
				mLen []byte
				mess []byte
			})
			toSend = append(toSend, c.nLen...)
			toSend = append(toSend, c.name...)
			toSend = append(toSend, c.mLen...)
			toSend = append(toSend, c.mess...)
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

	// Review Remark: Independant of switch. Move to outside of switch.
	conn.Write(toSend)
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

func (m *message) GenerateMessage() []byte {
	var b []byte
	switch m.mType {
	//TODO : add al cases
	case ForceTransform:
		{

		}
	case SimpleTransform:
		{

		}
	case CompleteTransform:
		{

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
