package main

import (
	"github.com/foxlovecat1989/blockchain/blockchain"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

const name = "ed"

func main() {
	chain := blockchain.NewBlockChain(name)
	chain.AddTransaction("A", "B", 100)
	chain.Mining()
	chain.AddTransaction("X", "Y", 50)
	chain.AddTransaction("Y", "Z", 25)
	chain.Mining()
	chain.AddTransaction("B", "C", 200)
	chain.AddTransaction("C", "D", 100)
	chain.AddTransaction("D", "E", 50)
	chain.Mining()
	chain.Print()
}
