package common

import (
	"encoding/json"
	"math/rand"
	"time"
	"unsafe"
)

const letterBytes = "012346789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func Clone[T any](src T) T {

	var dest T

	jsonStr, _ := json.Marshal(src)
	json.Unmarshal(jsonStr, &dest)
	return dest
}

/*
func CloneMap(src map[string]interface{}) map[string]interface{} {
	if src == nil || len(src) == 0 {
		return src
	}
	dest := make(map[string]interface{})

	jsonStr, _ := json.Marshal(src)
	json.Unmarshal(jsonStr, &dest)
	return dest
}


func CloneStruct(src interface{}) interface{} {
	if src == nil {
		return src
	}
	var dest interface{}

	jsonStr, _ := json.Marshal(src)
	json.Unmarshal(jsonStr, &dest)
	return dest
}*/
