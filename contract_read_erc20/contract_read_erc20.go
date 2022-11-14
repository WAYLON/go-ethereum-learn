package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	store "go-ethereum-learn/contracts"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	tokenAddress := common.HexToAddress("0xa4E9f795042D677eEe61E70aeE9e9f42a8EBF60B")
	fmt.Println(string(tokenAddress.Bytes()))
	client, err := ethclient.Dial("xxx")
	if err != nil {
		log.Fatal(err)
	}

	instance, err := store.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	fromAddress := common.HexToAddress("0xFcD1aef048EaA60cD07076f27FC6E4C4c642BC01")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)
	fmt.Printf("symbol: %s\n", symbol)
	fmt.Printf("decimals: %v\n", decimals)

	fmt.Printf("wei: %s\n", bal)

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f\n", value)
	toAddress := common.HexToAddress("0x4D96972a599FD930434DDF3712583512C18b3F21")

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("私钥")

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	transfer, err := instance.Transfer(auth, toAddress, big.NewInt(1000000000000000000))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(transfer)

	of, err := instance.BalanceOf(nil, fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(of)

	to, err := instance.BalanceOf(nil, toAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(to)

	supply, err := instance.TotalSupply(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(supply)

	approve, err := instance.Approve(auth, toAddress, big.NewInt(1000000000000000000))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(approve)

}
