package main

import (
	"fmt"
	"net"
)

func main() {
	// 1.创建服务器地址
	addr,_:=net.ResolveTCPAddr("tcp4","localhost:8899")

	// 2.创建连接
	conn,_:=net.DialTCP("tcp",nil,addr)

	// 3.发送数据 (无限发送)
	for  {
		var a string
		fmt.Scan(&a)
		conn.Write([]byte(a))
	}


	// 4.关闭连接
	conn.Close()
}
