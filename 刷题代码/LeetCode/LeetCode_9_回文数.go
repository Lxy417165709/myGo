package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x<0 {
		return false
	}
	s:=strconv.Itoa(x)
	for i:= 0;i< len(s)/2;i++  {
		if s[i]!=s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-0))
	fmt.Println(isPalindrome(-1))
	fmt.Println(isPalindrome(1234321))
	fmt.Println(isPalindrome(56765))
}

/*
	总结
	1. 判断字符串是否回文,只需要循环其索引[0,len/2)

*/