package did

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"sync"
)

var (
	mockWalletOnce sync.Once
)

var mockWallet *MockWallet

type MockWallet struct {
	priKey        *ecdsa.PrivateKey
	pubKey        *ecdsa.PublicKey
	walletAddress common.Address
}

func MockWalletInstance() *MockWallet {
	return mockWallet
}

func InitMockWallet() {
	mockWalletOnce.Do(func() {
		mockWallet = new(MockWallet)

		key, _ := crypto.GenerateKey()
		mockWallet.priKey = key
		mockWallet.pubKey = &key.PublicKey
		mockWallet.walletAddress = crypto.PubkeyToAddress(key.PublicKey)
	})
}

// GetAddress returns the organization wallet address
func (m *MockWallet) GetAddress() common.Address {
	return m.walletAddress
}

// GetPrivateKey returns the organization private key
func (m *MockWallet) GetPrivateKey() *ecdsa.PrivateKey {
	return m.priKey
}

// GetPrivateKey returns the organization private key
func (m *MockWallet) GetPublicKey() *ecdsa.PublicKey {
	return m.pubKey
}
