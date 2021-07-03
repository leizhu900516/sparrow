package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)
const   (
	mySigningKey  = "qazwsxedc0998615453"
)
//各种平台的代码
type platformType uint8
const (
	WEIXIN platformType  = iota
	BLOG
	CARPORT

)
//type MyCustomClaims struct {
//	Appid string `json:"appid"`
//	jwt.StandardClaims
//}

/**
jwt 生成 目前jwt 有效时间1小时，可自定义调整
iss: 签发者

sub: 面向的用户

aud: 接收方

exp: 过期时间

nbf: 生效时间

iat: 签发时间

jti: 唯一身份标识
*/
func NewJwt(platform uint8,userid uint) (string,error){
	claims := jwt.MapClaims{
		"id":userid, 		//userid
		"platform":platform, //平台id
		"appid":platform, 	//应用id
		"exp":time.Now().Add(time.Hour * time.Duration(24)).Unix(),
		"nbf":time.Now().Unix(),
		"iat":time.Now().Unix(),

	}
	//claims := MyCustomClaims{
	//	"",
	//	jwt.StandardClaims{
	//		ExpiresAt: time.Now().Add(time.Minute * time.Duration(1)).Unix(),
	//		Issuer:    apptype,
	//	},
	//}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//fmt.Println(token)
	tokenString, err := token.SignedString([]byte(mySigningKey))
	return tokenString,err
}
func ParseJwt(tokenstring string) (interface{},bool){

	//fmt.Println("token=",tokenstring)
	if tokenstring == "" ||tokenstring == "null"{
		log.Printf("请求未携带token无权限访问")
		return "请求未携带token无权限访问",false
	}
	token, err := jwt.Parse(tokenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})
	if err != nil{
		return "token无效",false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
	if token.Valid {
		log.Println("You look nice today")
		if claims, ok := token.Claims.(jwt.MapClaims);ok{
			fmt.Println("claims=",claims)
			return claims,true
		}
		return nil,true
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
			return "token无效",false
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
			return "token已经过期",false
		} else {
			fmt.Println("Couldn't handle this token:", err)
			return err,false
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
		return err,false
	}

}




