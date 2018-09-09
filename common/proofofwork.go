package main

import "math/big"

const targetBits  = 24
type ProofOfWork struct {
	block *Block
	targetBit *big.Int
}

func NewProofOfWork(b *Block)*ProofOfWork{
	var target = big.NewInt(1)
	//目标值,往左移动 256-24位,然后与算出来的hash值比较,如果hash值小于目标值,就任务找到了nonce
	target.Lsh(target,uint(256-targetBits))
	return &ProofOfWork{b,target}
}
