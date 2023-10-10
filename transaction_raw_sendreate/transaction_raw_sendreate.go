package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	client, err := ethclient.Dial("https://geth.inner.comeonbtc.com")
	if err != nil {
		log.Fatal(err)
	}

	rawTx := "f86e038459682f07825208944d96972a599fd930434ddf3712583512c18b3f21880de0b6b3a7640000808306d104a084cbed759da1e60939fa7f21ed40583d3034d53e81faf9594b1a42b8d470041ea003620d2570bc51bc928c82c09348af0b5b0ceb931ed92578181a674078ace48d"

	rawTxBytes, err := hex.DecodeString(rawTx)

	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}
