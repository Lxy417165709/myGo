package main

import "fmt"

func reverseString(s []byte)  {
	for i:=0;i< len(s)/2;i++  {
		s[i],s[len(s)-1-i] = s[len(s)-1-i],s[i]
	}
	fmt.Println(s)
}
func main() {
	reverseString([]byte{
		11,2,
	})
}

/*
	总结
	1. 这是一道水题
*/