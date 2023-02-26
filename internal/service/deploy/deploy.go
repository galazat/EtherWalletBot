package deploy

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	vote "github.com/galazat/EtherWalletBot/internal/vote"
)

func Deploy(client *ethclient.Client, ctx context.Context) {

	privateKey, err := crypto.HexToECDSA("26350ad0d6f61473246381d26c5b1cf17d79cc4194251561761d9454a42e13d8")
	if err != nil {
		log.Println(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Panicln("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Println(err, nonce)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Println(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce)) //big.NewInt(int64(21483089))
	auth.Value = big.NewInt(0)            // in wei
	auth.GasLimit = uint64(300000)        // in units
	auth.GasPrice = gasPrice              //big.NewInt(21000)        //gasPrice

	//input := bytes{'0x666f6f0000000000000000000000000000000000000000000000000000000000', 0x6261720000000000000000000000000000000000000000000000000000000000}
	var input [][32]byte
	var tm [32]byte
	input = append(input, tm)
	copy(input[0][:], "abc")
	address, _, instance, err := vote.DeployVote(auth, client, input)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(address.Hex()) // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54

	address2 := common.HexToAddress(address.Hex())
	instance, err2 := vote.NewVote(address2, client)
	if err != nil {
		log.Println(err2)
	}

	cand, err := instance.Candidats(nil, big.NewInt(0))
	if err != nil {
		log.Println(err)
	}

	fmt.Println(cand) // "1.0"

	fmt.Println("____________DEPLOYED_____________\n")
}
