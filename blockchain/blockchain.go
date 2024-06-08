package blockchain

import (
	"log"
	"strings"
	"time"
)

const (
	MiningDifficulty = 3
	MiningReward     = 1.0
	MiningSender     = "THE BLOCKCHAIN"
)

type BlockChain struct {
	transactionsPool []transaction
	blocks           []*Block
	recipient        string
}

func NewBlockChain(recipient string) *BlockChain {
	return &BlockChain{
		blocks: []*Block{
			{
				transactions: []transaction{},
				nonce:        0,
				prevHash:     [32]byte{},
				timestamp:    time.Now().UnixMilli(),
			},
		},
		transactionsPool: []transaction{},
		recipient:        recipient,
	}
}

func (bc *BlockChain) AddTransaction(sender string, recipient string, value float32) {
	t := transaction{
		sender:    sender,
		recipient: recipient,
		value:     value,
	}
	bc.transactionsPool = append(bc.transactionsPool, t)
}

func (bc *BlockChain) LastBlock() *Block {
	return bc.blocks[len(bc.blocks)-1]
}

func (bc *BlockChain) ProofOfWork() int {
	nonce := 0
	var copiedTransactionsPool []transaction
	copy(copiedTransactionsPool, bc.transactionsPool)

	for {
		guestBlock := &Block{
			transactions: copiedTransactionsPool,
			nonce:        nonce,
			prevHash:     Hash(bc.LastBlock().ToMarshalJSON()),
			timestamp:    0,
		}
		validateProof := bc.ValidateProof(guestBlock)
		if validateProof {
			break
		}

		nonce++
	}

	return nonce
}

func (bc *BlockChain) ValidateProof(block *Block) bool {
	zeros := strings.Repeat("0", MiningDifficulty)
	hashCode := Hash(block.ToMarshalJSON())

	return strings.HasPrefix(string(hashCode[:MiningDifficulty]), zeros)
}

func (bc *BlockChain) CreateBlock() {
	bc.AddTransaction(MiningSender, bc.recipient, MiningReward)
	lastBlock := bc.LastBlock()
	preHash := Hash(lastBlock.ToMarshalJSON())
	nonce := bc.ProofOfWork()
	block := NewBlock(preHash, nonce, bc.transactionsPool)
	bc.transactionsPool = []transaction{}
	bc.blocks = append(bc.blocks, block)
}

func (bc *BlockChain) Mining() {
	bc.CreateBlock()
}

func (bc *BlockChain) Print() {
	for i, block := range bc.blocks {
		log.Printf("%s Block %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
}
