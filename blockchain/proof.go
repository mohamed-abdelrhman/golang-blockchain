package blockchain

import "math/big"

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
	target.Lsh(target,uint(265-Difficulty))
	pow :=&ProofOfWork{b,target}
	return pow
}

