package tx

import (
	"bc-demo/wallet"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)
//挖矿奖励金
//单位体系
const satoshi  =1//1中本聪
const S  =satoshi
const KS  =1000*S//千
const MS  =1000*KS//百万
const GS  = 1000*MS// 十亿
const BTC = 100000000 * satoshi
const CoinbaseSubsidy  =12*BTC
//交易类型
type TX struct {
	Hash string
	//输出集合
	Inputs []*Input
	// 输出集合
	Outputs []*Output
}
//CoinBase交易构造器
func NewCoinbaseTx(to wallet.Address)*TX  {
	//输入，没输出
	ins:=[]*Input{}
	//输出，仅存在一个输出，给目标为to的用户挖矿奖励
	output:=&Output{
		Value: CoinbaseSubsidy,//常量，储存挖矿奖励金
		To:    to,
	}
	outs:=[]*Output{
		output,
	}
	return NewTX(ins,outs)
}
//TX的构造器
func NewTX(ins[]*Input,outs[]*Output)*TX  {
    //tx数据
    tx:=&TX{
		Inputs:  ins,
		Outputs: outs,

	}
    //设置hash
    tx.setHash()
    return tx
}
//设置哈希
func (tx*TX)setHash()*TX {
	//先序列化
	ser, err := SerializeTX(*tx)
	if err != nil {
		log.Fatal(err)
	}
	//hash
	//在生成hash sha256
	hash := sha256.Sum256(ser)
	tx.Hash = fmt.Sprintf("%x", hash)
	return tx
}
//序列化Tx数据
func SerializeTX(tx TX)([]byte,error)  {
	buffer:=bytes.Buffer{}
	enc:=gob.NewEncoder(&buffer)
	//序列化
	if err:=enc.Encode(tx);err!=nil{
		return nil,err
	}
	return buffer.Bytes(),nil
}
//反序列化（反串行化，解码）
func  UnserializeTX(data []byte) (TX, error){
	buffer:=bytes.Buffer{}
    dec:=gob.NewDecoder(&buffer)
    buffer.Write(data)
    tx:=TX{}
    if err:=dec.Decode(&tx);err!=nil{
    	return tx,err
	}
	return tx,nil
}