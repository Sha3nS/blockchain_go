package main

import(
		"fmt"
)

func main() {
  bc := NewBlockChain()
  fmt.Printf("bc.tip = %x\n", bc.tip)
  bci := bc.Inspect()
  fmt.Printf("bci.CurrentHash = %x\n", bci.CurrentHash)
  block := bci.Next()
  fmt.Printf("block is %x\n", *block)
  //fmt.Printf("block.Hash = %x\n", block.Hash)
  //bc.AddBlock("sent 1 BTC to alpha")
  //bc.AddBlock("sent 2 BTC to beta")
}
