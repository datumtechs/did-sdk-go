package types

import "github.com/datumtechs/did-sdk-go/common"

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
