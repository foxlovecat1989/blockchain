package blockchain

import (
	"encoding/json"
	"log"
)

type transaction struct {
	sender    string
	recipient string
	value     float32
}

func (t *transaction) Print() {
	log.Printf("SendAddress: %v\n", t.sender)
	log.Printf("RecipientAddress: %v\n", t.recipient)
	log.Printf("Value: %.1f\n", t.value)
}
func (t *transaction) ToMarshalJSON() []byte {
	bs, _ := json.Marshal(struct {
		SendAddress      string  `json:"sendAddress"`
		RecipientAddress string  `json:"recipientAddress"`
		Value            float32 `json:"value"`
	}{
		SendAddress:      t.sender,
		RecipientAddress: t.recipient,
		Value:            t.value,
	})

	return bs
}
