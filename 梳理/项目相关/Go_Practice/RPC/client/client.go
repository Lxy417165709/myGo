package main

import (
	"fmt"
	"net/rpc"
)

//
//func main() {
//	var client *rpc.Client
//	var err error
//	if client,err = rpc.Dial("tcp",":1234");err!=nil{
//		panic(err)
//	}
//	var reply string
//	if err = client.Call("RPCService.Hello","李学悦",&reply);err!=nil{
//		panic(err)
//	}
//	fmt.Println(reply)	// Hello:李学悦
//}

func main() {
	var client *rpc.Client
	var err error
	if client, err = rpc.Dial("tcp", ":1234"); err != nil {
		panic(err)
	}
	var reply int
	if err = client.Call("RPCService.CountSum", []int{1, 2, 3, 4, 5}, &reply); err != nil {
		panic(err)
	}
	fmt.Println(reply) // 15
}

// 调用成功后，客户端和服务器都退出了
