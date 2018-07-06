package main

import(
  "time"
  "log"
  "bytes"
  "encoding/gob"
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
  pow := NewPoW(block)
  nonce, hash := pow.Run()
  block.Nonce = nonce
  block.Hash = hash[:]

  return block
}

func GenesisBlock() *Block {
  return NewBlock("GenesisBlock", []byte{})
}

func (b *Block) Serialize() []byte {
  var result bytes.Buffer
  encoder := gob.NewEncoder(&result)

  err := encoder.Encode(b)
  if err != nil {
    log.Panic(err)
  }
  return result.Bytes()
}

func DeSerializeBlock(b []byte) *Block {
  var block Block
  decoder := gob.NewDecoder(bytes.NewReader(b))
  err := decoder.Decode(&block)
  if err != nil {
    log.Panic(err)
  }
  return &block
}

