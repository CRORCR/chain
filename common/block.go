package common

import (
	"bytes"
	"chain/utils_chain"
	"crypto/sha256"
	"time"
)

//block是不存储hash值,hash值是节点收到交易,自己计算得出,存储本地,当前版本是硬编码本地,
// 为了方便而做的简化,正常比特币区块是不包含hash值的
type Block struct {
	//版本号
	Version int64
	//父区块hash
	PrevBlockHash []byte
	Hash          []byte
	//时间戳
	TimeStamp int64
	//merkle根
	MerkleBoot []byte
	//难度值
	TargetBits int64
	Nonce      int64
	//区块体 这应该是一个交易(transaction)
	Date []byte
}

//需要交易和前一个区块hash值
func NewBlock(data, fatherHash []byte) *Block {
	block := &Block{
		Version:       1,
		PrevBlockHash: fatherHash,
		TimeStamp:     time.Now().Unix(),
		MerkleBoot:    []byte{},
		TargetBits:    1,
		Nonce:         0,
		Date:          data,
	}
	block.setHash()
	return block
}

func (block *Block) setHash() {
	tmp := [][]byte{
		utils_chain.IntToByte(block.Nonce),
		block.MerkleBoot,
		block.Date,
		utils_chain.IntToByte(block.TimeStamp),
		utils_chain.IntToByte(block.Version),
	}
	//二维byte,转为一维byte
	date := bytes.Join(tmp, []byte{})
	//入切片,返数组
	hash := sha256.Sum256(date)
	block.PrevBlockHash = hash[:]
}