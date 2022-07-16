package claim

import (
	"encoding/json"
	"github.com/datumtechs/did-sdk-go/common"
	"github.com/datumtechs/did-sdk-go/crypto"
)

const (
	PCTID         string = "pctId"
	NOT_DISCLOSED int    = 0
	DISCLOSED     int    = 1
	EXISTED       int    = 2
)

type Claim map[string]string

/*type Claim map[string]interface{}

type ClaimMapSalt map[string]interface{}
type ClaimListSalt []interface{}

func (claimMapSalt *ClaimMapSalt) generateClaimMapSalt(fixedSalt string) {
	for k, v := range *claimMapSalt {
		if m, ok := v.(ClaimMapSalt); ok {
			m.generateClaimMapSalt(fixedSalt)
		} else if l, ok := v.(ClaimListSalt); ok {
			isMapOrList := l.generateClaimListSalt(fixedSalt)
			if !isMapOrList {
				if len(fixedSalt) == 0 {
					(map[string]interface{}(*claimMapSalt))[k] = RandStringBytesMaskImprSrcUnsafe(8)
				} else {
					(map[string]interface{}(*claimMapSalt))[k] = fixedSalt
				}
			}
		} else {
			//替换value=salt
			if len(fixedSalt) == 0 {
				(map[string]interface{}(*claimMapSalt))[k] = RandStringBytesMaskImprSrcUnsafe(8)
			} else {
				(map[string]interface{}(*claimMapSalt))[k] = fixedSalt
			}
		}
	}
}

func (claimListSalt *ClaimListSalt) generateClaimListSalt(fixedSalt string) bool {
	for _, v := range *claimListSalt {
		if m, ok := v.(ClaimMapSalt); ok {
			m.generateClaimMapSalt(fixedSalt)
		} else if l, ok := v.(ClaimListSalt); ok {
			isMapOrList := l.generateClaimListSalt(fixedSalt)
			if !isMapOrList {
				return isMapOrList
			}
		} else {
			return false
		}
	}
	return true
}*/

func (c Claim) GetHash(disclosures map[string]int) string {
	if disclosures == nil {
		disclosures = make(map[string]int)
	}
	if len(disclosures) == 0 {
		//每个字段都需要披露
		for key, _ := range c {
			disclosures[key] = DISCLOSED
		}
	}

	dest := common.Clone(c)
	//对要披露的claim值进行hash
	for key, _ := range disclosures {
		dest[key] = crypto.SHA3Hex(dest[key])
	}

	cliamRawdata, _ := json.Marshal(dest)
	return crypto.SHA3Hex(string(cliamRawdata))
}

func GenerateClaimSaltForMap(claimMapSalt map[string]interface{}, fixedSalt string) {
	for k, v := range claimMapSalt {
		if m, ok := v.(map[string]interface{}); ok {
			GenerateClaimSaltForMap(m, fixedSalt)
		} else if l, ok := v.([]interface{}); ok {
			isMapOrList := GenerateClaimSaltForList(l, fixedSalt)
			if !isMapOrList {
				if len(fixedSalt) == 0 {
					claimMapSalt[k] = common.RandStringUnsafe(8)
				} else {
					claimMapSalt[k] = fixedSalt
				}
			}
		} else {
			//替换value=salt
			if len(fixedSalt) == 0 {
				claimMapSalt[k] = common.RandStringUnsafe(8)
			} else {
				claimMapSalt[k] = fixedSalt
			}
		}
	}
}

func GenerateClaimSaltForList(claimListSalt []interface{}, fixedSalt string) bool {
	for _, v := range claimListSalt {
		if m, ok := v.(map[string]interface{}); ok {
			GenerateClaimSaltForMap(m, fixedSalt)
		} else if l, ok := v.([]interface{}); ok {
			isMapOrList := GenerateClaimSaltForList(l, fixedSalt)
			if !isMapOrList {
				return isMapOrList
			}
		} else {
			return false
		}
	}
	return true
}
