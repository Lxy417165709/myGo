package main

import (
	"fmt"
	"strings"
)


func judge(x uint8) bool{

	return x>='0' &&x<='9' || x>='a' && x<='z'

}


func isPalindrome(s string) bool {
	l:=0
	r:= len(s)-1
	s = strings.ToLower(s)
	fmt.Println(s)
	for ;l<=r ;  {
		if judge(s[l]) && judge(s[r])  {
			if s[l]!=s[r] {
				return false
			}else{
				l++
				r--
			}
		}else{
			for ;l<len(s) && judge(s[l])==false ;  {
				l++
			}
			for ;r>= 0 && judge(s[r])==false ;  {
				r--
			}
		}
	}
	return true
}


func main() {
	fmt.Println(isPalindrome("race e   ,, car"))
}

/*
	总结
	1. 该题是字符串处理 + 判断回文串，我使用双指针法解
*/