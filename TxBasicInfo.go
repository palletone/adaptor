package adaptor

type TxBasicInfo struct {
	TxId           []byte //交易的ID，Hash
	TxRawData      []byte //交易的二进制数据
	CreatorAddress string //交易的发起人
	TargetAddress  string //交易的目标地址（被调用的合约、收款人）
	IsInBlock      bool   //是否已经被打包到区块链中
	IsSucess       bool   //是被标记为成功执行
	IsStable       bool   //是否已经稳定不可逆
	BlockId        []byte //交易被打包到了哪个区块ID
	BlockHeight    uint   //交易被打包到的区块的高度
	TxIndex        uint   //Tx在区块中的位置
	Timestamp      uint64 //交易被打包的时间戳
}
