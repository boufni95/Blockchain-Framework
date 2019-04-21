package gameserver

import (
	"errors"
	"fmt"
	"net"

	"github.com/davecgh/go-spew/spew"
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
			break
		}

	}

	// Review Remark: If you don't use it- delete it! You have code control versioning system which keeps track of code.
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
	//TODO : implement disconnection
	fmt.Println("disconnected:", s)
}

// Review Remark: Each message type should go to message object.
// Review Remark: With polymorphic behaviour you could avoid those switch statements
// Review Remark: Code smell: shotgun surgery. 3 places for the same responsibility.
// Review Remark: Looks very much like the previous methods for message type.
//FIXME : implement messages for future users
func StdReciveMessage(s Server, conn net.Conn) error {
	mType := make([]byte, 1)

	_, err := conn.Read(mType)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("new mess", mType, "from ", conn.RemoteAddr().String())
	switch (MessageType)(mType[0]) {
	case VoidMessage:
		{
			//TODO : implement VoidMessage
		}
	case InGameMessage:
		{
			// Review Remark: Split to a separate function
			mgType := make([]byte, 1)
			_, err := conn.Read(mgType)
			if err != nil {
				fmt.Println(err)
			}
			switch (MessageType)(mgType[0]) {
			case SimpleTransform:
				{
					code := make([]byte, 4)
					_, err := conn.Read(code)
					if err != nil {
						return err
					}
					pos := make([]byte, 12)
					_, err1 := conn.Read(pos)
					if err1 != nil {
						return err1
					}
					c := struct {
						code []byte
						pos  []byte
					}{code, pos}
					mc := NewMessage(SimpleTransform, c)
					m := NewMessage(InGameMessage, mc)
					s.BroadcastMessageRoom(m, conn.RemoteAddr().String())
				}
			case CompleteTransform:
				{

				}
			}
		}
		/*
			case SimpleTransform:
				{
					code := make([]byte, 4)
					_, err := conn.Read(code)
					if err != nil {
						return err
					}
					pos := make([]byte, 12)
					_, err1 := conn.Read(pos)
					if err1 != nil {
						return err1
					}
					c := struct {
						code []byte
						pos  []byte
					}{code, pos}
					m := NewMessage(SimpleTransform, c)
					s.BroadcastMessageRoom(m, conn.RemoteAddr().String())
					//spew.Dump(code)
					//spew.Dump(pos)
				}
			case CompleteTransform:
				{
					//TODO : implement CompleteTransform
				}
		*/
		// Review Remark: Split to a separate function
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
			b := make([]byte, 4)
			intTo4Byte(&b, code, true)
			wmess := NewMessage(WelcomeMessage, b)
			wmess.Send(s, conn)
			c := struct {
				code []byte
				name []byte
			}{b, name}
			wmess.Mutate(NewConnection, c)
			s.BroadcastMessage(wmess)
		}
	case NewInRoom:
		{
			// Review Remark: Split to a separate function
			owner := make([]byte, 4)
			_, err := conn.Read(owner)
			if err != nil {
				fmt.Println(err)
			}
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
				nLen  []byte
				name  []byte
			}{owner, nLen, name}
			mess := NewMessage(NewInRoom, c)
			s.BroadcastMessageRoom(mess, conn.RemoteAddr().String())
		}
	case ChatAll:
		{
			// Review Remark: Split to a separate function
			num := make([]byte, 1)
			_, err := conn.Read(num)
			if err != nil {
				fmt.Println(err)
			}
			name := make([]byte, num[0])
			// Review Remark: Do you really need err2-4? Why can't you reuse the same error variable?
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
			// Review Remark: Naming!!!
			cmess := NewMessage(ChatAll, c)
			s.BroadcastMessage(cmess)
		}
	case ChatRoom:
		{
			//TODO : implement ChatRoom
		}
	case ChatTo:
		{
			//TODO : implement ChatTo
		}
	default:
		{
			return errors.New("Reciven unknown message type")
		}
	}
	return nil
}


// Review Remark: Delete!!!
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
