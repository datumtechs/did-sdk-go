package did

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type ResponseStatus int

const (
	Response_FAILURE ResponseStatus = 0
	Response_SUCCESS ResponseStatus = 1
	Response_UNKNOWN ResponseStatus = 2
)

type TransactionInfo struct {
	BlockNumber      uint64
	TxHash           string
	TransactionIndex uint
}

func NewTransactionInfo(receipt *types.Receipt) TransactionInfo {
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
