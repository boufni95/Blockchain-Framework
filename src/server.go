package gameserver

import (
	"errors"
	"fmt"
	"net"
)

//------------------------------------------------------------------
//-------------------TYPES------------------------------------------
//------------------------------------------------------------------

//EventFunc : is the type of the function passed to the events handler
type EventFunc func(Server, net.Conn)

//------------------------------------------------------------------
//-------------------INTERFACES-------------------------------------
//------------------------------------------------------------------

//Server : the server interface
type Server interface {
	Start() error
	Close()
	AddListener(string, EventFunc) chan net.Conn
	RemoveListener(string, chan net.Conn)
	Emit(string, net.Conn)
	Status() string
	StatusIn(string)
	AddRoom(string, int) error
	RemoveRoom(string)
	AddPlayer(string, net.Conn) int
	RemovePlayer(string)
	BroadcastMessage(Message) error
	BroadcastMessageRoom(Message, string) error
	SendMessageToConn(Message, net.Conn) error
	SendMessageToAddr(Message, string) error
	AssignRoom(string, string)
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
	s.listeners = make(map[string][]chan net.Conn)
	s.players = make(map[string]Player)
	s.rooms = make(map[string]Room)
	s.status = "created"
	s.AddRoom(randStringBytes(8), 100)
	return &s

}

//------------------------------------------------------------------
//-------------------STRUCTS----------------------------------------
//------------------------------------------------------------------

type server struct {
	ln          net.Listener
	port        string
	listeners   map[string][]chan net.Conn
	status      string
	players     map[string]Player
	playersIndx []string
	rooms       map[string]Room
}

//-----------------------------------------------
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
			s.Emit("connected", conn)
		}

	}

}
func (s *server) Close() {
	//TODO : implement Close
}
func (s *server) AddListener(e string, ef EventFunc) chan net.Conn {
	ch := make(chan net.Conn)
	//s.status = "addingList"
	_, ok := s.listeners[e]
	if ok {
		s.listeners[e] = append(s.listeners[e], ch)
	} else {
		s.listeners[e] = []chan net.Conn{ch}
	}
	go func(s Server, ch chan net.Conn, ef EventFunc) {
		for {
			//fmt.Println("gonna fire it")
			f := <-ch
			go ef(s, f)
		}
	}(s, ch, ef)
	return ch
}
func (s *server) RemoveListener(e string, ch chan net.Conn) {
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
func (s *server) Emit(e string, conn net.Conn) {

	_, ok := s.listeners[e]
	if ok {
		for _, handler := range s.listeners[e] {
			go func(handler chan net.Conn) {
				handler <- conn
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
func (s *server) AddRoom(key string, maxP int) error {

	if _, ok := s.rooms[key]; ok == false {
		r := NewRoom(s, key, maxP)
		s.rooms[key] = r
		return nil
	}
	return errors.New("the key room already exists")
}
func (s *server) RemoveRoom(key string) {
	//TODO : implement RemoveRoom
}
func (s *server) AddPlayer(st string, conn net.Conn) int {
	var p player
	p.name = st
	p.conn = conn
	s.players[conn.RemoteAddr().String()] = &p
	found := false
	indx := -1
	for i, v := range s.playersIndx {
		indx = i
		if v == "" {
			v = conn.RemoteAddr().String()
			found = true
			break

		}
	}
	if found == false {
		indx++
		s.playersIndx = append(s.playersIndx, conn.RemoteAddr().String())
	}
	return indx
}
func (s *server) RemovePlayer(st string) {
	delete(s.players, st)
	for i := range s.playersIndx {
		if s.playersIndx[i] == st {
			s.playersIndx[i] = "removed"
			break
		}
	}
}
func (s *server) BroadcastMessage(m Message) error {
	for key := range s.players {
		s.players[key].SendMessage(m)
	}
	return nil
}
func (s *server) BroadcastMessageRoom(m Message, key string) error {
	s.players[key].BroadcastRoom(m)
	return nil
}
func (s *server) SendMessageToConn(m Message, conn net.Conn) error {
	//TODO : implement SendMessageToConn
	return nil
}
func (s *server) SendMessageToAddr(m Message, ip string) error {
	s.players[ip].SendMessage(m)
	return nil
}
func (s *server) AssignRoom(keyR string, keyP string) {
	if keyR == "" {
		//spew.Dump(s.rooms)
		for i := range s.rooms {
			if s.rooms[i].FreeSpots() > 0 {
				s.rooms[i].AddPlayer(keyP)
				s.players[keyP].SetRoom(s.rooms[i])
				break
			}
		}
		//spew.Dump(s.rooms)
	} else {
		//assign specific room
		//TODO : implement specific room assignement
	}
}

//==========================================================
type serverConfig struct {
	port     string
	maxRooms int
}

//----------------------------------------------------------
//Port : the port to use for the server
func (sc *serverConfig) Port() string {
	return sc.port
}
