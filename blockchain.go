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
	//fmt.Println(newBlock.Hash)
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
	  fmt.Println("No existing blockchain found. Creating a new one...")
	  genesis := GenesisBlock()
	  b, err := tx.CreateBucket([]byte(blocksBucket))
	  if err != nil {
		log.Panic(err)
	  }

	  err = b.Put(genesis.Hash, genesis.Serialize())
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
  return bci
}

func (bci *BlockChainInspect) Next() *Block {
  var block *Block

  err := bci.db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte(blocksBucket))
	encodedBlock := b.Get(bci.CurrentHash)
	block = DeSerializeBlock(encodedBlock)
	return nil
  })
  if err != nil {
	log.Panic(err)
  }
  bci.CurrentHash = block.PrevHash
  return block
}


