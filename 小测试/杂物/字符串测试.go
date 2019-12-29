package 杂物

import "fmt"

func main() {

	str:="abc"
	fmt.Println(str[3:])
	fmt.Println(str[-1:])
}

/*
	总结
	1. 对于合法的字符串[a:b], a必须>=0, b必须≤字符串的长度
*/
