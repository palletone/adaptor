package adaptor

import (
	"math/big"
)

type SimpleTransferTokenTx struct {
	TxId        []byte
	Asset       string
	FromAddress string
	ToAddress   string
	Amount      *big.Int
	Fee         *big.Int
	
}
