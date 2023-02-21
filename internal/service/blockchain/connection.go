package blockchain

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var BLOCHCHAIN_NET_ADRR = "http://188.120.229.42:8545"

func ConnectionInit() (*ethclient.Client, error) {
	client, err := ethclient.Dial(BLOCHCHAIN_NET_ADRR)
	if err != nil {
		log.Fatal(err)
	}

	return client, err
}
