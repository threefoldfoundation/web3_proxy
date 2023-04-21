package goethclient

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateKeypair() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

func KeyFromSecret(secret string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(secret)
}

func (c *Client) AddressFromKey() common.Address {
	publicKey := c.Key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA)
}
