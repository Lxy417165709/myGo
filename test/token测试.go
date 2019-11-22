package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func test(){
	// 定义一个key
	key:="Lxy"

	// 这是信息
	info := "我是要加密的信息"

	//这里是个例子，并包含了一个故意签发一个已过期的token
	// ExpiresAt:time.Now().Unix()-1000 过期时间
	data := jwt.StandardClaims{Subject:info,ExpiresAt:time.Now().Unix()-1000}
	tokenInfo := jwt.NewWithClaims(jwt.SigningMethodHS256,data)

	// 获取加密结果
	tokenStr,err := tokenInfo.SignedString([]byte(key))
	if err != nil {
		fmt.Println("出错了!")
		return
	}
	fmt.Println("myToken is: ",tokenStr)


	// 解密
	result,ok := tokenInfo.Claims.(jwt.StandardClaims)
	if !ok {
		fmt.Println("解析错误")
		return
	}

	//校验下token是否过期
	succ := result.VerifyExpiresAt(time.Now().Unix(),true)

	fmt.Println("succ",succ)
	//获取token中保存的用户信息
	fmt.Println(result.Subject)
}
func main() {
	test()
}
