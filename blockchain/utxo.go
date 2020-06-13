package blockchain

import (
	"bc-demo/tx"
	"bc-demo/wallet"
)
// 找到哪些 属于 address 的未被花费的输出 Unspend
func (bc *BlockChain) FindUTXO(address wallet.Address) []*tx.Output {
	utxo := []*tx.Output{}
	// 用于统计全部已花费的输入
	// map 的 key 是交易的 哈希
	// map 的 value 是输出的索引切片
	spentOuts := map[string][]int{}
	// 三层结构：区块链（区块） -> 区块的交易（每个交易）-> 交易中输入或输出（每个输入或输出）
	// 遍历区块，找
	bci := NewBCIterator(bc)
	for block, err := bci.Next(); err == nil; block, err = bci.Next() {
		// 遍历交易
		for _, tx := range block.GetTxs() {
			// tx 每个区块上的交易
			// 统计交易的输入
			for _, in := range tx.Inputs {
				// 记录下交易标识和输出索引
				if _, e := spentOuts[tx.Hash]; !e {
					// 该交易hash key 不存在，初始化
					spentOuts[tx.Hash] = []int{}
				}
				// 该交易key已经存在，追加输出索引即可
				spentOuts[tx.Hash] = append(spentOuts[tx.Hash], in.IndexSrcOutput)
			}
			// 遍历交易上的每个输出
			for i, o := range tx.Outputs {
				// o 交易上的某个输出
				// i 输出的索引
				// 属于我的输出
				if o.To == address && checkUnspent(spentOuts, tx.Hash, i) {
					utxo = append(utxo, o)
				}
			}
		}

	}

	return utxo
}
// 花费为false，未花费true
func checkUnspent(spentOuts map[string][]int, txHash string, i int) bool {
	indexs, e := spentOuts[txHash]
	// 该交易不在其中
	if !e {
		return true
	}
	// 继续监测索引是否匹配
	for _, index := range indexs {
		if index == i {
			return false
		}
	}

	return true
}
