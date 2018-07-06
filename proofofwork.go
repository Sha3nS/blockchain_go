package main

import (
  "fmt"
  "math"
  "math/big"
  "bytes"
  "strconv"
  "crypto/sha256"
)
const hardByte = 4
const maxNonce = math.MaxInt64

type PoW struct {
  block *Block
  hard big.Int
}

func NewPoW(b *Block) *PoW {
  hard := big.NewInt(1)
  hard.Lsh(hard, uint(256 - hardByte))
  pow := &PoW{b, *hard}
  return pow
}

func IntToHex(n int64) []byte {
    return []byte(strconv.FormatInt(n, 16))
}

func (pow *PoW) prepareData(nonce int) []byte {
  data := bytes.Join([][]byte{pow.block.PrevHash,
							  pow.block.Data,
							  IntToHex(pow.block.Timestamp),
							  IntToHex(int64(hardByte)),
							  IntToHex(int64(nonce))},
					 []byte{})
  return  data
}

func (pow *PoW) Run() (int, []byte) {
  var hashInt big.Int
  var hash [32]byte
  nonce := 0

  fmt.Printf("Mining the block containing %s\n", string(pow.block.Data[:]))
  for nonce < maxNonce {
    data := pow.prepareData(nonce)
	hash = sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	if hashInt.Cmp(&pow.hard) == -1 {
	  break
	} else {
	  nonce++
	}
  }
  return nonce, hash[:]
}

func (pow *PoW) Validate() bool {
  var hashInt big.Int
  data := pow.prepareData(pow.block.Nonce)
  hash := sha256.Sum256(data)
  hashInt.SetBytes(hash[:])

  isValid := hashInt.Cmp(&pow.hard) == -1
  return isValid
}
