package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
)

const (
	AES_KEY = "7uuO5PtrhlSOZpAm"
)

// 密码加密
func EncryptPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hash)
}

// 密码解密校验
func DecryptPassword(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	if err != nil {
		return false
	}

	return true
}

// pkcs7Padding填充
func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding填充反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	}
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

// base64加密
func base64Encrypt(data []byte) (string, error) {
	return base64.StdEncoding.EncodeToString(data), nil
}

// base64解密
func base64Decrypt(data string) ([]byte, error) {
	dataByte, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}

	return dataByte, nil
}

// AES加密
func AesEncrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	encryptBytes := pkcs7Padding(data, blockSize)
	encrypted := make([]byte, len(encryptBytes))
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	blockMode.CryptBlocks(encrypted, encryptBytes)
	return encrypted, nil
}

// AES解密
func AesDecrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	decrypted := make([]byte, len(data))
	blockMode.CryptBlocks(decrypted, data)
	decrypted, err = pkcs7UnPadding(decrypted)
	if err != nil {
		return nil, err
	}

	return decrypted, nil
}

// RSA加密
func RsaEncrypt(plainText []byte, path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 读取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	// pem解码
	block, _ := pem.Decode(buf)
	// x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	// 对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}

	return cipherText
}

// RSA解密
func RsaDecrypt(cipherText []byte, path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)

	block, _ := pem.Decode(buf)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	// 对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	return plainText
}

func main() {
	// 加、解密
	/*password := "123456"
	hash := EncryptPassword([]byte(password))
	fmt.Println("加密结果：", hash)
	fmt.Println("解密结果：", DecryptPassword([]byte(hash), []byte(password)))*/

	// AES加、解密
	/*data := []byte("CoinSky APP")
	key := []byte(AES_KEY)
	encryptResult, err := AesEncrypt(data, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(encryptResult)

	encryptResultBase64, _ := base64Encrypt(encryptResult)
	fmt.Println("AES加密结果：", encryptResultBase64)

	decryptResultBase64, err := base64Decrypt(encryptResultBase64)
	if err != nil {
		fmt.Println(err)
	}

	decryptResult, err := AesDecrypt(decryptResultBase64, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("AES解密结果：", string(decryptResult))*/

	// RSA加、解密
	plainText := []byte("token=ada&")
	encryptResult := RsaEncrypt(plainText, "cert/public_key.pem")
	fmt.Println("RSA加密结果：", string(encryptResult))

	decryptResult := RsaDecrypt(encryptResult, "cert/private_key.pem")
	fmt.Println("RSA解密结果：", string(decryptResult))

}
