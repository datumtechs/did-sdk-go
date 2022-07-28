package claim

import (
	"encoding/json"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/crypto"
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

type ClaimKey string

const (
	SEED      ClaimKey = "seed"
	ROOT_HASH ClaimKey = "rootHash"
)

type Claim map[string]interface{}

func (c Claim) GetSeed() uint64 {
	if c[string(SEED)] == nil {
		return 0
	} else {
		seedString := c[string(SEED)].(string)
		if seed, err := strconv.ParseUint(seedString, 10, 64); err != nil {
			log.Errorf("cannot parse seed, %s", seedString)
			return 0
		} else {
			return seed
		}
	}
}

func (c Claim) GetHash(disclosures map[string]int, seed uint64) string {
	if disclosures == nil {
		disclosures = make(map[string]int)
	}
	if len(disclosures) == 0 {
		//每个字段都需要披露
		for key, _ := range c {
			disclosures[key] = int(DISCLOSED)
		}
	}

	//对claim进行加盐，并计算ClaimRootHash
	if seed == 0 {
		seed = rand.Uint64()
		c[string(SEED)] = strconv.FormatUint(seed, 10)
		//fmt.Printf("generate new seed: %d\n", seed)
	}

	//为claim的有效字段，生成新的值: json(original_value)+string(seedHash)
	dest := common.Clone(c)
	//遍历时只需要有效字段，因此需要删除可能存在的key
	delete(dest, string(SEED))
	delete(dest, string(ROOT_HASH))

	hashStringBuilder := strings.Builder{}
	GenerateClaimSaltForMap(dest, common.Uint64ToBigEndianBytes(seed), &hashStringBuilder)
	claimRootHash := crypto.SHA3Hex(hashStringBuilder.String())
	c[string(ROOT_HASH)] = claimRootHash

	//json.Marshal会对key按字典顺序排列
	claimRawdata, _ := json.Marshal(c)
	//fmt.Printf("claimRawdata:%s\n", claimRawdata)
	return crypto.SHA3Hex(string(claimRawdata))

}

func GenerateClaimSaltForMap(claimMapSalt map[string]interface{}, seed []byte, builder *strings.Builder) {
	var keys []string
	for key := range claimMapSalt {
		keys = append(keys, key)
	}
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
