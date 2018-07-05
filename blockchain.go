package main

type BlockChain struct {
  Block []*Block
}

func (bc *BlockChain) AddBlock(data string) {
  //prevBlock := bc.Block[len(bc.Block) - 1]
  prevBlock := bc.Block[len(bc.Block)-1]
  newBlock := NewBlock(data, prevBlock.Hash)
  bc.Block = append(bc.Block, newBlock)
}

func NewBlockChain() *BlockChain {
  return &BlockChain{[]*Block{GenesisBlock()}}
}
