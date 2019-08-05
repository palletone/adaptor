/*
 *
 *    This file is part of go-palletone.
 *    go-palletone is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *    go-palletone is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *    You should have received a copy of the GNU General Public License
 *    along with go-palletone.  If not, see <http://www.gnu.org/licenses/>.
 * /
 *
 *  * @author PalletOne core developer <dev@pallet.one>
 *  * @date 2018-2019
 *
 */

package adaptor

type TxBasicInfo struct {
	TxID           []byte `json:"tx_id"`           //交易的ID，Hash
	TxRawData      []byte `json:"tx_raw"`          //交易的二进制数据
	CreatorAddress string `json:"creator_address"` //交易的发起人
	TargetAddress  string `json:"target_address"`  //交易的目标地址（被调用的合约、收款人）
	IsInBlock      bool   `json:"is_in_block"`     //是否已经被打包到区块链中
	IsSuccess      bool   `json:"is_success"`      //是被标记为成功执行
	IsStable       bool   `json:"is_stable"`       //是否已经稳定不可逆
	BlockID        []byte `json:"block_id"`        //交易被打包到了哪个区块ID
	BlockHeight    uint   `json:"block_height"`    //交易被打包到的区块的高度
	TxIndex        uint   `json:"tx_index"`        //Tx在区块中的位置
	Timestamp      uint64 `json:"timestamp"`       //交易被打包的时间戳
}
