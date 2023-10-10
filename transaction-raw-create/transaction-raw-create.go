package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://geth.inner.comeonbtc.com")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		log.Fatal(err)
	}

	fromAddress := common.HexToAddress("0x12d5E2A37c7814BF53C2231669ba0c426A79ce2b")

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000000) // in wei (1 eth)

	gasLimit := uint64(21000) // in units

	gasPrice := big.NewInt(30000000000) // in wei (30 gwei)

	gasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4D96972a599FD930434DDF3712583512C18b3F21")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

	h := types.NewEIP155Signer(chainID).Hash(tx)
	fmt.Println(h.String())

	// 签名
	sign, err := hexutil.Decode("0x84cbed759da1e60939fa7f21ed40583d3034d53e81faf9594b1a42b8d470041e03620d2570bc51bc928c82c09348af0b5b0ceb931ed92578181a674078ace48d01")
	signedTx, err = tx.WithSignature(types.NewEIP155Signer(chainID), sign)
	if err != nil {
		log.Fatal(err)
	}
	ts := types.Transactions{signedTx}
	var w = new(bytes.Buffer)
	ts.EncodeIndex(0, w)
	fmt.Println(hex.EncodeToString(w.Bytes())) // f86...772

}
