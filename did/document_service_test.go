package did

import (
	"encoding/hex"
	"github.com/datumtechs/did-sdk-go/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_createDid(t *testing.T) {
	setup()
	req := CreateDidReq{}
	req.PrivateKey = MockWalletInstance().GetPrivateKey()
	req.PublicKey = hex.EncodeToString(crypto.FromECDSAPub(&MockWalletInstance().GetPrivateKey().PublicKey))
	req.PublicKeyType = types.PublicKey_SECP256K1
	response := didService.DocumentService.CreateDID(req)

	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)

	t.Logf("response:%+v", response)

}
