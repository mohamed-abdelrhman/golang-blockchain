package main

import (
	"fmt"
	"github.com/mohamed-abdelrhman/golang-blockchain/blockchain"
	"math/big"
)

func main()  {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First")
	chain.AddBlock("second")
	chain.AddBlock("Third")
	target := big.NewInt(1)
	fmt.Println(target,uint(265))

	target.Lsh(target,uint(265))
	//fmt.Println(target)
	//pow :=&ProofOfWork{b,target}
	//return pow

}

