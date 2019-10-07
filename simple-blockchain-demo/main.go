package main

import (
	"fmt"
	. "github.com/chen-shiwei/Blockchain/core"
)

func main() {
	bc := NewBlockchain()
	bc.AddBlock("send 1 btc to pangzi")
	bc.AddBlock("send 1100 btc to pangzi")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PreviousHash)
		fmt.Printf("Transactions: %s\n", block.Transactions)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
