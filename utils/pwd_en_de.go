//对用户密码进行加密
package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

//加密
func Encryption(pwd string) string {

	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	encodePW := string(hash)  // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	fmt.Println(encodePW)
	return encodePW
}

//解密
func Decryption(encryptionPwd,pwd string) bool{
	// 正确密码验证
	err := bcrypt.CompareHashAndPassword([]byte(encryptionPwd), []byte(pwd))
	if err != nil {
		fmt.Println("passwd wrong")
		return false
	} else {
		return true
	}
}
