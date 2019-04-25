package blockchain

import (
	"errors"
	"fmt"
	"net"

	core "GGS/src/core"
)

//StdBCServer : create a standard server
func StdBCServer(sc core.ServerConfig) core.Server {

	s := core.NewServer(sc)
	StdAddListeners(s)
	return s

}
func stdServerConfig() core.ServerConfig {
	sc := core.NewServerConfig(":8080", 1)
	return sc
}

var defListen map[string]chan net.Conn

//StdAddListeners :  add the standard listeners to the server
//NEW CONNECTION
//NEW DOSCONNECTION
func StdAddListeners(s core.Server) {
	defListen = make(map[string]chan net.Conn)
	s.StatusIn("adding listener")
	defListen["connected"] = s.AddListener("connected", stdOnConnected)

}
func stdOnConnected(s core.Server, conn net.Conn) {
	for {
		if err := StdReciveMessage(s, conn); err != nil {
			fmt.Println(err)
			break
		}

	}
}
func discF(s string) {
	//TODO : implement disconnection
	fmt.Println("disconnected:", s)
}

//FIXME : implement messages for future users

//StdReciveMessage : Standard function to recive messages
func StdReciveMessage(s core.Server, conn net.Conn) error {
	mType := make([]byte, 1)

	_, err := conn.Read(mType)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("new mess", mType, "from ", conn.RemoteAddr().String())
	switch (core.MessageType)(mType[0]) {
	case core.VoidMessage:
		{
		}
	case core.BChainMessage:
		{
			fmt.Println("recived bc message")
			HandleBCmessage(s, conn)
		}
	case core.NameString:
		{

		}
	case core.NewInRoom:
		{
		}
	case core.ChatAll:
		{
		}
	case core.ChatRoom:
		{
			//TODO : implement ChatRoom
		}
	case core.ChatTo:
		{
			//TODO : implement ChatTo
		}
	default:
		{
			return errors.New("Recived unknown message type")
		}
	}
	return nil
}
func HandleBCmessage(s core.Server, conn net.Conn) {
	mType := make([]byte, 1)

	_, err := conn.Read(mType)
	if err != nil {
		fmt.Println(err)
	}
	switch (core.MessageType)(mType[0]) {
	case core.IAmNode:
		{
			fmt.Println("recived I am node")
		}
	}
}
