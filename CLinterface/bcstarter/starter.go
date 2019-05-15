package bcstarter

import (
	conf "GGS/CLinterface/configs"
	bchain "GGS/src/blockchain"
	"GGS/src/core"
	"fmt"
	"io/ioutil"
	"net"
)

var blockchain []bchain.Block

//Starterb : start a blockchian node
func Starterb(pathConfig string) error {
	var sc conf.BChainConfig
	h, err := conf.ExtractBChainConfig(&sc, pathConfig, false)
	if err != nil {
		return err
	}
	fmt.Println("hash", h)
	//------------------------------------------------
	blockchain, err = retriveChain()
	if err != nil {
		return err
	}
	s := bchain.StdBCServer(sc)
	if err := s.SetVar("ConfigHash", h); err != nil {
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
		s.Emit("connected", conn)

	}
	s.Start()
	return nil
}
func retriveChain() ([]bchain.Block, error) {
	var chain []bchain.Block
	_, err := ioutil.ReadFile("blockchain.ggs")
	if err != nil {
		chain = make([]bchain.Block, 1)
		genesis, err := bchain.GenerateGenesisBlock()
		if err != nil {
			return chain, err
		}

		chain[0] = genesis
	}
	return chain, nil

}
func replaceChain(newBlocks []bchain.Block) {
	if len(newBlocks) > len(blockchain) {
		blockchain = newBlocks
	}
}
