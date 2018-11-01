package main

import (
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/dianabejan/blockchain/block"
	"github.com/dianabejan/blockchain/server"
	"github.com/joho/godotenv"
)

func main() {
	var blockchain = []block.Block{}
	var server = server.NewServer(blockchain)
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := block.Block{0, t.String(), 0, 1, "", "", ""}
		spew.Dump(genesisBlock)
		blockAux := append(server.Blockchain, genesisBlock)
		server.ReplaceChain(blockAux)
	}()
	log.Fatal(server.Run())
}
