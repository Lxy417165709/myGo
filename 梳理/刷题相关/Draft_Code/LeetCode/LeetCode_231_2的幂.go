package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n==1{
		return true
	}
	if n%2==1 || n==0 {
		return false
	}
	return isPowerOfTwo(n/2)

}
func main() {
	fmt.Println(isPowerOfTwo(0))
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(218))
	fmt.Println(isPowerOfTwo(1026))
	fmt.Println(isPowerOfTwo(2048))

}

/*
	总结
	1. 递归很有用!

*/