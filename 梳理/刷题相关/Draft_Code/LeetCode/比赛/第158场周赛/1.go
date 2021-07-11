package main

import "fmt"

func balancedStringSplit(s string) int {
	l,r:=0,0
	ans:=0
	for i:=0;i< len(s);i++  {
		if s[i]=='L' {
			l++
		}else{
				r++
		}
		if l==r {
			ans++
		}
	}
	return ans
}

func main() {
	fmt.Println( balancedStringSplit("L"))
}
