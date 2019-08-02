package adaptor

//一个简单的Token转账交易
type SimpleTransferTokenTx struct {
	TxBasicInfo
	FromAddress string       //转出地址
	ToAddress   string       //转入地址
	Amount      *AmountAsset //转账金额
	Fee         *AmountAsset //转账交易费
	AttachData  []byte       //附加的数据（备注之类的）
}

//多地址对多地址的转账交易
// type MultiAddrTransferTokenTx struct {
// 	TxBasicInfo
// 	FromAddress map[string]*AmountAsset //转出地址
// 	ToAddress   map[string]*AmountAsset //转入地址
// 	Fee         *AmountAsset            //转账交易费
// 	AttachData  []byte                  //附加的数据（备注之类的）
// }
