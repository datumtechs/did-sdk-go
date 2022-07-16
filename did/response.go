package did

import (
	"github.com/ethereum/go-ethereum/common"
)

type ResponseStatus int

const (
	Response_FAILURE ResponseStatus = 0
	Response_SUCCESS ResponseStatus = 1
	Response_UNKNOWN ResponseStatus = 2
)

type Response[K any] struct {
	CallMode bool
	TxHash   common.Hash
	Status   ResponseStatus
	Msg      string
	Data     K
}

func CopyResp[K, V any](from *Response[K], to *Response[V]) {
	to.Status = from.Status
	to.Msg = from.Msg
}
