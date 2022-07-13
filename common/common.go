package common

import (
	"github.com/ethereum/go-ethereum/common"
	"time"
)

type ResponseStatus int

const (
	Response_FAILURE ResponseStatus = 0
	Response_SUCCESS ResponseStatus = 1
	Response_UNKNOWN ResponseStatus = 2
)

type Response struct {
	CallMode bool
	TxHash   common.Hash
	Status   ResponseStatus
	Msg      string
	Data     interface{}
}

func FormatUTC(utcTime time.Time) string {
	return utcTime.Format("2006-01-02T15:04:05.000")
}

func MustParseUTC(utcTime string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02T15:04:05.000", utcTime, time.UTC)
	return t
}
