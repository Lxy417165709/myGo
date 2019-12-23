package main

import "fmt"

func lengthOfLastWord(s string) int {

	flag:=0
	ans:=0
	for i:= len(s)-1;i>=0 ;i--  {
		if flag==0 {
			if s[i]!=' ' {
				ans++
				flag=1
			}
		}else{
			if s[i] == ' ' {
				break
			}else{
				ans++
			}
		}

	}
	return ans
}

func main() {
	fmt.Println(lengthOfLastWord(""))
}


/*
	总结
	1. 这是一道简单题！
*/