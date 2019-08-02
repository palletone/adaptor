package adaptor

//加密货币相关的API
type ICryptoCurrency interface {
	//获取某地址下持有某资产的数量,返回数量为该资产的最小单位
	GetBalance(addr string, asset string) (*AmountAsset, error)
	//获取某资产的小数点位数
	GetAssetDecimal(asset string) (uint, error)
	//创建一个转账交易，但是未签名
	CreateTransferTokenTx(from, to string, amount *AmountAsset, fee *AmountAsset) ([]byte, error)
	//获取某个地址对某种Token的交易历史,支持分页和升序降序排列
	GetAddrTxHistory(addr, asset string, pageSize, pageIndex int, desc bool) ([]*SimpleTransferTokenTx, error)
	//根据交易ID获得对应的转账交易
	GetTransferTxByTxId(txid string) (*SimpleTransferTokenTx, error)
	//创建一个多签地址，该地址必须要满足signCount个签名才能解锁
	CreateMultiSigAddress(keys [][]byte, signCount int, extra []byte) (addr string, extraResult []byte, err error)
}
