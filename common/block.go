package common

import (
	"bytes"
	"chain/utils_chain"
	"crypto/sha256"
	"time"
)

type Block struct {
	//版本号
	Version int64
	//父区块hash
	PrevBlockHash []byte
	Hash []byte
	//时间戳
	TimeStamp int64
	//merkle根
	MerkleBoot []byte
	//难度
	TargetBits int64
	Nonce int64
	//区块体
	Date []byte
}
//前一个区块ha;sh值
func  NewBlock(data ,fatherHash []byte)*Block {
	block:=&Block{
		Version: 1,
		PrevBlockHash: fatherHash,
		TimeStamp: time.Now().Unix(),
		MerkleBoot:[]byte{},
		TargetBits: 1,
		Nonce:0,
		Date: data,
		}
	block.setHash()
	return block
}

func (block *Block)setHash(){
	tmp:=[][]byte{
		utils_chain.IntToByte(block.Nonce),
		block.MerkleBoot,
		block.Date,
		utils_chain.IntToByte(block.TimeStamp),
		utils_chain.IntToByte(block.Version),
	}
	//二维byte,转为一维byte
	date:=bytes.Join(tmp,[]byte{})
	hash := sha256.Sum256(date)
	block.PrevBlockHash=hash[:]

}