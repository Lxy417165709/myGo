package main

import "fmt"

func opt( a int, b int) (int,int,int,float64) {
	// go函数支持多返回值
	return  a+b, a-b, a*b, float64(a)/float64(b)
}

func main(){
	a, b := 20, 10
	fmt.Println(opt(a, b))
}
