package utils_chain

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)
//使用二进制来做一个中间转换,参数 int64
func IntToByte(num int64)[]byte{
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	CheckErr(err)
	return buffer.Bytes()
}
func CheckErr(err error) {
	if err != nil {
		fmt.Println("falied", err)
		os.Exit(1)
	}
}
