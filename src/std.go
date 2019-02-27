package gameserver

import (
	"fmt"
	"net"
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
	defListen["connected"] = s.AddListener("connected", connF)

}
func connF(s Server, conn net.Conn) {
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

	s.AddPlayer(string(name), conn)
	NewMessage(VoidMessage, nil)
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
}
func discF(s string) {
	fmt.Println("disconnected:", s)
}
