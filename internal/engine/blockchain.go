package engine

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/alwindoss/fragmos"
)

func NewBlockchain(addr string) *Blockchain {
	bc := new(Blockchain)
	bc.blockchainAddress = addr
	bc.createGenesisBlock(0)
	return bc
}

type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
}

func (bc *Blockchain) CreateBlock(nonce int, prevHash [32]byte) *Block {
	blk := NewBlock(nonce, prevHash, bc.transactionPool)
	bc.chain = append(bc.chain, blk)
	bc.transactionPool = []*Transaction{}
	return blk
}

func (bc Blockchain) LastBlock() *Block {
	blk := bc.chain[len(bc.chain)-1]
	return blk
}

func (bc *Blockchain) createGenesisBlock(nonce int) *Block {
	var prevHash [32]byte
	txns := []*Transaction{}
	blk := NewBlock(nonce, prevHash, txns)
	bc.chain = append(bc.chain, blk)
	return blk
}

func (bc Blockchain) Print() {
	for i, b := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("#", 25), i, strings.Repeat("#", 25))
		b.Print()
	}

	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *Blockchain) AddTransaction(sender, recepient string, value float32) {
	txn := &Transaction{
		senderBlockchainAddress:    sender,
		recepientBlockchainAddress: recepient,
		value:                      value,
	}
	bc.transactionPool = append(bc.transactionPool, txn)
}

func (bc Blockchain) CopyTransactionPool() []*Transaction {
	txns := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		txns = append(txns, NewTransaction(t.senderBlockchainAddress, t.recepientBlockchainAddress, t.value))
	}
	return txns
}

func (bc Blockchain) ValidProof(nonce int, prevHash [32]byte, txns []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlk := Block{
		timestamp:    0,
		nonce:        nonce,
		previousHash: prevHash,
		transactions: txns,
	}
	guessBlkHashStr := fmt.Sprintf("%x", guessBlk.Hash())
	return guessBlkHashStr[:difficulty] == zeros
}

func (bc Blockchain) ProofOfWork() int {
	txns := bc.CopyTransactionPool()
	prevHash := bc.LastBlock().Hash()
	nonce := 0
	start := time.Now()
	for !bc.ValidProof(nonce, prevHash, txns, fragmos.MINING_DIFFICULTY) {
		nonce += 1
	}
	fmt.Printf("Time Taken to calculate Proof %s\n", time.Since(start))
	return nonce
}

// Mining accepts a receiver pointer even though there is no mutation done in this method.
// However, it calls other mutating methods like Add Transaction
// to which reference needs to be passed so that the mutation is successful
func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(fragmos.MINING_SENDER, bc.blockchainAddress, fragmos.MINING_REWARD)
	nonce := bc.ProofOfWork()
	prevHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, prevHash)
	log.Println("action=mining, status=success")
	return true
}

func (bc *Blockchain) CalculateTotalAmount(addr string) float32 {
	var totalAmount float32
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			value := t.value
			if addr == t.recepientBlockchainAddress {
				totalAmount += value
			}
			if addr == t.senderBlockchainAddress {
				totalAmount -= value
			}
		}
	}
	return totalAmount
}
