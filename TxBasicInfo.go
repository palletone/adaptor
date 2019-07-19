package adaptor

type TxBasicInfo struct {
	TxId        []byte
	TxData      []byte
	IsInBlock   bool //是否已经被打包到区块链中
	IsSucess    bool //是被标记为成功执行
	IsStable    bool //是否已经稳定不可逆
	BlockId     []byte
	BlockHeight uint
	TxIndex     uint //Tx在区块中的位置
	Timestamp   uint64
}
