package main

import (
	"fmt"
	"time"
)

type block struct {
	transactions []string
	nonce        string
	prevHash     string
	timestamp    int64
}

func NewBlock(prevHash string, nonce string) *block {
	return &block{
		transactions: []string{},
		nonce:        nonce,
		prevHash:     prevHash,
		timestamp:    time.Now().Unix(),
	}
}

func (b *block) Print(block block) {
	fmt.Printf("Block: %v\n", block)
	fmt.Printf("transactions: %v\n", block.transactions)
	fmt.Printf("nonce: %v\n", block.nonce)
	fmt.Printf("prevHash: %v\n", block.prevHash)
	fmt.Printf("timestamp: %v\n", block.timestamp)
}

func main() {
	block := NewBlock("123", "456")
	block.Print(*block)
}
