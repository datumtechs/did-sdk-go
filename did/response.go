package did

import (
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type ResponseStatus int

const (
	Response_FAILURE      ResponseStatus = 0
	Response_SUCCESS      ResponseStatus = 1
	Response_EXIST        ResponseStatus = 2
	Response_NOT_FOUND    ResponseStatus = 3
	Response_DEACTIVATION ResponseStatus = 3
	Response_UNKNOWN      ResponseStatus = 4
)

type TransactionInfo struct {
	BlockNumber      uint64
	TxHash           string
	TransactionIndex uint
}

func NewTransactionInfo(receipt *ethtypes.Receipt) TransactionInfo {
	return TransactionInfo{
		BlockNumber:      receipt.BlockNumber.Uint64(),
		TxHash:           receipt.TxHash.Hex(),
		TransactionIndex: receipt.TransactionIndex,
	}
}

type Response[K any] struct {
	CallMode bool
	TxInfo   TransactionInfo
	Status   ResponseStatus
	Msg      string
	Data     K
}

func CopyResp[K, V any](from *Response[K], to *Response[V]) {
	to.Status = from.Status
	to.Msg = from.Msg
}
