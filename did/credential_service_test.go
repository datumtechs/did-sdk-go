package did

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	platoncommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/bglmmz/chainclient"
	"github.com/datumtechs/did-sdk-go/crypto"
	"github.com/datumtechs/did-sdk-go/types"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethhexutil "github.com/ethereum/go-ethereum/common/hexutil"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var didService *DIDService

var doc_context = "http://datumtech.com/did/v1"
var pctId = new(big.Int).SetUint64(1000)

var privateKey, _ = ethcrypto.HexToECDSA("68efa6466edaed4918f0b6c3b1b9667d37cad591482d672e8abcb4c5d1720f89")
var publicKey = ethhexutil.Encode(ethcrypto.FromECDSAPub(&privateKey.PublicKey))

var bech32Addr = platoncommon.Address(ethcrypto.PubkeyToAddress(privateKey.PublicKey))
var address = ethcrypto.PubkeyToAddress(privateKey.PublicKey)
var issuer = "did:pid:lat1d7zjh2vx8xsqrgc4qe0v4usxn368naxvlpu70r"
var did = "did:pid:lat1d7zjh2vx8xsqrgc4qe0v4usxn368naxvlpu70r"
var publicKeyId string = did + "#keys-1" //, 需要首先初始化hrp

var applicantPriKey, _ = ethcrypto.HexToECDSA("42fe5edfa8327ceb7fe8ae059251f38528fd3bf8d65e26619bafcf60849790ec")
var applicantPublicKey = ethhexutil.Encode(ethcrypto.FromECDSAPub(&applicantPriKey.PublicKey))
var applicantBech32Addr = platoncommon.Address(ethcrypto.PubkeyToAddress(applicantPriKey.PublicKey))
var applicantDid string = "did:pid:lat1cq9svdd8vc83u74relncn6cyxywr5mjqccqlea"
var applicantPublicKeyId = applicantDid + "#keys-1" //= bech32Addr.String() + "#keys-1", 需要首先初始化hrp

var credential types.Credential
var vc = "{\"context\":\"http://datumtech.com/did/v1\",\"version\":\"1.0.0\",\"id\":\"a4a47370-a75b-41fb-bd87-0c823a72be07\",\"type\":[\"VerifiableCredential\"],\"issuer\":\"did:pid:lat1d7zjh2vx8xsqrgc4qe0v4usxn368naxvlpu70r\",\"issuanceDate\":\"2022-08-09T04:06:26.901\",\"expirationDate\":\"2029-07-06-18T21:19:10\",\"claimData\":{\"nodeID\":\"0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\",\"nodeName\":\"The PlatON Node\",\"url\":\"http://www.platon.network\"},\"claimMeta\":{\"pctId\":\"1000\"},\"proof\":{\"claimRootHash\":\"0x7d17357718f54ff5df50b087f6e95adae5e20cb54365b11c0c332667f1bc7eca\",\"created\":\"2022-08-09T04:06:26.901\",\"jws\":\"0xe3a62b3a0aad740e2f8ae693a049b1f9660936c416d3037bb98ce13c61b232f849f3061d294d5276c56126dee8c19febdea6f467b452df09bdbf92644576780200\",\"seed\":\"9828766684487745566\",\"type\":\"Secp256k1\",\"verificationMethod\":\"did:pid:lat1d7zjh2vx8xsqrgc4qe0v4usxn368naxvlpu70r#keys-1\"},\"holder\":\"did:pid:lat1cq9svdd8vc83u74relncn6cyxywr5mjqccqlea\"}"
var credentialHash ethcommon.Hash

func setup() {
	fmt.Println("initing........")
	InitMockWallet()
	MockWalletInstance().SetPrivateKey(privateKey)
	ethcontext := chainclient.NewEthClientContext("ws://8.219.126.197:6790", "lat", MockWalletInstance())
	didService = NewDIDService(ethcontext)
	fmt.Println("publicKey:" + publicKey)
	fmt.Println("bech32Addr:" + bech32Addr.String())
	//publicKeyId = did + "#keys-1"
	fmt.Println("publicKeyId:" + publicKeyId)

	err := json.Unmarshal([]byte(vc), &credential)
	if err != nil {
		panic(err)
	}

	seed, err := credential.Proof.GetSeed()
	if err != nil {
		panic(err)
	}

	digest, _ := credential.GetDigest(seed)
	credentialHash = ethcommon.BytesToHash(digest)
}

func Test_CreateCredentialSimple(t *testing.T) {
	setup()
	claimVar := make(types.Claim)
	claimVar["nodeID"] = "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	claimVar["nodeName"] = "The PlatON Node"
	claimVar["url"] = "http://www.platon.network"
	expirationDate := "2019-06-06-18T21:19:10"
	req := new(CreateCredentialReq)
	req.Did = applicantDid
	req.Context = doc_context
	req.PctId = pctId
	req.Claim = claimVar
	req.ExpirationDate = expirationDate
	req.Issuer = issuer
	req.PrivateKey = privateKey
	req.PublicKeyId = publicKeyId
	req.Type = types.CREDENTIAL_TYPE_VC
	response := didService.CredentialService.CreateCredentialSimple(*req)

	/*b, _ := json.Marshal(response.Data)*/

	t.Logf("response.Data:%+v", *response)

	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)

}

func Test_VerifySimpleVC(t *testing.T) {
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

	response := didService.CredentialService.CreateCredentialSimple(*req)

	b, _ := json.Marshal(response.Data)

	t.Logf("VerifiableCredential:%s", string(b))

	t.Logf("response.Status:%d", response.Status)
	cred := response.Data
	t.Logf("%#v", cred)
	a := assert.New(t)
	if !a.Equal(Response_SUCCESS, response.Status) {
		t.Fatal()
	}

	/*ok, _ := didService.CredentialService.VerifyCredential(&cred)
	a.Equal(true, ok)*/

	ok := didService.CredentialService.VerifyCredentialWithPublicKey(&cred, didService.CredentialService.ctx.GetPublicKey())
	a.Equal(true, ok)
}

func Test_CreateCredential(t *testing.T) {
	setup()
	claimVar := make(types.Claim)
	claimVar["nodeID"] = "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	claimVar["nodeName"] = "The PlatON Node"
	claimVar["url"] = "http://www.platon.network"
	expirationDate := "2029-06-06-18T21:19:10"
	req := new(CreateCredentialReq)
	req.Did = applicantDid
	req.Context = doc_context
	req.PctId = pctId
	req.Claim = claimVar
	req.ExpirationDate = expirationDate
	req.Issuer = issuer
	req.PrivateKey = privateKey
	req.PublicKeyId = publicKeyId
	req.Type = types.CREDENTIAL_TYPE_VC
	response := didService.CredentialService.CreateCredentialSimple(*req)

	/*b, _ := json.Marshal(response.Data)*/

	t.Logf("response.Data:%+v", *response)

	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)

}

func Test_CreateCredential2(t *testing.T) {
	setup()
	claimVar := make(types.Claim)
	claimVar["nodeID"] = "did:pid:lat1wdv3hh6auk0um7yr6lsxu8rwljk8942uxezekr"
	claimVar["nodeName"] = "org_9_156"
	claimVar["url"] = "ipfs://QmdJTKxgiVjKd4NNwhA2jgRS2JJCJ2iSxLSW7Lidc3YnGv"
	expirationDate := "2122-08-22T07:56:47.061"
	req := new(CreateCredentialReq)
	req.Did = "did:pid:lat1wdv3hh6auk0um7yr6lsxu8rwljk8942uxezekr"
	req.Context = doc_context
	req.PctId = pctId
	req.Claim = claimVar
	req.ExpirationDate = expirationDate
	req.Issuer = issuer
	req.PrivateKey = privateKey
	req.PublicKeyId = publicKeyId
	req.Type = types.CREDENTIAL_TYPE_VC
	response := didService.CredentialService.CreateCredential(*req)

	/*b, _ := json.Marshal(response.Data)*/

	t.Logf("response.Data:%+v", *response)

	a := assert.New(t)
	a.Equal(Response_SUCCESS, response.Status)

}

func Test_CreateAndVerifyVC(t *testing.T) {
	setup()
	t.Helper()
	claimVar := make(types.Claim)
	claimVar["nodeID"] = "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	claimVar["nodeName"] = "The PlatON Node"
	claimVar["url"] = "http://www.platon.network"
	expirationDate := "2029-07-06-18T21:19:10"

	req := new(CreateCredentialReq)
	req.Did = applicantDid
	req.Context = doc_context
	req.PctId = pctId
	req.Claim = claimVar
	req.ExpirationDate = expirationDate
	req.Issuer = issuer
	req.PrivateKey = privateKey
	req.PublicKeyId = publicKeyId
	req.Type = types.CREDENTIAL_TYPE_VC

	response := didService.CredentialService.CreateCredential(*req)

	b, _ := json.Marshal(response.Data)

	t.Logf("VerifiableCredential:%s", string(b))

	t.Logf("response.Status:%d", response.Status)
	cred := response.Data
	t.Logf("%#v", cred)
	a := assert.New(t)
	if !a.Equal(Response_SUCCESS, response.Status) {
		t.Fatal()
	}

	/*ok, _ := didService.CredentialService.VerifyCredential(&cred)
	a.Equal(true, ok)*/

	ok := didService.CredentialService.VerifyCredential(&cred)
	a.Equal(Response_SUCCESS, ok.Status)
	a.Equal(true, ok.Data)
}

func Test_VerifyVC(t *testing.T) {
	setup()
	t.Helper()
	ok := didService.CredentialService.VerifyCredential(&credential)
	a := assert.New(t)
	a.Equal(Response_SUCCESS, ok.Status)
}

func Test_VerifySecp256k1Signature(t *testing.T) {
	setup()
	t.Helper()
	seed, _ := credential.Proof.GetSeed()
	credentialHash, _ := credential.GetDigest(seed)

	pubKey := ethcrypto.FromECDSAPub(&privateKey.PublicKey)
	t.Logf("publicKey:%s", ethhexutil.Encode(pubKey))

	ok := crypto.VerifySecp256k1Signature(credentialHash, "0xe3a62b3a0aad740e2f8ae693a049b1f9660936c416d3037bb98ce13c61b232f849f3061d294d5276c56126dee8c19febdea6f467b452df09bdbf92644576780200", &privateKey.PublicKey)
	t.Logf("ok:%t", ok)
}

func Test_VerifySecp256k1Signature2(t *testing.T) {
	plain := "test"
	reqHash := HashSHA256([]byte(plain))
	sig := crypto.SignSecp256k1(reqHash, privateKey)
	t.Logf("reqHash:%s", ethhexutil.Encode(reqHash))
	t.Logf("sig:%s", ethhexutil.Encode(sig))

	ok := crypto.VerifySecp256k1Signature(reqHash, ethhexutil.Encode(sig), &privateKey.PublicKey)
	t.Logf("ok:%t", ok)

}

func HashSHA256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}
