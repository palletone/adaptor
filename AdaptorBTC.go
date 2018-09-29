/*
   This file is part of go-palletone.
   go-palletone is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-palletone is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with go-palletone.  If not, see <http://www.gnu.org/licenses/>.
*/
/*
 * @author PalletOne core developers <dev@pallet.one>
 * @date 2018
 */
package adaptor

type RPCParams struct {
	Host      string `json:"host"`
	RPCUser   string `json:"rpcUser"`
	RPCPasswd string `json:"rpcPasswd"`
	CertPath  string `json:"certPath"`
}

type AdaptorBTC struct {
	NetID int
	RPCParams
}

type NetID int

const (
	NETID_MAIN = iota
	NETID_TEST
)

type adapterbtc interface {
	NewPrivateKey() (wifPriKey string)
	GetPublicKey(wifPriKey string) (pubKey string)
	GetAddress(wifPriKey string) (address string)
	GetAddressByPubkey(pubKeyHex string) (string, error)
	CreateMultiSigAddress(params *CreateMultiSigParams)

	RawTransactionGen(params *RawTransactionGenParams) (string, error)
	DecodeRawTransaction(params *DecodeRawTransactionParams) (string, error)
	GetTransactionByHash(params *GetTransactionByHashParams) (string, error)

	SignTransaction(params *SignTransactionParams) (string, error)
	SignTxSend(params *SignTxSendParams) (string, error)

	GetBalance(params *GetBalanceParams) (string, error)
	GetTransactions(params *GetTransactionsParams) (string, error)
	ImportMultisig(params *ImportMultisigParams) (string, error)
	SendTransaction(params string) string
}

//
type CreateMultiSigParams struct {
	PublicKeys []string `json:"publicKeys"`
	N          int      `json:"n"`
	M          int      `json:"m"`
}
type CreateMultiSigResult struct {
	P2ShAddress  string   `json:"p2sh_address"`
	RedeemScript string   `json:"redeem_script"`
	Addresses    []string `json:"addresses"`
}

//
type RawTransactionGenParams struct {
	Inputs   []Input  `json:"inputs"`
	Outputs  []Output `json:"outputs"`
	Locktime int64    `json:"locktime"`
}
type RawTransactionGenResult struct {
	Rawtx string `json:"rawtx"`
}

//
type DecodeRawTransactionParams struct {
	Rawtx string `json:"rawtx"`
}
type Input struct {
	Txid string `json:"txid"`
	Vout uint32 `json:"vout"`
}
type Output struct {
	Address string  `json:"address"`
	Amount  float64 `json:"amount"`
}
type DecodeRawTransactionResult struct {
	Inputs   []Input  `json:"inputs"`
	Outputs  []Output `json:"outputs"`
	Locktime uint32   `json:"locktime"`
}

//
type GetTransactionByHashParams struct {
	TxHash string `json:"txhash"`
}
type GetTransactionByHashResult struct {
	Inputs  []Input  `json:"inputs"`
	Outputs []OutputIndex `json:"outputs"`
}

//
type SignTransactionParams struct {
	TransactionHex string   `json:"transactionhex"`
	RedeemHex      string   `json:"redeemhex"`
	Privkeys       []string `json:"privkeys"` //wif private keys
}
type SignTransactionResult struct {
	Complete       bool   `json:"complete"`
	TransactionHex string `json:"transactionhex"`
}

//
type SendTransactionParams struct {
	TransactionHex string `json:"transactionhex"`
}
type SendTransactionResult struct {
	TransactionHah string `json:"transactionhash"`
}

//
type SignTxSendParams struct {
	TransactionHex string   `json:"transactionhex"`
	RedeemHex      string   `json:"redeemhex"`
	Privkeys       []string `json:"privkeys"` //wif private keys
}
type SignTxSendResult struct {
	TransactionHah string `json:"transactionhash"`
	Complete       bool   `json:"complete"`
	TransactionHex string `json:"transactionhex"`
}

//
type GetBalanceParams struct {
	Address string `json:"address"`
	Minconf int    `json:"minconf"`
}
type GetBalanceResult struct {
	Value float64 `json:"value"`
}

//
type GetTransactionsParams struct {
	Account string `json:"account"`
	Count   int    `json:"count"`
	Skip    int    `json:"skip"`
}

type InputIndex struct {
	TxHash string `json:"txHash"`
	Index  uint32 `json:"index"`
	Addr   string `json:"addr"`
	Value  int64  `json:"value"`
}
type OutputIndex struct {
	Index uint32 `json:"index"`
	Addr  string `json:"addr"`
	Value int64  `json:"value"` //satoshi
}
type Transaction struct {
	TxHash        string   `json:"txHash"`
	BlanceChanged int64    `json:"blanceChanged"`
	Inputs        []InputIndex  `json:"inputs"`
	Outputs       []OutputIndex `json:"outputs"`
}
type TransactionsResult struct {
	Transactions []Transaction `json:"transactions"`
}

//
type ImportMultisigParams struct {
	PublicKeys   []string `json:"publicKeys"`
	MRequires    int      `json:"mRequires"`
	WalletPasswd string   `json:"walletPasswd"`
}
type ImportMultisigResult struct {
	Import bool `json:"import"`
}

/* not used current

type GetUTXOParams struct {
	Addresses    []string `json:"addresses"`
	Minconf      int      `json:"minconf"`
	Maxconf      int      `json:"maxconf"`
	MaximumCount int      `json:"maximumCount"`
}

func (abtc AdaptorBTC) GetUnspendUTXO(params string) string {
	return GetUnspendUTXO(params, &abtc.RPCParams, abtc.NetID)
}

*/
