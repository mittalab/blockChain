package model

import (
	"fmt"
)

type BlockChain struct {
	blocks []*Block
}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, new)
}

func (bc *BlockChain) AddOneBlock(b *Block) {
	bc.blocks = append(bc.blocks, b)
}

func (bc *BlockChain) PrintChain() {
	for _, b := range bc.blocks {
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("PrevHash: %x\n", b.PrevHash)
	}
}
