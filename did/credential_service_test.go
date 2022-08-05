package did

import (
	"encoding/json"
	"fmt"
	platoncommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/bglmmz/chainclient"
	"github.com/datumtechs/did-sdk-go/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var didService *DIDService

var did = "did:pid:lat1x4w7852dxs69sy2mgf8w0s7tmvqx3cz2ydaxq4"
var doc_context = "http://datumtech.com/did/v1"
var pctId = new(big.Int).SetUint64(1)
var issuer = "did:pid:lat1x4w7852dxs69sy2mgf8w0s7tmvqx3cz2ydaxq4"

var privateKey, _ = crypto.HexToECDSA("b24285967575de7d5563e35213a806c60d69094faa509025f2ab5437017d343a")
var publicKey = hexutil.Encode(crypto.FromECDSAPub(&privateKey.PublicKey))
var address = platoncommon.Address(crypto.PubkeyToAddress(privateKey.PublicKey))
var publicKeyId string //= address.String() + "#keys-1", 需要首先初始化hrp

func setup() {
	fmt.Println("initing........")
	InitMockWallet()
	MockWalletInstance().SetPrivateKey(privateKey)
	ethcontext := chainclient.NewEthClientContext("ws://8.219.126.197:6790", "lat", MockWalletInstance())
	didService = NewDIDService(ethcontext)
	fmt.Println("publicKey:" + publicKey)
	fmt.Println("address:" + address.String())

	publicKeyId = address.String() + "#keys-1"
}

func Test_bech32(t *testing.T) {
	setup()
	addr := string(([]byte(did))[len("did:pid:"):])
	platonAddress, err := platoncommon.Bech32ToAddress(addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(platonAddress)
}

func Test_createVC(t *testing.T) {
	setup()
	claimVar := make(types.Claim)
	claimVar["age"] = "12"
	claimVar["name"] = "Alice"

	expirationDate := "1989-06-06-18T21:19:10"
	req := new(CreateCredentialReq)
	req.Did = did
	req.Context = doc_context
	req.PctId = pctId
	req.Claim = claimVar
	req.ExpirationDate = expirationDate
	req.Issuer = issuer
	req.PrivateKey = privateKey
	req.PublicKeyId = publicKeyId
	req.Type = types.CREDENTIAL_TYPE_VC
	response := didService.VcService.CreateCredentialSimple(*req)

	b, _ := json.Marshal(response.Data)

	t.Logf("response.Data:%s", string(b))

	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)

}

func Test_verifyVC(t *testing.T) {
	setup()

	claimVar := make(types.Claim)
	claimVar["age"] = "12"
	claimVar["name"] = "Alice"

	expirationDate := "1989-06-06-18T21:19:10"

	req := new(CreateCredentialReq)
	req.Did = did
	req.Context = doc_context
	req.PctId = pctId
	req.Claim = claimVar
	req.ExpirationDate = expirationDate
	req.Issuer = issuer
	req.PrivateKey = privateKey
	req.PublicKeyId = publicKeyId
	req.Type = types.CREDENTIAL_TYPE_VC

	response := didService.VcService.CreateCredentialSimple(*req)

	b, _ := json.Marshal(response.Data)

	t.Logf("VerifiableCredential:%s", string(b))

	t.Logf("response.Status:%d", response.Status)
	cred := response.Data
	t.Logf("%#v", cred)
	a := assert.New(t)
	if !a.Equal(Response_SUCCESS, response.Status) {
		t.Fatal()
	}

	/*ok, _ := didService.VcService.VerifyVC(&cred)
	a.Equal(true, ok)*/

	ok := didService.VcService.VerifyVCWithPublicKey(&cred, didService.VcService.ctx.GetPublicKey())
	a.Equal(true, ok)
}
