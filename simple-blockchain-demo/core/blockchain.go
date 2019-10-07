package Blockchain

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(transactions string) {
	newBlock := NewBlock(transactions, bc.previousBlock().Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func CreateGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{CreateGenesisBlock()}}
}

func (bc *Blockchain) previousBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}
