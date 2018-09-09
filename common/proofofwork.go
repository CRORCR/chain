package main

import (
	"bytes"
	"chain/utils_chain"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 2

type ProofOfWork struct {
	block     *Block
	targetBit *big.Int
}

//创建pow,设置难度值
func NewProofOfWork(b *Block) *ProofOfWork {
	var target = big.NewInt(1)
	//目标值,往左移动 256-24位,然后与算出来的hash值比较,如果hash值小于目标值,就任务找到了nonce
	target.Lsh(target, uint(256-targetBits))
	return &ProofOfWork{b, target}
}

//拼接hash,nonce需要传入
func (pow *ProofOfWork) PrepareData(nonce int64) []byte {
	block := pow.block
	tmp := [][]byte{
		utils_chain.IntToByte(nonce), //nonce需要传入,其他跟block sethash函数是一样的
		block.PrevBlockHash,
		block.MerkleBoot,
		block.Date,
		utils_chain.IntToByte(block.TimeStamp),
		utils_chain.IntToByte(block.Version),
		utils_chain.IntToByte(targetBits), //难度值
	}
	//区块的各个字段 二维byte,转为一维byte,使用join
	date := bytes.Join(tmp, []byte{})
	return date
}

//hash碰撞 找到对应难度值的nonce
func (pow *ProofOfWork) Run() (int64, []byte) {
	var nonce int64
	var hash [32]byte
	var hashInt big.Int

	//防止死循环
	for nonce < math.MaxInt64 {
		data := pow.PrepareData(nonce)
		//算出来的hash不能直接参与运行,需要 byte转int(使用bigInt)
		hash = sha256.Sum256(data)
		//需要切片,数组转切片
		hashInt.SetBytes(hash[:])

		//hashInt < targetBit=-1
		if hashInt.Cmp(pow.targetBit) == -1 {
			fmt.Printf("find nonce,hash:%v \n",hashInt)
			break
		} else {
			nonce++
		}
	}
	return nonce,hash[:]
}

//校验是否有效
func(pow *ProofOfWork)IsValid()bool{
	data := pow.PrepareData(pow.block.Nonce)
	hash:=sha256.Sum256(data)
	var intHash big.Int
	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.targetBit)==-1
}