package bcstarter

import (
	conf "Blockchain-Framework/configs"
	bchain "Blockchain-Framework/src/blockchain"
	"Blockchain-Framework/src/core"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var blockchainVar []bchain.Block

//Starterb : start a blockchian node
func Starterb(pathConfig string, chainDir string) error {
	var sc conf.BChainConfig
	h, err := conf.ExtractBChainConfig(&sc, pathConfig, false)
	if err != nil {
		return err
	}
	fmt.Println("hash", h)
	//------------------------------------------------
	blockchainVar, err = retriveChain(chainDir)
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
	go serveHttp()
	s.Start()
	return nil
}
func retriveChain(dir string) ([]bchain.Block, error) {
	CreateDirIfNotExist(dir)
	var chain []bchain.Block
	_, err := ioutil.ReadFile(dir + "/blocks-0.ggs")
	if err != nil {
		chain = make([]bchain.Block, 1)
		genesis, err := bchain.GenerateGenesisBlock()
		b, err := json.MarshalIndent(genesis, " ", "    ")
		SaveToFile(dir+"/blocks-0.ggs", b)
		if err != nil {
			return chain, err
		}

		chain[0] = genesis
	}
	return chain, nil

}
func replaceChain(newBlocks []bchain.Block) {
	if len(newBlocks) > len(blockchainVar) {
		blockchainVar = newBlocks
	}
}
func serveHttp() {

	http.HandleFunc("/set-tx", HttpPostTx)
	http.HandleFunc("/", HttpGetChain)
	log.Fatal(http.ListenAndServe(":80", nil))
}
