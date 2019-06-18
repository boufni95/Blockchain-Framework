package bcstarter

import (
	"Blockchain-Framework/src/blockchain"
	"Blockchain-Framework/src/configs"
	"Blockchain-Framework/src/core"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var blockchainVar blockchain.Chain

//Starterb : start a blockchian node
func Starterb(pathConfig string, chainDir string) error {
	var sc configs.BChainConfig
	h, err := configs.ExtractBChainConfig(&sc, pathConfig, false)
	if err != nil {
		return err
	}
	fmt.Println("hash", h)
	//------------------------------------------------
	blockchainVar, err = retriveChain(chainDir)
	if err != nil {
		return err
	}
	s := blockchain.StdBCServer(sc)
	if err := s.SetVar("ConfigHash", h); err != nil {
		return err
	}

	for _, v := range sc.SOURCEIPS {
		conn, err := net.Dial("tcp", v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		/*iam := core.NewMessage(core.IAmNode, nil)
		bcm := core.NewMessage(core.BChainMessage, iam)
		bcm.Send(nil, conn)*/
		confMex := core.NewMessage(core.Config, ([]byte)(h))
		bcm := core.NewMessage(core.BChainMessage, confMex)
		//spew.Dump(confMex)
		bcm.Send(nil, conn)
		s.Emit("connected", conn)

	}
	go serveHttp()
	if sc.VALIDATE {
		go blockchain.Validate()
	}
	s.Start()
	return nil
}
func retriveChain(dir string) (blockchain.Chain, error) {
	CreateDirIfNotExist(dir)
	var bChain []blockchain.Block
	_, err := ioutil.ReadFile(dir + "/blocks-0.ggs")
	if err != nil {
		bChain = make([]blockchain.Block, 1)
		genesis, err := blockchain.GenerateGenesisBlock()
		b, err := json.MarshalIndent(genesis, " ", "    ")
		SaveToFile(dir+"/blocks-0.ggs", b)
		if err != nil {
			return nil, err
		}

		bChain[0] = genesis
	}
	chain := blockchain.NewChain(bChain, dir)
	return chain, nil

}

func serveHttp() {

	http.HandleFunc("/set-tx", HttpPostTx)
	http.HandleFunc("/", HttpGetChain)
	log.Fatal(http.ListenAndServe(":80", nil))
}
