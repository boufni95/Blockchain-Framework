package main

import (
	conf "GGS/CLinterface/configs"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	game "GGS/src/game"

	bcstart "GGS/CLinterface/bcstarter"

	spew "github.com/davecgh/go-spew/spew"
)

func main() {
	args := os.Args
	/*spew.Dump(args)
	reader := bufio.NewReader(os.Stdin)
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
	printIco()
	fmt.Println("-vcg - Verify config game")
	fmt.Println("-vcb - Verify config bchain")
	fmt.Println("-vcw - Verify config wallet")
	fmt.Println("-sg - start game server")
	fmt.Println("-sb - start blockchain node")
	fmt.Println("-sw - start blockchain wallet->create priv/pub key")
	fmt.Print("-> ")*/
	//text, _ := reader.ReadString('\n')
	fmt.Println("welcome to ggs...")
	if len(args) < 2 {
		fmt.Println("missing arguments")
		return
	}
	switch args[1] {
	case "-vcg":
		{
			fmt.Println("Verify config game")
			importerg()
		}
	case "-vcb":
		{
			fmt.Println("Verify config bchain")
			if len(args) < 3 {
				fmt.Println("missing arguments")
				return
			}
			path := args[2]
			config, err := conf.ExtractBChainConfig(path, true)
			if err != nil {
				fmt.Println(err)
			}
			spew.Dump(config)
		}
	case "-vcw":
		{
			fmt.Println("Verify config wallet")
			importerw()
		}
	case "-sg":
		{
			fmt.Println("start game server")
			starterg()
		}
	case "-sb":
		{
			fmt.Println("start blockchain node")
			err := bcstart.Starterb("../Templates/Examples/BChainConfig.ggs")
			if err != nil {
				fmt.Println(err)
			}
		}
	case "-sw":
		{
			fmt.Println("start blockchain wallet->create priv/pub key")
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
	path := "../Templates/Examples/BChainConfig.ggs"
	config, err := conf.ExtractBChainConfig(path, true)
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
func PrintIco() {
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
