package Blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Transaction struct {
	Sender   string
	Recipent string
	Amount   float64
}

type Block struct {
	Timestamp int64
	//证明
	Transactions []byte
	Hash         []byte
	PreviousHash []byte
}

func NewBlock(transactions string, previousHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(transactions), previousHash, []byte{}}
	block.CreateHash()
	return block
}

func (this *Blockchain) newTransaction(sender string, recipent string, amount float64) (Transactions []Transaction) {
	return append(Transactions, Transaction{sender, recipent, amount})
}

func (b *Block) CreateHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PreviousHash, b.Transactions, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
