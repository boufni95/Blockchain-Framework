package blockchain

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

var TxsQueue []Transaction

func ProcessTx(Tx Transaction) bool {
	TxsQueue = append(TxsQueue, Tx)
	spew.Dump(TxsQueue)
	return true
}
func Validate() {
	fmt.Println("*****VALIDATE OK******")
}
