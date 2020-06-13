package pow

import (
	"bc-demo/block"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

type ProofOfWork struct {
	//需要pow工作量区块的区块
	block *block.Block
	//证明参数目标
	target *big.Int
}
//构造方法
func NewPOW(b *block.Block) *ProofOfWork {
	  p:=&ProofOfWork{
		  block: b,
		  target:big.NewInt(1),
	  }
	  //计算target
	p.target.Lsh(p.target, uint(block.HashLen- b.GetBits() - 1))
	return p
}
//hashcash证明
//返回使用的nonce和形成的区块hash
func (p *ProofOfWork)Proof()(int, block.Hash)  {
    var hashInt big.Int
    //基于block准备servicestr
    serviceStr:=p.block.GenServiceStr()
    //noce计数器
    nonce:=0
    //迭代计算hash，设置防noncce溢出的条件
    for nonce<=math.MaxInt64{
    	//生成hash
    	hash:=sha256.Sum256([]byte(serviceStr+strconv.Itoa(nonce)))
    	//得到hash的big.Int
    	hashInt.SetBytes(hash[:])

		fmt.Printf("%x \t %d \n",hash,nonce)
    	//判断是否满足难度
    	if hashInt.Cmp(p.target)==-1{
    		//解决问题
    		return nonce, block.Hash(fmt.Sprintf("%x",hash))
		}
    	nonce++
	}
    return 0,""
}
//验证
func (p*ProofOfWork)Validate()bool  {
	//验证区块hash是否正确
	//再次生成hash
	serviceStr:=p.block.GenServiceStr()
	data:=serviceStr+strconv.Itoa(p.block.GentNonce())
	hash:=sha256.Sum256([]byte(data))
	//比较是否相等
	if p.block.GetHashCurr()!=fmt.Sprintf("%x",hash){
		return false
	}
	//比较是否满足难题
	target:=big.NewInt(1)
	target.Lsh(target,uint(block.HashLen- p.block.GetBits() - 1))
	hashInt:=new(big.Int)
	hashInt.SetBytes(hash[:])
	//不小于
	if hashInt.Cmp(target)!=-1{
		return false
	}
	return  true

}
