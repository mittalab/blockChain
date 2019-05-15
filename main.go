package main

import (
	"github.com/mittalab/blockChain/model"
	log "github.com/sirupsen/logrus"
)

func InitBlockChain() *model.BlockChain {
	bc := &model.BlockChain{}
	b := model.CreateGenesisBlock()

	bc.AddOneBlock(b)
	return bc
}

func main() {
	log.Println("Hello")

	bc := InitBlockChain()
	bc.AddBlock("Block1")
	bc.AddBlock("Block2")
	bc.AddBlock("Block3")
	bc.AddBlock("Block4")

	bc.PrintChain()
}
