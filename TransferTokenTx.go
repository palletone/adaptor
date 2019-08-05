package adaptor

//一个简单的Token转账交易
type SimpleTransferTokenTx struct {
	TxBasicInfo
	FromAddress string       `json:"from_address"` //转出地址
	ToAddress   string       `json:"to_address"`   //转入地址
	Amount      *AmountAsset `json:"amount"`       //转账金额
	Fee         *AmountAsset `json:"fee"`          //转账交易费
	AttachData  []byte       `json:"attach_data"`  //附加的数据（备注之类的）
}

//多地址对多地址的转账交易
// type MultiAddrTransferTokenTx struct {
// 	TxBasicInfo
// 	FromAddress map[string]*AmountAsset //转出地址
// 	ToAddress   map[string]*AmountAsset //转入地址
// 	Fee         *AmountAsset            //转账交易费
// 	AttachData  []byte                  //附加的数据（备注之类的）
// }
