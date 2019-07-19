package adaptor

type ISmartContract interface {
	//上传合约代码到区块链网络
	InstallContract(addr string)
	//初始化合约实例
	InitialContract(addr string)
	//调用合约方法
	InvokeContract(addr string)
	//查询合约方法
	QueryContract(addr string)
	//销毁合约
	DestoryContract(addr string)
}
