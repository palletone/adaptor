package adaptor

type ISmartContract interface {
	//创建一个安装合约的交易，未签名
	CreateContractInstallTx(addr string, fee *AmountAsset, contract []byte, extra []byte) (tx []byte, err error)
	//上传合约代码到区块链网络,返回签名后的交易
	// InstallContract(priKey []byte, fee *AmountAsset, contract []byte,extra []byte) (signedTx []byte,error)
	//查询合约安装的结果
	GetContractInstallTxByTxId(txid []byte) (*ContractInstallTx, error)
	//初始化合约实例
	CreateContractInitialTx(addr string, fee *AmountAsset, contract []byte, args [][]byte, extra []byte) (tx []byte, err error)
	GetContractInitialTxByTxId(txid []byte) (*ContractInitialTx, error)
	//调用合约方法
	CreateContractInvokeTx(addr string, fee *AmountAsset, contractAddress string, function string, args [][]byte, extra []byte) (tx []byte, err error)
	GetContractInvokeTxByTxId(txid []byte) (*ContractInvokeTx, error)
	//查询合约方法
	QueryContract(addr string, fee *AmountAsset, contractAddress string, function string, args [][]byte, extra []byte) (result []byte, err error)
	//销毁合约
	// CreateContractDestoryTx(addr string, fee *AmountAsset, contractAddress string, args [][]byte, extra []byte) (tx []byte, err error)
	// GetContractDestoryTxByTxId(txid []byte) (*ContractDestoryTx, error)
}
