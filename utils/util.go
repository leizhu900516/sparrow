package utils

import (
	"crypto/aes"
	"crypto/cipher"
	cryptoRand "crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"reflect"
	"strings"
	"time"
)

func CheckErr(err error){
	if err !=nil{
		fmt.Println(err)
	}
}


func Isfile(filename string) bool {
	_,err := os.Stat(filename)
	if err != nil{
		return false
	}
	return true

}

func getIp()(string,bool){
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		return "",false
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						fmt.Println(ipnet.IP.String())
						return ipnet.IP.String(),true
					}
				}
			}
		}
	}
	return "",false
}

//判断文件是否存在
func CheckFileState(filepath string)  bool{
	_,err := os.Stat(filepath)
	if err != nil{
		return false
	}
	return true
}

// 生成指定长度的随机字符串
func  GetRandomString(l int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}


func FormatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

/*
AES CFB加解密
secretkey必须是16 24 32 位的整数
*/
var secretKey = "6368616e676520896869732070617373"
//解密
func NewCFBDecrypter(enCode string) (bool,string){
	key, _ := hex.DecodeString(secretKey)
	ciphertext, _ := hex.DecodeString(enCode)
	block, err := aes.NewCipher(key)
	if err != nil {
		return false,""
	}
	if len(ciphertext) < aes.BlockSize {

		return false,""
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	//fmt.Printf("%s\n", ciphertext)
	return true,string(ciphertext)
}
//加密
func NewCFBEncrypter(str string) string {
	key, _ := hex.DecodeString(secretKey)
	plaintext := []byte(str)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(cryptoRand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return hex.EncodeToString(ciphertext)
}


func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	//if t.Kind() == reflect.Ptr {
	//	t = t.Elem()
	//}
	//if t.Kind() != reflect.Struct {
	//	log.Println("Check type error not Struct")
	//	return nil
	//}
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[strings.ToLower(t.Field(i).Name)] = v.Field(i).Interface()
	}
	return data
}
