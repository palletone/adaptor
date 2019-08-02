package adaptor

//一个简单的Token转账交易
type ContractInstallTx struct {
	TxBasicInfo
	TemplateId []byte
}
type ContractInitialTx struct {
	TxBasicInfo
	ContractAddress string
}
type ContractInvokeTx struct {
	TxBasicInfo
	UpdateStateSuccess bool //读写集一致，成功更新StateDB
	InvokeResult       []byte
}
