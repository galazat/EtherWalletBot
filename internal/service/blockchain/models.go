package blockchain

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

type Account struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	Address    common.Address
}
