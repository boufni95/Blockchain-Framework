package bcstarter

import (
	conf "GGS/CLinterface/configs"
	bchain "GGS/src/blockchain"
	"GGS/src/core"
	"fmt"
	"net"
	"time"

	"github.com/davecgh/go-spew/spew"
)

//Starterb : start a blockchian node
func Starterb(pathConfig string) error {
	var sc conf.BChainConfig
	h, err := conf.ExtractBChainConfig(&sc, pathConfig, false)
	if err != nil {
		return err
	}
	fmt.Println("hash", h)
	s := bchain.StdBCServer(sc)
	if err := s.SetVar("ConfigHash", h); err != nil {
		return err
	}
	t := time.Now()
	txs := make([]bchain.Transaction, 0)
	genesis := bchain.Block{0, t.String(), txs, "", ""}
	spew.Dump(genesis)
	for _, v := range sc.SOURCEIPS {
		conn, err := net.Dial("tcp", v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		iam := core.NewMessage(core.IAmNode, nil)
		bcm := core.NewMessage(core.BChainMessage, iam)
		bcm.Send(nil, conn)
		s.Emit("connected", conn)

	}
	s.Start()
	return nil
}
