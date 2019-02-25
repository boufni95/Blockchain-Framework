package gameserver

import "fmt"

//StdServer : create a standard server
func StdServer() Server {

	sc := stdServerConfig()
	s := NewServer(sc)
	addStdListeners(s)
	return s

}
func stdServerConfig() ServerConfig {
	sc := NewServerConfig(":8080", 1)
	return sc
}

var defListen map[string]chan string

func addStdListeners(s Server) {
	defListen = make(map[string]chan string)
	s.StatusIn("adding listener")
	defListen["connected"] = s.AddListener("connected", connF)

}
func connF(s string) {
	fmt.Println("connected:", s)
}
func discF(s string) {
	fmt.Println("disconnected:", s)
}
