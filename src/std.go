package gameserver

import (
	"fmt"
	"net"

	"github.com/davecgh/go-spew/spew"
)

//StdServer : create a standard server
func StdServer() Server {

	sc := stdServerConfig()
	s := NewServer(sc)
	AddStdListeners(s)
	return s

}
func stdServerConfig() ServerConfig {
	sc := NewServerConfig(":8080", 1)
	return sc
}

var defListen map[string]chan net.Conn

//AddStdListeners :  add the standard listeners to the server
//NEW CONNECTION
//NEW DOSCONNECTION
func AddStdListeners(s Server) {
	defListen = make(map[string]chan net.Conn)
	s.StatusIn("adding listener")
	defListen["connected"] = s.AddListener("connected", onConnected)

}
func onConnected(s Server, conn net.Conn) {
	//s := conn.RemoteAddr().String()
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
	//type cont struct {
	//	code int
	//	name []byte
	//}
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
	}
}
func discF(s string) {
	fmt.Println("disconnected:", s)
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
