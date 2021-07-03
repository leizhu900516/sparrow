package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"sparrow/utils"
	"strconv"
	"strings"
	"time"
)


//var secretKey = "abcdefghjklmnopqrstuvwxyz1234567"
var secretKey = "6368616e676520746869732070617373"
func NewCFBDecrypter(destr string) string{
	key, _ := hex.DecodeString(secretKey)
	ciphertext, _ := hex.DecodeString(destr)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	fmt.Printf("%s\n", ciphertext)
	return string(ciphertext)
}
func NewCFBEncrypter(str string) string {
	fmt.Println(len(secretKey))
	key, _ := hex.DecodeString(secretKey)
	plaintext := []byte(str)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return hex.EncodeToString(ciphertext)
}
func main() {
	localTimestemp := time.Now().AddDate(0,0,3).Unix()
	timeTempstr := strconv.FormatInt(localTimestemp,10)
	enStr := "abcdefg" +timeTempstr  + utils.GetRandomString(10)
	fmt.Println(">>>",enStr)
	//enStr := "1234567890"
	aaa:=NewCFBEncrypter(enStr)
	ddd:=NewCFBDecrypter("7c4c68ecf59061facd3405ccb25d13dca0886fbbba58ac8a8207986a13065bf720e593341a")
	fmt.Println(aaa)
	fmt.Println(ddd)
	fmt.Println(strings.Split(ddd,"="))
	fmt.Println(time.Now().AddDate(0,0,3).Unix())
}
