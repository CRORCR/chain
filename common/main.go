package main

import "fmt"

func main() {
	chain := NewBlockChain()
	chain.AddBlockChain("爸爸给果果一个btc")
	chain.AddBlockChain("爸爸给妈妈一个btc")
	for k,v:=range chain.blocks{
		fmt.Printf("======block num:%d \n",k)
		fmt.Printf("Date:%s \n",v.Date)
		fmt.Printf("PrevBlockHash:%x \n",v.PrevBlockHash)
		fmt.Printf("Version:%x \n",v.Version)
		fmt.Printf("Nonce:%x \n",v.Nonce)
		fmt.Printf("TimeStamp:%x \n",v.TimeStamp)
		fmt.Printf("MerkleBoot:%x \n",v.MerkleBoot)
		fmt.Printf("TargetBits:%x \n",v.TargetBits)

		pow := NewProofOfWork(v)
		fmt.Println("isvalid:",pow.IsValid())
	}
}
