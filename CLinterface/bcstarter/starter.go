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
		iam := core.NewMessage(core.IAmNode, nil)
		bcm := core.NewMessage(core.BChainMessage, iam)
		bcm.Send(nil, conn)
	}
	s := bchain.StdBCServer(sc)
	s.Start()
	return nil
}
