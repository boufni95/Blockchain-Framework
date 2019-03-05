package gameserver

import (
	"errors"
	"fmt"
	"net"
)

//StdServer : create a standard server
func StdServer() Server {

	sc := stdServerConfig()
	s := NewServer(sc)
	StdAddListeners(s)
	return s

}
func stdServerConfig() ServerConfig {
	sc := NewServerConfig(":8080", 1)
	return sc
}

var defListen map[string]chan net.Conn

//StdAddListeners :  add the standard listeners to the server
//NEW CONNECTION
//NEW DOSCONNECTION
func StdAddListeners(s Server) {
	defListen = make(map[string]chan net.Conn)
	s.StatusIn("adding listener")
	defListen["connected"] = s.AddListener("connected", stdOnConnected)

}
func stdOnConnected(s Server, conn net.Conn) {
	for {
		if err := StdReciveMessage(s, conn); err != nil {
			fmt.Println(err)
			continue
		}

	}
	/*num := make([]byte, 1)
	_, err := conn.Read(num)
	if err != nil {
		fmt.Println(err)
	}
	name := make([]byte, num[0])
	_, err2 := conn.Read(name)
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println("connected:", string(name))


	c := struct {
		code int
		name []byte
	}{code, name}
	wmess.Mutate(NewConnection, c)
	s.BroadcastMessage(wmess)
	for {
		_, err := conn.Read(num)
		if err != nil {
			fmt.Println(err)
			break
		}
		if num[0] == (byte)(SimpleTransform) {
			posB := make([]byte, 12)
			_, err := conn.Read(posB)
			if err != nil {
				fmt.Println(err)
				break
			}
			spew.Dump(posB)
		}
	}*/
}
func discF(s string) {
	fmt.Println("disconnected:", s)
}
func StdReciveMessage(s Server, conn net.Conn) error {
	mType := make([]byte, 1)
	_, err := conn.Read(mType)
	if err != nil {
		fmt.Println(err)
	}
	switch (MessageType)(mType[0]) {
	case VoidMessage:
		{

		}
	case SimpleTransform:
		{

		}
	case CompleteTransform:
		{

		}
	case NameString:
		{
			num := make([]byte, 1)
			_, err := conn.Read(num)
			if err != nil {
				fmt.Println(err)
			}
			name := make([]byte, num[0])
			_, err2 := conn.Read(name)
			if err2 != nil {
				fmt.Println(err)
			}
			fmt.Println("connected:", string(name))

			code := s.AddPlayer(string(name), conn)
			wmess := NewMessage(WelcomeMessage, code)
			wmess.Send(s, conn)
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
	default:
		{
			return errors.New("Reciven unknown message type")
		}
	}
	return nil
}

/*
	var st = []byte("012345")
	buff := make([]byte, 0)
	buff = append(buff, (byte)(6))
	buff = append(buff, st...)
	for i := 0; i < 10000; i++ {
		go func(buff []byte) {
			fmt.Println("writing.. ", buff)
			_, err3 := conn.Write(buff)
			if err3 != nil {
				fmt.Println(err3)
			}
		}(buff)
	}
*/
