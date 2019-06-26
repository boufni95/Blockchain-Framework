package blockchain

import (
	"Blockchain-Framework/src/core"
	"fmt"
	"net"
	"sync"

	"github.com/davecgh/go-spew/spew"
)

var TxsQueue struct {
	sync.RWMutex
	Txs []Transaction
}

func ProcessTx(Tx Transaction) bool {
	TxsQueue.Lock()
	TxsQueue.Txs = append(TxsQueue.Txs, Tx)
	TxsQueue.Unlock()
	TxsQueue.RLock()
	spew.Dump(TxsQueue)
	TxsQueue.RUnlock()
	return true
}
func Validate(s core.Server) {
	fmt.Println("*****VALIDATE OK******")
	s.AddListener("ready", ReadyToValidate)
}
func ReadyToValidate(s core.Server, conn net.Conn) {
	fmt.Println("** ready to validate **")
	for {
		TxsQueue.RLock()
		if len(TxsQueue.Txs) == 0{
			TxsQueue.RUnlock()
			break
		}
		//TODO --
	}

}
