package blockchain

import (
	"encoding/json"
	"time"
)

//Block : struct of a blockchan block
type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	Hash         string
	PrevHash     string
}

//Transaction : struct of a blockchain transaction
type Transaction struct {
	nTxsIn  int
	TxsIn   []TxIn
	nTxsOut int
	TxsOut  []TxOut
}

//TxIn : struct for bc transaction input
type TxIn struct {
	PrevTx    string
	Index     int
	ScriptSig string
}

//TxOut : struct for bc transaction output
type TxOut struct {
	Value        int
	ScriptPubKey string
}

//CalculateHash : calcumate has of the block
func (b *Block) CalculateHash() (string, error) {
	toHash := struct {
		Index        int
		Timestamp    string
		Transactions []Transaction
		PrevHash     string
	}{b.Index, b.Timestamp, b.Transactions, b.PrevHash}
	record, err := json.Marshal(toHash)
	if err != nil {
		return "", err
	}
	h := HashSha256(record)
	return h, nil
}

//GenerateBlock : generate a new block from an old block and some txs
func GenerateBlock(oldBlock Block, txs []Transaction) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Transactions = txs
	newBlock.PrevHash = oldBlock.Hash
	hash, err := newBlock.CalculateHash()
	if err != nil {
		return newBlock, err
	}
	newBlock.Hash = hash

	return newBlock, nil
}

//GenerateGenesisBlock : generate a genesis block
func GenerateGenesisBlock() (Block, error) {

	var newBlock Block

	t := time.Now()
	txs := make([]Transaction, 0)
	newBlock.Index = 0
	newBlock.Timestamp = t.String()
	newBlock.Transactions = txs
	newBlock.PrevHash = ""
	hash, err := newBlock.CalculateHash()
	if err != nil {
		return newBlock, err
	}
	newBlock.Hash = hash

	return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	h, err := newBlock.CalculateHash()
	if err != nil {
		return false
	}
	if h != newBlock.Hash {
		return false
	}

	return true
}
