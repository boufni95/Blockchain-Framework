package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	Hash         string
	PrevHash     string
}
type Transaction struct {
	nTxsIn  int
	TxsIn   []TxIn
	nTxsOut int
	TxsOut  []TxOut
}
type TxIn struct {
	PrevTx    string
	Index     int
	ScriptSig string
}
type TxOut struct {
	Value        int
	ScriptPubKey string
}

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
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed), nil
}
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
