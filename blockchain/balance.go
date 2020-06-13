package blockchain

import "bc-demo/wallet"

func (bc *BlockChain) GetBalance(address wallet.Address) int {
	// 获取 address 对应的 UTXO
	// 统计余额
	balance := 0
	for _, utxo := range bc.FindUTXO(address) {
		balance += utxo.Value
	}

	return balance
}