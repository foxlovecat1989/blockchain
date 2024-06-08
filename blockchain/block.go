package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"time"
)

type Block struct {
	transactions []transaction
	nonce        int
	prevHash     [32]byte
	timestamp    int64
}

func NewBlock(prevHash [32]byte, nonce int, txs []transaction) *Block {
	return &Block{
		transactions: txs,
		nonce:        nonce,
		prevHash:     prevHash,
		timestamp:    time.Now().UnixMilli(),
	}
}

func (b *Block) Print() {
	log.Printf("transactionsPool:\t%v\n", b.transactions)
	log.Printf("nonce:\t\t\t%v\n", b.nonce)
	log.Printf("prevHash:\t\t%x\n", b.prevHash)
	log.Printf("timestamp:\t\t%v\n", b.timestamp)
}

func (b *Block) ToMarshalJSON() []byte {
	bs, _ := json.Marshal(struct {
		Transactions []transaction `json:"transactions"`
		Nonce        int           `json:"nonce"`
		PrevHash     [32]byte      `json:"prevHash"`
		Timestamp    int64         `json:"timestamp"`
	}{
		Transactions: b.transactions,
		Nonce:        b.nonce,
		PrevHash:     b.prevHash,
		Timestamp:    b.timestamp,
	})

	return bs
}

func Hash(bytes []byte) [32]byte {
	return sha256.Sum256(bytes)
}
