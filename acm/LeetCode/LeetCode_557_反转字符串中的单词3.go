package main

import (
	"fmt"
	"strings"
)
func rev(s string) string{
	x := []byte(s)
	for i:=0;i<len(x)/2 ;i++  {
		x[len(x)-1-i],x[i] = x[i],x[len(x)-i-1]
	}
	return string(x)
}
func reverseWords(s string) string {
	slice:=strings.Split(s," ")
	ans:=""
	for i:=0;i< len(slice);i++  {
		if i==0 {
			ans+=rev(slice[i])
		}else{
			ans+=" " + rev(slice[i])
		}
	}
	return ans
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest a ab"))
}
/*
	总结
	1. 我的做法是暴力法，先切割出单词，分别对单词反转，再合并 0.0..
*/