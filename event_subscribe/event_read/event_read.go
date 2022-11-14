package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	store "go-ethereum-learn/event_read/contracts" // for demo
)

func main() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws/v3/6d46879db5974ac9a39ac6d37db50507")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x33400f7e34c846B9B5dE730860847A5b043AAB77")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(12713843),
		ToBlock:   big.NewInt(12713845),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(vLog.BlockNumber)     // 2394201
		fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6

		/*event := struct {
			Key   [32]byte
			Value [32]byte
		}{}*/
		event, err := contractAbi.Unpack("ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		bytes := event[0].([32]byte)
		fmt.Println(string(bytes[:]))
		bytes = event[1].([32]byte)
		fmt.Println(string(bytes[:]))

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println(topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
	}

	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex()) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
}
