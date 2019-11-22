package main

import "fmt"

func reverseBits(num uint32) uint32 {
	var ans uint32=0
	for i:=0;i<=31 ;i++  {
		x:=num&1
		num >>=1
		ans = ans<<1 + x
	}
	return ans
}
func main() {
	fmt.Println(reverseBits(2147483648*2-1))
}
/*
	总结
	1. 这是一道简单的位运算题
*/