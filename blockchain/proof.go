package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

//take the data from the block

//create a counter (nonce) which starts at 0

//create a hash for the data plus the counter

//check the hash to see if it meets a set of requirements

//Requirements:
// the first few bytes must contain 0s

const Difficulty = 12

type ProofOfWork struct {
	Block *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork  {
	target := big.NewInt(1)
	target.Lsh(target,uint(256-Difficulty))
	pow :=&ProofOfWork{b,target}
	return pow
}

func (pow *ProofOfWork)IntData(nonce int)[]byte  {

	data := bytes.Join([][]byte{
		pow.Block.PrevHash,
		pow.Block.Data,
		ToHex(int64(nonce)),
		ToHex(int64(Difficulty)),
	},
	[]byte{})
	return data
}

func (pow *ProofOfWork)Run() (int, []byte)  {
	var intHash big.Int
	var hash [32]byte
	
	nonce :=0
	for nonce < math.MaxInt64 {
		data := pow.IntData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1{
			break
		}else {
			nonce ++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool  {
	var initHash big.Int
	data:= pow.IntData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	initHash.SetBytes(hash[:])
	return initHash.Cmp(pow.Target) == -1
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff,binary.BigEndian,num)
	if err !=nil{
		log.Panic(err)
	}
	return buff.Bytes()
}