package main

import "os"

type BlockChain struct {
	blocks []*Block
}

//构造区块链 先存储一个创世区块
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGenesisBlock()}}
}

//区块链添加一个区块
func (bc *BlockChain) AddBlockChain(data string) {
	//校验下标越界
	if len(bc.blocks)==0{
		os.Exit(1)
	}
	//hash是上一个区块hash值
	b := bc.blocks[len(bc.blocks)-1]
	block := NewBlock([]byte(data), b.Hash)
	bc.blocks = append(bc.blocks, block)
}
