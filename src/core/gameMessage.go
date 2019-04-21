package core

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
	//1byte -> player type
	//12byte -> Vector3int position
	SimpleTransform MessageType = 5 // s <-> c

	//CompleteTransform : message containing a coplete transform
	//STRUCTURE:
	//1byte -> message type
	//
	CompleteTransform MessageType = 7 // s <-> c

	//NewObject : message to declare creation of new object
	//STRUCTURE:
	//1byte -> message type
	//4byte -> owner
	//1byte -> owner type
	//1byte -> code type
	//1byte -> code obj
	//12byte -> position
	NewObject MessageType = 10

	//UpdateObject : message to update object
	//STRUCTURE:
	//1byte -> message type
	//4byte -> owner
	//1byte -> owner type
	//1byte -> code type
	//1byte -> code obj
	//12byte -> position
	UpdateObject MessageType = 13

	//DestroyObject :	destroy object in game
	//STRUCTURE:
	//1byte -> message type
	//4byte -> owner
	//1byte -> owner type
	//1byte -> code type
	//1byte -> code obj
	DestroyObject MessageType = 15

	//NewCollision : message to inform the server of a new collision
	//STRUCTURE:
	//1byte -> message type
	//4byte -> code owner
	//4byte -> player code
	//1byte -> bullet code type
	//1byte -> bullet code obj
	NewCollision MessageType = 20
)
