package types

import (
	"encoding/json"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/crypto"
	"github.com/datumtechs/did-sdk-go/keys/claim"
	log "github.com/sirupsen/logrus"
	"math/rand"
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

func (c Claim) GetSeed() uint64 {
	if c[claimkeys.SEED] == nil {
		return 0
	} else {
		switch value := c[claimkeys.SEED].(type) {
		case uint64:
			return value
		case string:
			seedString := c[claimkeys.SEED].(string)
			if seed, err := strconv.ParseUint(seedString, 10, 64); err != nil {
				log.Errorf("cannot parse seed, %s", seedString)
				return 0
			} else {
				return seed
			}
		default:
			return 0
		}
	}
}

//
func (c Claim) GetHash(seed uint64) string {
	newClaim := common.Clone(c)
	//遍历时只需要有效字段，因此需要删除可能存在的key
	delete(newClaim, claimkeys.SEED)
	delete(newClaim, claimkeys.ROOT_HASH)
	//对claim进行加盐，并计算ClaimRootHash
	if seed == 0 {
		seed = rand.Uint64()
		//fmt.Printf("generate new seed: %d\n", seed)
	}

	//为claim的有效key，生成新newValue:= json(original_value)+string(hash(seed))
	//并hash(newValue).hex(), 写入builder
	allNewValueHashesBuilder := strings.Builder{}
	GenerateClaimSaltForMap(newClaim, common.Uint64ToBigEndianBytes(seed), &allNewValueHashesBuilder)
	newClaim[claimkeys.SEED] = seed
	newClaim[claimkeys.ROOT_HASH] = crypto.SHA3Hex(allNewValueHashesBuilder.String())

	c[claimkeys.SEED] = seed
	c[claimkeys.ROOT_HASH] = crypto.SHA3Hex(allNewValueHashesBuilder.String())

	//json.Marshal会对key按字典顺序排列
	claimRawData, _ := json.Marshal(newClaim)
	//fmt.Printf("claimRawdata:%s\n", claimRawData)
	return crypto.SHA3Hex(string(claimRawData))
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
				builder.WriteString(crypto.SHA3Hex(newValue))
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
			builder.WriteString(crypto.SHA3Hex(newValue))
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

func SplitForMap(originalClaim, disclosureMap, undisclosed Claim, seed []byte) {
	var disclosedKeys []string
	for key := range disclosureMap {
		disclosedKeys = append(disclosedKeys, key)
	}

	//排序keys
	sort.Strings(disclosedKeys)
	for _, disclosedKey := range disclosedKeys {
		disclosedValue := disclosureMap[disclosedKey]
		originalValue := originalClaim[disclosedKey]

		if disclosureMapNextLevel, ok := disclosedValue.(map[string]interface{}); ok {
			if originalClaimNextLevel, ok := originalValue.(map[string]interface{}); ok {
				SplitForMap(originalClaimNextLevel, disclosureMapNextLevel, undisclosed, seed)
			} else {
				//todo: handle error:
			}
		} else if disclosureMapNextLevel, ok := disclosedValue.([]interface{}); ok {
			if originalClaimNextLevel, ok := originalValue.([]interface{}); ok {
				isMapOrList := SplitForList(originalClaimNextLevel, disclosureMapNextLevel, undisclosed, seed)
				if !isMapOrList {
					//替换value= json(value)+salt
					originalValueJson, _ := json.Marshal(originalValue)
					seed = common.GetHash(seed)
					newValue := string(originalValueJson) + strconv.FormatUint(common.BigEndianBytesToUint64(seed), 10)
					//fmt.Printf("claim key:%s newValue:=%s\n", key, newValue)

					disclosedValueJson, _ := json.Marshal(disclosedValue)
					if string(disclosedValueJson) == "0" { //不披露
						delete(originalClaim, disclosedKey)
						undisclosed[disclosedKey] = crypto.SHA3Hex(newValue)
					} else {
						delete(undisclosed, disclosedKey)
					}
				}
			} else {
				//todo: handle error:
			}

		} else {
			//替换value= json(value)+salt
			originalValueJson, _ := json.Marshal(originalValue)
			seed = common.GetHash(seed)
			newValue := string(originalValueJson) + strconv.FormatUint(common.BigEndianBytesToUint64(seed), 10)
			//fmt.Printf("claim key:%s newValue:=%s\n", key, newValue)

			disclosedValueJson, _ := json.Marshal(disclosedValue)
			if string(disclosedValueJson) == "0" { //不披露
				delete(originalClaim, disclosedKey)
				undisclosed[disclosedKey] = crypto.SHA3Hex(newValue)
			} else {
				delete(undisclosed, disclosedKey)
			}
		}
	}

}

func SplitForList(originalClaim []interface{}, disclosedClaim []interface{}, undisclosed Claim, seed []byte) bool {
	for idx, v := range disclosedClaim {
		if m, ok := v.(map[string]interface{}); ok {
			if original, ok := (originalClaim[idx]).(map[string]interface{}); ok {
				SplitForMap(original, m, undisclosed, seed)
			} else {
				//todo: handle error:
			}
		} else if l, ok := v.([]interface{}); ok {
			if original, ok := (originalClaim[idx]).([]interface{}); ok {
				SplitForList(original, l, undisclosed, seed)
			} else {
				//todo: handle error:
			}
		} else {
			return false
		}
	}
	return true
}
