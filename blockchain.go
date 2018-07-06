package main

import (
		"fmt"
"log"
"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

type BlockChain struct {
	tip []byte
	db  *bolt.DB
}

func (bc *BlockChain) AddBlock(data string) {
  var lastHash []byte
  err := bc.db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(blocksBucket))
    lastHash = b.Get([]byte("l"))
    return nil
  })
  if err != nil {
    log.Panic(err)
  }

  newBlock := NewBlock(data, lastHash)

  err = bc.db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(blocksBucket))
	err := b.Put(newBlock.Hash, newBlock.Serialize())
	if err != nil {
      log.Panic(err)
    }
	err = b.Put([]byte("l"), newBlock.Hash)
	if err != nil {
	  log.Panic(err)
	}
	bc.tip = newBlock.Hash
	fmt.Println(newBlock.Hash)
    return nil
  })
}

//func NewBlockChain() *BlockChain {
//  return &BlockChain{[]*Block{GenesisBlock()}}
//}
func NewBlockChain() *BlockChain {
  var tip []byte
  db, err := bolt.Open(dbFile, 0600, nil)
  if err != nil {
	log.Panic(err)
  }
  err = db.Update(func(tx *bolt.Tx) error { //db.View 0400
    b := tx.Bucket([]byte(blocksBucket))
    if b == nil {
	  genesis := GenesisBlock()
	  b, err := tx.CreateBucket([]byte(blocksBucket))
	  if err != nil {
		log.Panic(err)
	  }
	  err = b.Put(genesis.Hash, genesis.Serialize())
	  fmt.Printf("genesis.Hash = %x\ngenesis.Serialize = %x\n", genesis.Hash, genesis.Serialize())
	  if err != nil {
        log.Panic(err)
      }
	  err = b.Put([]byte("l"), genesis.Hash)
	  if err != nil {
		log.Panic(err)
	  }
	  tip = genesis.Hash
	} else {
	  tip = b.Get([]byte("l"))
	}
	return nil
  })
  bc := BlockChain{tip, db}
  return &bc
}

type BlockChainInspect struct {
  CurrentHash []byte
  db		  *bolt.DB
}

func (bc *BlockChain) Inspect() *BlockChainInspect {
  bci := &BlockChainInspect{bc.tip, bc.db}
  fmt.Printf("bci %x\n", *bci)
  return bci
}

func (bci *BlockChainInspect) Next() *Block {
  var block *Block

  err := bci.db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(blocksBucket))
	encodedBlock := b.Get(bci.CurrentHash)
	fmt.Printf("bci.CurrentHash = %x\nget = %x\n", bci.CurrentHash, b.Get(bci.CurrentHash))//ok
	block = DeSerializeBlock(encodedBlock)
	//fmt.Printf("bci.block = %x\n", b.Get(bci.CurrentHash))
	return nil
  })
  if err != nil {
	log.Panic(err)
  }
  bci.CurrentHash = block.PrevHash
  fmt.Printf("Next.CurrentHash = %x\n", bci.CurrentHash)
  fmt.Printf("block.Next.Hash = %x\n", block.Hash)
  return block
}


