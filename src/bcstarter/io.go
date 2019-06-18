package bcstarter

import (
	"Blockchain-Framework/src/blockchain"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
)

//CreateDirIfNotExist creted a dir if it doesent exist
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

//SaveToFile save a slice of bytes to a file, if non existant, create it
func SaveToFile(fileName string, b []byte) {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(fileName, b, 0644)
}
func HttpGetChain(w http.ResponseWriter, req *http.Request) {
	BJson, err := json.MarshalIndent(blockchainVar.Get(), "", "   ")
	if err != nil {
		return
	}
	io.WriteString(w, string(BJson))
}
func HttpPostTx(w http.ResponseWriter, req *http.Request) {
	// Set CORS headers for the preflight request
	if req.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	b := make([]byte, 10000)
	n, err := req.Body.Read(b)
	if err != nil {
		fmt.Println(err)

	}
	b = b[0:n]
	var Tx blockchain.Transaction
	err = json.Unmarshal(b, &Tx)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	spew.Dump(Tx)
	err = checkTx(Tx)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, "Tx structure ok, will now check consistency")
	blockchain.ProcessTx(Tx)
}

func checkTx(Tx blockchain.Transaction) error {
	res := Tx.Hash != "" &&
		Tx.NTxsIn > 0 &&
		Tx.NTxsOut > 0
	if res == false {
		return errors.New("Missing Hash or TxsIn or TxsOut")
	}
	err := checkTxsIn(Tx.NTxsIn, Tx.TxsIn)
	if err != nil {
		return err
	}
	err = checkTxsOut(Tx.NTxsOut, Tx.TxsOut)
	if err != nil {
		return err
	}
	return nil
}
func checkTxsIn(NTxsIn int, TxsIn []blockchain.TxIn) error {
	if NTxsIn != len(TxsIn) {
		return errors.New("inconsistent TxIn")
	}
	for _, v := range TxsIn {
		if v.PrevTx == "" {
			return errors.New("invalid prev Tx")
		}
	}
	return nil

}
func checkTxsOut(NTxsOut int, TxsOut []blockchain.TxOut) error {
	if NTxsOut != len(TxsOut) {
		return errors.New("inconsistent TxOut")
	}
	for _, v := range TxsOut {
		if v.Value == 0 {
			return errors.New("output can't be 0")
		}
	}
	return nil
}
