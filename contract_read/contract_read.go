package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	store "go-ethereum-learn/contracts"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("xxx")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("私钥")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0xc592878BbF8F9d6133e954f3afe378E16eE9F110")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	_ = instance // we'll be using this in the 下个章节

	//api不一样
	/*version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"*/
}
