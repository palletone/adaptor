package adaptor

type ISmartContract interface {
	//创建一个安装合约的交易，未签名
	CreateContractInstallTx(input *CreateContractInstallTxInput) (*CreateContractInstallTxOutput, error)
	//查询合约安装的结果的交易
	GetContractInstallTx(input *GetContractInstallTxInput) (*GetContractInstallTxOutput, error)
	//初始化合约实例
	CreateContractInitialTx(input *CreateContractInitialTxInput) (*CreateContractInitialTxOutput, error)
	//查询初始化合约实例的交易
	GetContractInitialTx(input *GetContractInitialTxInput) (*GetContractInitialTxOutput, error)
	//调用合约方法
	CreateContractInvokeTx(input *CreateContractInvokeTxInput) (*CreateContractInvokeTxOutput, error)
	//查询调用合约方法的交易
	GetContractInvokeTx(input *GetContractInvokeTxInput) (*GetContractInvokeTxOutput, error)
	//调用合约的查询方法
	QueryContract(input *QueryContractInput) (*QueryContractOutput, error)
	//销毁合约
	// CreateContractDestoryTx(addr string, fee *AmountAsset, contractAddress string, args [][]byte, extra []byte) (tx []byte, err error)
	// GetContractDestoryTxByTxId(txid []byte) (*ContractDestoryTx, error)
}
type CreateContractInstallTxInput struct {
	Address  string       `json:"address"`
	Fee      *AmountAsset `json:"fee"`
	Contract []byte       `json:"contract"`
	Extra    []byte       `json:"extra"`
}
type CreateContractInstallTxOutput struct {
	RawTranaction []byte `json:"raw_transaction"`
}
type CreateContractInitialTxInput struct {
	Address  string       `json:"address"`
	Fee      *AmountAsset `json:"fee"`
	Contract []byte       `json:"contract"`
	Args     [][]byte     `json:"args"`
	Extra    []byte       `json:"extra"`
}
type CreateContractInitialTxOutput struct {
	RawTranaction []byte `json:"raw_transaction"`
}
type CreateContractInvokeTxInput struct {
	Address         string       `json:"address"`
	Fee             *AmountAsset `json:"fee"`
	ContractAddress string       `json:"contract_address"`
	Function        string       `json:"function"`
	Args            [][]byte     `json:"args"`
	Extra           []byte       `json:"extra"`
}
type CreateContractInvokeTxOutput struct {
	RawTranaction []byte `json:"raw_transaction"`
}
type QueryContractInput struct {
	Address         string       `json:"address"`
	Fee             *AmountAsset `json:"fee"`
	ContractAddress string       `json:"contract_address"`
	Function        string       `json:"function"`
	Args            [][]byte     `json:"args"`
	Extra           []byte       `json:"extra"`
}
type QueryContractOutput struct {
	QueryResult []byte `json:"query_result"`
}

type GetContractInstallTxInput struct {
	TxID []byte `json:"tx_id"`
}
type GetContractInstallTxOutput struct {
	TxBasicInfo
	TemplateId []byte `json:"template_id"`
}
type GetContractInitialTxInput struct {
	TxID []byte `json:"tx_id"`
}
type GetContractInitialTxOutput struct {
	TxBasicInfo
	ContractAddress string `json:"contract_address"`
}
type GetContractInvokeTxInput struct {
	TxID []byte `json:"tx_id"`
}
type GetContractInvokeTxOutput struct {
	TxBasicInfo
	UpdateStateSuccess bool   `json:"update_state_success"` //读写集一致，成功更新StateDB
	InvokeResult       []byte `json:"invoke_result"`
}
