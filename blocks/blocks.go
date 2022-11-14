package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-ethereum-learn/contracts"
	"log"
	"math/big"
	"strings"
)

func main() {
	client, err := ethclient.Dial("https://geth.mm.comeonbtc.com:8443")
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String()) // 5671744

	blockNumber := header.Number
	_ = blockNumber
	block, err := client.BlockByNumber(context.Background(), big.NewInt(238425))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions()))   // 144
	contractAbi, err := abi.JSON(strings.NewReader(contracts.ERC1155MetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}
	for _, transaction := range block.Transactions() {
		receipt, err := client.TransactionReceipt(context.Background(), transaction.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(receipt.Status) // 1
		fmt.Println(receipt.Logs)   // .

		logTransferSig := []byte("Transfer(address,address,uint256)")
		LogApprovalSig := []byte("Approval(address,address,uint256)")
		logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
		logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

		for _, vLog := range receipt.Logs {
			fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
			fmt.Printf("Log Index: %d\n", vLog.Index)

			switch vLog.Topics[0].Hex() {
			case logTransferSigHash.Hex():
				fmt.Printf("Log Name: Transfer\n")

				transferEvent, err := contractAbi.Unpack("Transfer", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(transferEvent)
				fmt.Println(common.HexToAddress(vLog.Topics[1].Hex()))
				fmt.Println(common.HexToAddress(vLog.Topics[2].Hex()))

			case logApprovalSigHash.Hex():
				fmt.Printf("Log Name: Approval\n")

				logApproval, err := contractAbi.Unpack("Approval", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(common.HexToAddress(vLog.Topics[1].Hex()))
				fmt.Println(common.HexToAddress(vLog.Topics[2].Hex()))
				fmt.Println(logApproval)
			}

			fmt.Printf("\n\n")
		}

	}

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 144
}
