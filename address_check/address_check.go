package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
	fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

	client, err := ethclient.Dial("xxx")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol Token (ZRX) smart contract address
	address := common.HexToAddress("0x712899f45f5c92117cbBBD937363dF38b2354853")
	bytecode, err := client.CodeAt(context.Background(), address, big.NewInt(81679)) // nil is lat est block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: true

	// a random user account address
	address = common.HexToAddress("0xFcD1aef048EaA60cD07076f27FC6E4C4c642BC01")
	bytecode, err = client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract = len(bytecode) > 0

	fmt.Printf("is   contract: %v\n", isContract) // is contract: false
}
