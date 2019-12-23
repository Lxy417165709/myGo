package main

import "fmt"


func longestPalindrome(s string) string {
	ansS:=""
	for i:=0;i< len(s);i++  {
		l:=i-1
		r:=i+1
		tmpS:=string(s[i])
		for ;l>=0 && r< len(s) && s[l]==s[r];  {
			tmpS=string(s[l])+tmpS + string(s[r])
			l--
			r++
		}
		if len(tmpS)> len(ansS) {
			ansS=tmpS
		}

	}

	for i:=0;i< len(s)-1;i++  {
		l:=i-1
		r:=i+2
		if s[i]==s[i+1] {
			tmpS:=string(s[i]) + string(s[i+1])
			for ;l>=0 && r< len(s) && s[l]==s[r];  {
				tmpS=string(s[l])+tmpS + string(s[r])
				l--
				r++

			}
			if len(tmpS)> len(ansS) {
				ansS=tmpS
			}
		}
	}
	return ansS
}



func main() {
	fmt.Println(longestPalindrome("cdc"))
}
/*
	总结
	1. 该题我采用了中心扩展算法， 回文字符串的中心可能是1个字符，也可能是2个，所以我先遍历1个的，再左右拓展。
	   之后我再遍历2个的，左右拓展~ 0.0. 时间复杂度O(n^2)
	2. 该题最好的是使用马拉车算法 0.0 课本还有nlogn时间复杂度的
*/