package main

import(
  "fmt"
  "strconv"
)

func main() {
  bc := NewBlockChain()
  bc.AddBlock("sent 1 BTC to alpha")
  bc.AddBlock("sent 2 BTC to beta")
  bc.AddBlock("sent 3 BTC to gamma")

  for _, b := range bc.Block {
    fmt.Printf("Prev. hash: %x\n", b.PrevHash)
	fmt.Printf("Data: %s\n", b.Data)
	fmt.Printf("Hash: %x\n", b.Hash)
	pow := newPoW(b)
    fmt.Println(strconv.FormatBool(pow.Validate()))
    fmt.Println("")
  }
}
