package core

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

	//BChainMessage : defines a type of messages that are send diring game
	//STRUCTURE:
	//1byte -> message type
	//1byte -> blockchain message type
	//Nbyte -> message
	BChainMessage MessageType = 7

	//NameString : message containing the name of the player
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
	//1byte : player type
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
)
