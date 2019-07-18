package adaptor

import (
	"math/big"
)

type SimpleTransferTokenTx struct {
	Asset       string
	FromAddress string
	ToAddress   string
	Amount      *big.Int
	Fee         *big.Int
	TxId        string
}
