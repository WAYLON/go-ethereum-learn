package main

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.HexToECDSA("e2bd0a5de14054d9c10af50766813b0e4d0d7c0b1a2cf58bd71d0ca8d64514dc")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	data := []byte("balxabala")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301
	encoding := base64.StdEncoding.EncodeToString(signature)
	fmt.Println(encoding)
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hexutil.Encode(sigPublicKey))
	fmt.Println(hexutil.Encode(publicKeyBytes))
	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches) // true
	decode, _ := hexutil.Decode("0xb98740ea326025a71b7748fc6aa8fe4daf7e8e659be54bc1e1950a43d7a38943757d9ecd130234c34b6bd2fa8a3b1517fa7db6ee0bb668c7800b04080406b3721c")
	sigPublicKeyECDSA, err := crypto.SigToPub(common.HexToHash("0x30FB84B1CB46A5B65A37C528BFF8C1A1DCBF0BCBFA1E5585F0707AAAAD8A74E1").Bytes(), decode)
	if err != nil {
		log.Fatal(err)
	}

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches) // true

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true
}
