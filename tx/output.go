package tx
import "bc-demo/wallet"

type Output struct {
	Value int //金额
	To wallet.Address // 目标用户地址
}