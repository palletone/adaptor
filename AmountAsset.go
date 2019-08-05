package adaptor

import (
	"math/big"
)

type AmountAsset struct {
	Amount big.Int `json:"amount"`
	Asset  string  `json:"asset"`
}
