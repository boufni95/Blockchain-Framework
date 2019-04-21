package blockchain

import (
	"errors"
	"fmt"
	"net"

	core "gameserver/src/core"

	"github.com/davecgh/go-spew/spew"
)

//StdServer : create a standard server
func StdServer() core.Server {
	/*sMods := []GameMode{{"stdMode"}}
	sObjs := []GMObject{{"meteor"}}

	sConfig := struct {
		maxRooms int
		clans    bool
		gamemods []GameMode
		objects  []GMObject
	}{1, true, sMods, sObjs}

	testObjs := []PObject{{"testObj", 1, 1, 1}}
	pTypes := []PlayerType{{"testType", 1, 1, testObjs}}

	pConfig := struct {
		clan  bool
		types []PlayerType
	}{true, pTypes}*/

	sc := stdServerConfig()
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
			//TODO : implement VoidMessage
		}
	case core.NameString:
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

			code := s.AddConnection(string(name), conn)
			b := make([]byte, 4)
			core.IntTo4Byte(&b, code, true)
			wmess := core.NewMessage(core.WelcomeMessage, b)
			wmess.Send(s, conn)
			c := struct {
				code []byte
				name []byte
			}{b, name}
			wmess.Mutate(core.NewConnection, c)
			s.BroadcastMessage(wmess)
		}
	case core.NewInRoom:
		{
			owner := make([]byte, 4)
			_, err := conn.Read(owner)
			if err != nil {
				fmt.Println(err)
			}
			pType := make([]byte, 1)
			_, err4 := conn.Read(pType)
			if err4 != nil {
				fmt.Println(err4)
			}
			//TODO : read player type here!!
			nLen := make([]byte, 1)
			_, err2 := conn.Read(nLen)
			if err2 != nil {
				fmt.Println(err)
			}
			name := make([]byte, nLen[0])
			_, err3 := conn.Read(name)
			if err3 != nil {
				fmt.Println(err3)
			}
			s.AssignRoom("", conn.RemoteAddr().String())
			c := struct {
				owner []byte
				pType byte
				nLen  []byte
				name  []byte
			}{owner, pType[0], nLen, name}
			mess := core.NewMessage(core.NewInRoom, c)
			s.BroadcastMessageRoom(mess, conn.RemoteAddr().String())
		}
	case core.ChatAll:
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
			n := make([]byte, 1)
			_, err3 := conn.Read(n)
			if err3 != nil {
				fmt.Println(err)
			}
			mess := make([]byte, n[0])
			_, err4 := conn.Read(mess)
			if err4 != nil {
				fmt.Println(err)
			}
			spew.Dump(name)
			spew.Dump(mess)
			c := struct {
				nLen []byte
				name []byte
				mLen []byte
				mess []byte
			}{num, name, n, mess}
			cmess := core.NewMessage(core.ChatAll, c)
			s.BroadcastMessage(cmess)
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
