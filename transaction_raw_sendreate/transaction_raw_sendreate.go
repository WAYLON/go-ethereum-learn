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
	client, err := ethclient.Dial("xxx")
	if err != nil {
		log.Fatal(err)
	}

	rawTx := "f86e0d843b9aca07825208944d96972a599fd930434ddf3712583512c18b3f21880de0b6b3a7640000808306d104a05c831b3d01211018c10a2966f5d5b82519b70103a6f015b224ddbbfef70bbbefa01fbdd2e6cefe45535ed99403706bd8e62118a2550b8632c468a3dbdbd93a372e\n"

	rawTxBytes, err := hex.DecodeString(rawTx)

	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}
