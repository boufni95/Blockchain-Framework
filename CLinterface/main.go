package main

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	conf "gameserver/CLinterface/configs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	game "gameserver/src/game"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	args := os.Args
	spew.Dump(args)
	reader := bufio.NewReader(os.Stdin)
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
	printIco()
	fmt.Println("1 - Verify config game")
	fmt.Println("2 - Verify config bchain")
	fmt.Println("3 - Verify config wallet")
	fmt.Println("4 - start game server")
	fmt.Println("5 - start blockchain node")
	fmt.Println("6 - start blockchain wallet->create priv/pub key")
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	switch text[0] {
	case '1':
		{
			fmt.Println("selected 1")
			importerg()
		}
	case '2':
		{
			fmt.Println("selected 2")
			importerc()
		}
	case '3':
		{
			fmt.Println("selected 3")
			importerw()
		}
	case '4':
		{
			fmt.Println("selected 4")
			starterg()
		}
	case '5':
		{
			fmt.Println("selected 5")
			starterb()
		}
	case '6':
		{
			fmt.Println("selected 6")
			starterw()
		}
	}

}

func importerg() {

	b, err := ioutil.ReadFile("config.ggs")
	if err != nil {
		fmt.Println("error :(")
		return
	}
	_, err = conf.ExtractGameConfig(b)
	if err != nil {
		fmt.Println(err)
	}
}
func importerc() {

	b, err := ioutil.ReadFile("../Templates/Examples/BChainConfig.ggs")
	if err != nil {
		fmt.Println("error :(")
		return
	}
	var config conf.BChainConfig
	config, err = conf.ExtractBChainConfig(b)
	if err != nil {
		fmt.Println(err)
	}
	spew.Dump(config)
}
func importerw() {

	fmt.Println("not implemented")
}
func starterg() {
	s := game.StdServer()
	s.Start()
}
func starterb() {
	fmt.Println("not implemented")
}
func starterw() {
	fmt.Println("not implemented")
	fmt.Println("creating a key pair...")
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	bPriv, err := json.MarshalIndent(privateKey, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(bPriv)
	f1, err := os.OpenFile("privkey.ggs", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	if err := f1.Close(); err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("privkey.ggs", bPriv, 0644)

	bPub, err := json.MarshalIndent(privateKey.PublicKey, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(bPub)
	f2, err := os.OpenFile("pubkey.ggs", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()
	if err := f2.Close(); err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("pubkey.ggs", bPub, 0644)

}
func printIco() {
	fmt.Println("***************************************")
	fmt.Println("****  _______   _______   _______  ****")
	fmt.Println("**** |         |         |         ****")
	fmt.Println("**** |         |  *-*    |_______  ****")
	fmt.Println("**** |      _  |      _          | ****")
	fmt.Println("**** |______|  |______|          | ****")
	fmt.Println("**** ____________________________| ****")
	fmt.Println("****                               ****")
	fmt.Println("***************************************")
}

/*
func generate() {
	fmt.Println("selected 1")
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.MarshalIndent(group, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	f, err := os.OpenFile("config.ggs", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("config.ggs", b, 0644)
}*/
