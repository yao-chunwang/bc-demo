package tx

type Input struct {
	HashSrcTx string //输入来源交易的hash
	IndexSrcOutput int//输入来源交易输出的索引
}
