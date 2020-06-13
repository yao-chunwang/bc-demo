package block

import (
	"bc-demo/tx"
	"bytes"
	"encoding/gob"
	"time"
)

//区块数据
type BlockData struct {
	Version        int
	HashPrevBlock  Hash
	HashMerkleRoot Hash
	Time           time.Time
	Bits           int
	Nonce          int
	Txs            []*tx.TX
	TxCounter      int
	HashCurr       Hash
}
//区块序列化
func BlockSerialize(b Block)([]byte,error)  {
	// 由于区块的字段都是 unexported field 非导出字段
	// 使用中间的数据结构作为桥梁，完成序列化。
	// 将区块数据赋值到 BlockData
	bd:=BlockData{
		Version:        b.header.version,
		HashPrevBlock:  b.header.hashPrevBlock,
		HashMerkleRoot: b.header.hashMerkleRoot,
		Time:           b.header.time,
		Bits:           b.header.bits,
		Nonce:          b.header.nonce,
		Txs:            b.txs,
		TxCounter:      b.txCounter,
		HashCurr:      b.hashCurr,
	}
//执行gob序列化即可
buffer:=bytes.Buffer{}
//编码器
enc:=gob.NewEncoder(&buffer)
//编码，序列化
if err:=enc.Encode(bd);err!=nil{
	return nil,err
}
//编码成功
return buffer.Bytes(),nil
}
//区块反序列化
func BlockUnSerialize(data []byte)(Block,error)  {
	//得到装有内容的缓冲
	buffer:=bytes.Buffer{}
	buffer.Write(data)
	//解码器
	dec:=gob.NewDecoder(&buffer)
	//解码，反序列化
	bd:=BlockData{}
	if err:=dec.Decode(&bd);err!=nil{
		return Block{},err
	}
	//反序列化成功
	return Block{
		header:    BlockHeader{
			version:bd.Version,
			hashPrevBlock:bd.HashPrevBlock,
			hashMerkleRoot:bd.HashMerkleRoot,
			time:bd.Time,
			bits:bd.Bits,
			nonce:bd.Nonce,

		},
		txs:bd.Txs,
		txCounter: bd.Bits,
		hashCurr:  bd.HashCurr,
	},nil
}