package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestJoin(t *testing.T) {
	s1:=[]byte("hello")
	s2:=[]byte("world")
	s3:=[]byte("你好")
	b:=bytes.Join([][]byte{s1,s2,s3},[]byte{})
	fmt.Println(string(b))
	fmt.Printf("%s\n",b)

}