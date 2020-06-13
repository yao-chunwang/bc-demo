package block

import (
	"bc-demo/tx"
	"fmt"
	"strings"
	"time"
)

type Hash = string

const blockBits  =16
const nodeVersion  = 0
const HashLen = 256
//Block主体的主要作用存储数据（交易）
type Block struct {
	header BlockHeader
	txs  []*tx.TX //交易列表
	txCounter int//交易计数器
	hashCurr Hash //当前区块hash，算法sha256
}
//BlockHeader区块的头信息，包含GG对区块的说明部分（区块的元数据）
type BlockHeader struct {
	version int //版本信息，节点更新时版本迭代
	hashPrevBlock Hash //前一个节点的hash值
	hashMerkleRoot Hash//默克尔树根节点
	time time.Time//节点生成时间
	bits int//难度系数
	nonce int//随机计时器
}
//设置当前区块hash
func (b*Block)SetHashCurr(hash Hash)*Block  {
	//计算hash 值
	b.hashCurr=hash
	return b
}
//头信息的字符串化
func (bh *BlockHeader)Stringify()string  {
	return  fmt.Sprintf("%d%s%s%d%d%d",
		bh.version,
		bh.hashPrevBlock,
		bh.hashMerkleRoot,
		bh.time.UnixNano(),
		bh.bits,
		bh.nonce,
		)
}
// 构造区块
func NewBlock(prevHash Hash) *Block {
	// 实例化Block
	b := &Block{
		header:    BlockHeader{
			version: nodeVersion,
			hashPrevBlock: prevHash, // 设置前面的区块哈希
			time: time.Now(),
			bits:blockBits,
		},
	/*	txs:       txs, // 设置数据
		txCounter: 1, // 计数交易*/
	}
	// 计算设置当前区块的哈希
	return b
}
//bits属性的getter
func (b*Block)GetBits()int  {
	return b.header.bits
}
func (b*Block) GenServiceStr() string  {
	return fmt.Sprintf("%d%s%s%s%d",
		b.header.version,
		b.header.hashPrevBlock,
		b.header.hashMerkleRoot,
		b.header.time.Format("2006-01-02 15:04:05.999999999 -0700 MST"),
		b.header.bits,
	)
}
func (b*Block)SetNonce(nonce int)*Block  {
	b.header.nonce=nonce
	return b
}
func (b *Block) GetHashCurr() Hash {
	return b.hashCurr
}
func (b *Block) GetTxs() []*tx.TX{
	return b.txs
}
func (b *Block) GetTime() time.Time {
	return b.header.time
}
func (b *Block) GetHashPrevBlock() Hash {
	return b.header.hashPrevBlock
}
func (b*Block)GentNonce()int  {
	return b.header.nonce
}
//添加交易
func (b*Block)AddTX(tx*tx.TX)*Block  {
	//添加
	b.txs=append(b.txs,tx)
	b.txCounter++
	return b
}
func (b*Block) GetTxsString() string {
  show:=fmt.Sprintf("%d tansactions\n",b.txCounter)
  txStr:=[]string{}
  for i,t:=range b.txs{
	  txStr = append(txStr, fmt.Sprintf("\tindex:%d, Hash: %s, Inputs: %d, Ouputs: %d", i, t.Hash, len(t.Inputs), len(t.Outputs)))
  }
  return show +strings.Join(txStr,"\n")
}