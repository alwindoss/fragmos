package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

func NewBlock(nonce int, prevHash [32]byte, txns []*Transaction) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		previousHash: prevHash,
		transactions: txns,
	}
}

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*Transaction
}

func (b Block) Print() {
	fmt.Printf("Nonce               \t%d\n", b.nonce)
	fmt.Printf("PreviousHash        \t%x\n", b.previousHash)
	fmt.Printf("Timestamp           \t%d\n", b.timestamp)
	for _, t := range b.transactions {
		t.Print()
	}
}

func (b Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

func (b Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}
