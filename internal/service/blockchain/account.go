package blockchain

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func InitAccount() *Account {
	return &Account{}
}

func (a *Account) GetPrivateKey() string {
	privateKeyBytes := crypto.FromECDSA(a.privateKey)
	return hexutil.Encode(privateKeyBytes)[2:]
}

func (a *Account) GetPublicKey() string {
	publicKeyBytes := crypto.FromECDSAPub(a.publicKey)
	return hexutil.Encode(publicKeyBytes)[4:]
}

func (a *Account) GetAddress() string {
	return a.Address.Hex()
}

func (a *Account) SetPrivateKey(pk *ecdsa.PrivateKey) {
	a.privateKey = pk
}

func (a *Account) PublicFromPrivate() {
	public := a.privateKey.Public()
	a.publicKey = public.(*ecdsa.PublicKey)
}

func (a *Account) AddressFromPublic() {
	a.Address = crypto.PubkeyToAddress(*a.publicKey)
}
