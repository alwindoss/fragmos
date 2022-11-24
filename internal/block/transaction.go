package block

import (
	"encoding/json"
	"fmt"
	"strings"
)

func NewTransaction(sender, recepient string, value float32) *Transaction {
	return &Transaction{
		senderBlockchainAddress:    sender,
		recepientBlockchainAddress: recepient,
		value:                      value,
	}
}

type Transaction struct {
	senderBlockchainAddress    string
	recepientBlockchainAddress string
	value                      float32
}

func (t Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("SenderAddress       \t%s\n", t.senderBlockchainAddress)
	fmt.Printf("RecepientAddress    \t%s\n", t.recepientBlockchainAddress)
	fmt.Printf("Value               \t%.1f\n", t.value)
}

func (t Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
		RecepientBlockchainAddress string  `json:"recepient_blockchain_address"`
		Value                      float32 `json:"value"`
	}{
		SenderBlockchainAddress:    t.senderBlockchainAddress,
		RecepientBlockchainAddress: t.recepientBlockchainAddress,
		Value:                      t.value,
	})
}
