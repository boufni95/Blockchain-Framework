package bcstarter

import (
	"fmt"
	conf "gameserver/CLinterface/configs"
	bchain "gameserver/src/blockchain"
	"net"
)

func Starterb(pathConfig string) error {
	fmt.Println("not implemented")
	sc, err := conf.ExtractBChainConfig(pathConfig, false)
	if err != nil {
		return err
	}
	for _, v := range sc.SOURCEIPS {
		_, err := net.Dial("tcp", v)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
	s := bchain.StdBCServer(sc)
	s.Start()
	return nil
}
