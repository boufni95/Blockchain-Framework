package bcstarter

import (
	conf "GGS/CLinterface/configs"
	bchain "GGS/src/blockchain"
	"GGS/src/core"
	"fmt"
	"net"
)

func Starterb(pathConfig string) error {
	sc, err := conf.ExtractBChainConfig(pathConfig, false)
	if err != nil {
		return err
	}
	for _, v := range sc.SOURCEIPS {
		conn, err := net.Dial("tcp", v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		b := make([]byte, 1)
		b[0] = (byte)(core.StrangeMessage)
		conn.Write(b)
	}
	s := bchain.StdBCServer(sc)
	s.Start()
	return nil
}
