package encrypt

import (
	"DAGoPkg/perr"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AesEncryptRaw(origData []byte, key []byte) ([]byte, error) {
	// 分组秘钥
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// 创建数组
	crypted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(crypted, origData)

	return crypted, nil
}

func AesEncryptBytes(origData []byte, key string) (string, error) {
	k := []byte(key)

	crypted, err := AesEncryptRaw(origData, k)
	if err != nil {
		return "", err
	}

	result := base64.StdEncoding.EncodeToString(crypted)

	return result, nil
}

// AesEncrypt .
func AesEncrypt(orig string, key string) (string, error) {
	return AesEncryptBytes([]byte(orig), key)
}

func AesDecryptBytes(cryted string, key string) ([]byte, error) {
	// 转成字节数组
	crytedByte, err := base64.StdEncoding.DecodeString(cryted)
	if err != nil {
		return nil, err
	}

	k := []byte(key)

	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig, err = PKCS7UnPadding(orig)
	if err != nil {
		return nil, err
	}

	return orig, nil
}

// AesDecrypt .
func AesDecrypt(cryted string, key string) (string, error) {
	orig, err := AesDecryptBytes(cryted, key)
	var result string
	if orig != nil {
		result = string(orig)
	}

	return result, err
}

// PKCS7Padding 补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7Padding 去码
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, perr.NewErr("len(data)=0", perr.EncryptAESPaddingErrCode)
	}

	unpadding := int(origData[length-1])

	if length < unpadding {
		return nil, perr.NewErr("data length < unpadding", perr.EncryptAESPaddingErrCode)
	}

	return origData[:(length - unpadding)], nil
}
