package did

import (
	"crypto/ecdsa"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"sync"
)

var (
	mockWalletOnce sync.Once
)

var mockWallet *MockWallet

type MockWallet struct {
	priKey        *ecdsa.PrivateKey
	pubKey        *ecdsa.PublicKey
	walletAddress ethcommon.Address
}

func MockWalletInstance() *MockWallet {
	return mockWallet
}

func InitMockWallet() {
	mockWalletOnce.Do(func() {
		mockWallet = new(MockWallet)

		key, _ := ethcrypto.GenerateKey()
		mockWallet.priKey = key
		mockWallet.pubKey = &key.PublicKey
		mockWallet.walletAddress = ethcrypto.PubkeyToAddress(key.PublicKey)
	})
}

// GetAddress returns the organization wallet bech32Addr
func (m *MockWallet) GetAddress() ethcommon.Address {
	return m.walletAddress
}

// GetPrivateKey returns the organization private key
func (m *MockWallet) SetPrivateKey(privateKey *ecdsa.PrivateKey) {
	m.priKey = privateKey
	m.pubKey = &privateKey.PublicKey
	m.walletAddress = ethcrypto.PubkeyToAddress(privateKey.PublicKey)
}

// GetPrivateKey returns the organization private key
func (m *MockWallet) GetPrivateKey() *ecdsa.PrivateKey {
	return m.priKey
}

// GetPrivateKey returns the organization private key
func (m *MockWallet) GetPublicKey() *ecdsa.PublicKey {
	return m.pubKey
}
