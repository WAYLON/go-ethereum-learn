package main

import (
	"context"
	"fmt"
	"go-ethereum-learn/contracts"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// LogTransfer ..
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

func main() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws/v3/6d46879db5974ac9a39ac6d37db50507")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress("0xD44983eaA84f1b94B8D0d13C8558f20a6BE5ef35")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(12714196),
		ToBlock:   big.NewInt(12714239),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.TokenABI)))
	if err != nil {
		log.Fatal(err)
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
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
