package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	store "go-ethereum-learn/contracts"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("xxx")
	if err != nil {
		log.Fatal(err)
	}

	//代币的地址
	tokenAddress := common.HexToAddress("0xa4E9f795042D677eEe61E70aeE9e9f42a8EBF60B")
	instance, err := store.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	//我的地址3
	fromAddress := common.HexToAddress("0x3DcB8292Ff80fEfd750F36aE83A7156158180B12")

	//我的地址2
	toAddress := common.HexToAddress("0x9CF52a2Ec3F920F88fDcA4eeb152dce45cD6A431")

	//我的地址1
	fromAddress1 := common.HexToAddress("0xFcD1aef048EaA60cD07076f27FC6E4C4c642BC01")

	//返回链ID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//私钥3
	privateKey, err := crypto.HexToECDSA("私钥")

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	//获取随机数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(30000000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasPrice = gasPrice

	//查询接收人可以转走我几个猫币
	allowance, err := instance.Allowance(&bind.CallOpts{}, fromAddress1, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(allowance)

	//我的地址1 -》 接收人地址
	from, err := instance.TransferFrom(auth, fromAddress1, toAddress, big.NewInt(5000000000000000000))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(from)

	time.Sleep(10 * time.Second)
	//查询接收人可以转走我几个猫币
	allowance, err = instance.Allowance(&bind.CallOpts{}, fromAddress1, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(allowance)

}
