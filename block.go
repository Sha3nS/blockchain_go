package main

import(
//		"fmt"
  "time"
)

type Block struct {
  Timestamp     int64
  Data          []byte
  PrevHash      []byte
  Hash          []byte
  Nonce         int
}

func NewBlock(data string, PrevHash []byte) *Block {
  block := &Block{time.Now().Unix(), []byte(data), PrevHash, []byte{}, 0}
  pow := newPoW(block)
  nonce, hash := pow.Run()
  block.Nonce = nonce
  block.Hash = hash[:]

  return block
}

func GenesisBlock() *Block {
  return NewBlock("GenesisBlock", []byte{})
}
