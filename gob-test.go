package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Block struct {
	CurrHash string
	Txs string
}

func main()  {
	//Block数据，数据不需要指针类型
	b:=Block{
		CurrHash: "191918123456789123",
		Txs:      "fist transaction",
	}
	//gob编码
	//先得到编译器
	//定义可以写入内容的容器，通常使用byte型的缓存
	var bb bytes.Buffer
	//编码器需要该缓存，将未将编码的结果写入该缓存
	//提供的缓存，应该具备可写功能
	enc:=gob.NewEncoder(&bb)
	//编码数据，编码的结果，写入了编码器的缓存中。
	enc.Encode(b)
	//fmt.Println(bb.Bytes(),bb.String())
	result:=bb.Bytes()
	//解码
	//解码时，编码的数据，从缓冲中获取
	//提供的缓存，应该具备可读功能
	var bbr bytes.Buffer
	//将之前编码的数据，放入缓冲中
	bbr.Write(result)
	dec:=gob.NewDecoder(&bbr)
	b1:=Block{}
	dec.Decode(&b1)
	fmt.Println(b1)
}
