package main

import (
	"fmt"
	"net"
)

func main() {
	// 1.创建服务器地址
	addr,_ := net.ResolveTCPAddr("tcp4","localhost:8899")

	// 2. 创建监听器
	lis,_ := net.ListenTCP("tcp4",addr)


	// 3. 通过监听器获取客户端传递过来的数据
	// 阻塞式
	conn,_:=lis.Accept()

	// 无限接收数据
	for {
		// 4.转换数据
		b:=make([]byte,1024)
		n,_ :=conn.Read(b)
		fmt.Println("获取到的数据是:",string(b[:n]))
	}




	// 5.关闭连接
	conn.Close()



}
