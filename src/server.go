package gameserver

import (
	"fmt"
	"net"
)

//------------------------------------------------------------------
//-------------------TYPES------------------------------------------
//------------------------------------------------------------------

//EventFunc : is the type of the function passed to the events handler
type EventFunc func(string)

//------------------------------------------------------------------
//-------------------INTERFACES-------------------------------------
//------------------------------------------------------------------

//Server : the server interface
type Server interface {
	Start() error
	Close()
	AddListener(string, EventFunc) chan string
	RemoveListener(string, chan string)
	Emit(string, string)
	Status() string
	StatusIn(string)
}

//ServerConfig : the interface of the server config
type ServerConfig interface {
	Port() string
}

//------------------------------------------------------------------
//-------------------FUNCTIONS--------------------------------------
//------------------------------------------------------------------

//NewServerConfig : create a new server config
func NewServerConfig(port string, maxRooms int) ServerConfig {
	sc := serverConfig{port, maxRooms}
	return &sc
}

//NewServer : create a new server with given server config
func NewServer(sc ServerConfig) Server {
	var s server
	s.port = sc.Port()
	s.ln = nil
	s.listeners = make(map[string][]chan string)
	s.status = "created"
	return &s

}

//------------------------------------------------------------------
//-------------------STRUCTS----------------------------------------
//------------------------------------------------------------------

type server struct {
	ln        net.Listener
	port      string
	listeners map[string][]chan string
	status    string
}

func (s *server) Start() error {
	var err error
	s.ln, err = net.Listen("tcp", s.port)
	if err != nil {
		return err
	}

	for {
		conn, err := s.ln.Accept()
		if err != nil {
			fmt.Println(err)
		} else {
			s.Emit("connected", conn.RemoteAddr().String())
		}

	}

}
func (s *server) Close() {

}
func (s *server) AddListener(e string, ef EventFunc) chan string {
	ch := make(chan string)
	//s.status = "addingList"
	_, ok := s.listeners[e]
	if ok {
		s.listeners[e] = append(s.listeners[e], ch)
	} else {
		s.listeners[e] = []chan string{ch}
	}
	go func(ch chan string, ef EventFunc) {
		for {
			f := <-ch
			ef(f)
		}
	}(ch, ef)
	return ch
}
func (s *server) RemoveListener(e string, ch chan string) {
	_, ok := s.listeners[e]
	if ok {
		for i := range s.listeners[e] {
			if s.listeners[e][i] == ch {
				s.listeners[e] = append(s.listeners[e][:i], s.listeners[e][i+1:]...)
				break
			}
		}
	}
}
func (s *server) Emit(e string, st string) {

	_, ok := s.listeners[e]
	if ok {
		for _, handler := range s.listeners[e] {
			go func(handler chan string) {
				handler <- st
			}(handler)
		}
	}
}
func (s *server) Status() string {
	return s.status
}
func (s *server) StatusIn(st string) {
	s.status = st
}

type serverConfig struct {
	port     string
	maxRooms int
}

//Port : the port to use for the server
func (sc *serverConfig) Port() string {
	return sc.port
}
