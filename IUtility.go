package adaptor

import (
	"math/big"
)

//钱包相关的API接口,区块链相关API接口
type IUtility interface {
	//创建一个新的私钥
	NewPrivateKey() (priKey []byte, err error)
	//根据私钥创建公钥
	GetPublicKey(priKey []byte) (pubKey []byte, err error)
	//根据Key创建地址
	GetAddress(key []byte) (address string, err error)
	//对一条消息进行签名
	SignMessage(addr string, message []byte) (signature []byte, err error)
	//对一条交易进行签名，并返回签名结果
	SignTransaction(addr string, tx []byte) (signature []byte, err error)
	//创建一个多签地址，该地址必须要满足signNeed个签名才能解锁
	CreateMultiSigAddress(keys [][]byte, signNeed int) (addr string, extra []byte, err error)
	//将未签名的原始交易与签名进行绑定，返回一个签名后的交易
	BindTxAndSignature(tx []byte, signs [][]byte, extra []byte) (signedTx []byte, err error)
	//根据交易内容，计算交易Hash
	GetTxHash(tx []byte) (hash []byte, err error)
	//将签名后的交易广播到网络中,如果发送交易需要手续费，指定最多支付的手续费
	SendTx(tx []byte, fee *big.Int) (txID []byte, err error)
	//根据交易ID获得交易的基本信息
	GetTxBasicInfo(txID []byte) (*TxBasicInfo, error)
}
