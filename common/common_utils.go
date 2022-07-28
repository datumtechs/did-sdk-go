package common

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
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

func FormatUTC(utcTime time.Time) string {
	return utcTime.Format("2006-01-02T15:04:05.000")
}

func MustParseUTC(utcTime string) time.Time {
	t, _ := time.ParseInLocation("2006-01-02T15:04:05.000", utcTime, time.UTC)
	return t
}

func GenerateSequence256(seed []byte, count int) [][]byte {
	result := [][]byte{}
	current := seed
	for i := 0; i < count; i++ {
		current = GetHash(current)
		result = append(result, current)
	}
	return result
}

func GetHash(seed []byte) []byte {
	h := sha256.New()
	h.Write(seed)
	return h.Sum(nil)
}

func Uint64ToBigEndianBytes(seed uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, seed)
	return b
}

func BigEndianBytesToUint64(seed []byte) uint64 {
	return binary.BigEndian.Uint64(seed)
}

func VerifyWithJsonSchema(jsonSchema string, content map[string]interface{}) bool {
	contentJson, err := json.Marshal(content)
	if err != nil {
		log.WithError(err).Errorf("cannot marshal content to json, content: %+v", content)
		return false
	}

	schemaLoader := gojsonschema.NewStringLoader(jsonSchema)
	documentLoader := gojsonschema.NewStringLoader(string(contentJson))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		log.WithError(err).Error("failed to validate content")
		return false
	}

	if !result.Valid() {
		log.WithError(err).Errorf("content is not valid: %+v", result.Errors())
		return false
	}
	return true
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
