package adaptor

//钱包相关的API接口,区块链相关API接口
type IUtility interface {
	//创建一个新的私钥
	NewPrivateKey() (priKey string)
	//根据私钥创建公钥
	GetPublicKey(priKey string) (pubKey string)
	//根据Key创建地址
	GetAddress(priKey string) (address string)
	SignMessage(message []byte) (signature []byte)
	SignTransaction(params string) string

	//根据区块Hash获得区块头信息
	//GetBlockHeader(blockHash string) []byte
}
