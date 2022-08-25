package types

import (
	"encoding/json"
	"errors"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/crypto"
	"sort"
	"strconv"
	"strings"
)

const (
	NOT_DISCLOSED int = 0
	DISCLOSED     int = 1
	EXISTED       int = 2
)

type Claim map[string]interface{}

//
func (c Claim) GetHash(seed uint64) (claimDigest string, rootHash string) {
	newClaim := common.Clone(c)
	//对claim进行加盐，并计算ClaimRootHash
	//为claim的有效key，生成新newValue:= json(original_value)+string(hash(seed))
	//并hash(newValue).hex(), 写入builder
	allNewValueHashesBuilder := strings.Builder{}
	GenerateClaimSaltForMap(newClaim, common.Uint64ToBigEndianBytes(seed), &allNewValueHashesBuilder)

	//json.Marshal会对key按字典顺序排列
	claimRawData, _ := json.Marshal(newClaim)
	//fmt.Printf("claimRawdata:%s\n", claimRawData)
	return crypto.LegacyKeccak256SHA3Hex(string(claimRawData)), crypto.LegacyKeccak256SHA3Hex(allNewValueHashesBuilder.String())
}

func GenerateClaimSaltForMap(claimMapSalt map[string]interface{}, seed []byte, builder *strings.Builder) {
	var keys []string
	for key := range claimMapSalt {
		keys = append(keys, key)
	}
	//排序keys
	sort.Strings(keys)
	for _, key := range keys {
		v := claimMapSalt[key]
		if m, ok := v.(map[string]interface{}); ok {
			GenerateClaimSaltForMap(m, seed, builder)
		} else if l, ok := v.([]interface{}); ok {
			isMapOrList := GenerateClaimSaltForList(l, seed, builder)
			if !isMapOrList {
				//替换value= json(value)+hash(seed)
				vJson, _ := json.Marshal(v)
				seed = common.GetHash(seed)
				newValue := string(vJson) + strconv.FormatUint(common.BigEndianBytesToUint64(seed), 10)
				claimMapSalt[key] = newValue
				builder.WriteString(crypto.LegacyKeccak256SHA3Hex(newValue))
				builder.WriteString(" ")
			}
		} else {
			//替换value= json(value)+salt
			//vJson, _ := json.Marshal(v)
			vJson, _ := json.Marshal(v)
			seed = common.GetHash(seed)
			newValue := string(vJson) + strconv.FormatUint(common.BigEndianBytesToUint64(seed), 10)
			//fmt.Printf("claim key:%s newValue:=%s\n", key, newValue)
			claimMapSalt[key] = newValue
			builder.WriteString(crypto.LegacyKeccak256SHA3Hex(newValue))
			builder.WriteString(" ")
		}
	}
}

func GenerateClaimSaltForList(claimListSalt []interface{}, seed []byte, builder *strings.Builder) bool {
	for _, v := range claimListSalt {
		if m, ok := v.(map[string]interface{}); ok {
			GenerateClaimSaltForMap(m, seed, builder)
		} else if l, ok := v.([]interface{}); ok {
			isMapOrList := GenerateClaimSaltForList(l, seed, builder)
			if !isMapOrList {
				return isMapOrList
			}
		} else {
			return false
		}
	}
	return true
}

func getType(i interface{}) string {
	switch i.(type) {
	case map[string]interface{}:
		return "map"
	case []interface{}:
		return "list"
	}
	return "final"
}

func SplitForMap(originalClaim, disclosureMap Claim, seed []byte) error {
	var originalKeys []string
	for key := range originalClaim {
		originalKeys = append(originalKeys, key)
	}

	//排序keys
	sort.Strings(originalKeys)
	for _, key := range originalKeys {
		originalValue := originalClaim[key]
		disclosedValue := disclosureMap[key]

		originalType := getType(originalValue)
		disclosedType := getType(disclosedValue)

		if originalType == "map" && disclosedType == "map" {
			SplitForMap(originalValue.(map[string]interface{}), disclosedValue.(map[string]interface{}), seed)
		} else if disclosedType == "list" {
			isMapOrList, err := SplitForList(originalValue.([]interface{}), disclosedValue.([]interface{}), seed)
			if err != nil {
				return err
			}
			if !isMapOrList {
				//替换value= json(value)+salt
				originalValueJson, _ := json.Marshal(originalValue)
				seed = common.GetHash(seed)
				newValue := string(originalValueJson) + strconv.FormatUint(common.BigEndianBytesToUint64(seed), 10)
				//fmt.Printf("claim key:%s newValue:=%s\n", key, newValue)

				disclosedValueJson, _ := json.Marshal(disclosedValue)
				if string(disclosedValueJson) == "0" { //不披露
					originalClaim[key] = crypto.LegacyKeccak256SHA3Hex(newValue)
				}
			}
		} else {
			//替换value= json(value)+salt
			originalValueJson, _ := json.Marshal(originalValue)
			seed = common.GetHash(seed)
			newValue := string(originalValueJson) + strconv.FormatUint(common.BigEndianBytesToUint64(seed), 10)
			//fmt.Printf("claim key:%s newValue:=%s\n", key, newValue)

			disclosedValueJson, _ := json.Marshal(disclosedValue)
			if string(disclosedValueJson) == "0" { //不披露
				originalClaim[key] = crypto.LegacyKeccak256SHA3Hex(newValue)
			}
		}
	}
	return nil
}

func SplitForList(originalClaim []interface{}, disclosedClaim []interface{}, seed []byte) (bool, error) {
	if len(originalClaim) == len(disclosedClaim) {
		for idx := 0; idx < len(originalClaim); idx++ {
			originalValue := originalClaim[idx]
			disclosedValue := disclosedClaim[idx]

			originalType := getType(originalValue)
			disclosedType := getType(disclosedValue)

			if originalType == "map" && disclosedType == "map" {
				SplitForMap(originalValue.(map[string]interface{}), disclosedValue.(map[string]interface{}), seed)
			} else if originalType == "list" && disclosedType == "list" {
				isMapOrList, err := SplitForList(originalValue.([]interface{}), disclosedValue.([]interface{}), seed)
				if err != nil {
					return isMapOrList, err
				}
				if !isMapOrList {
					return isMapOrList, nil
				}
			} else {
				return false, nil
			}
		}
	} else {
		return false, errors.New("claim list error")
	}
	return true, nil
}
