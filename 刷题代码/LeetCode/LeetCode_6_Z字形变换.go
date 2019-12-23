package main

import "fmt"

var ans [1005][1005]uint8

func convert(s string, numRows int) string {

	if numRows==1 {
		return s
	}else{
		if numRows==0 {
			return ""
		}
	}

	ans = [1005][1005]uint8{}
	num:=0
	time:=0
	for ;num<len(s) ;  {
		for i:=0;i<numRows && num<len(s);i,num = i+1,num+1  {
			ans[i][time*(numRows-1)] = s[num]
		}

		for i:=numRows-2;i>=1 && num<len(s);i,num = i-1,num+1  {
			ans[i][(time+1)*(numRows-1)-i] = s[num]
		}
		time++
	}
	mm:=""
	for i:=0;i< 1005;i++  {
		for t:=0;t< 1005;t++  {
			if ans[i][t]!=0 {
				mm += string(ans[i][t])
			}
		}
	}
	return mm
}
func main() {
	fmt.Println(convert("abcdefghijklmn",1))
}

/*
	总结
	1. 自己用的是暴力构造 O(n)，之后暴力读取的方法空间复杂度O(1005*1005)
	2. 官方有更好的方法！时间复杂度O(n)空间复杂度O(n)
	3. 官方使用vector，可以压缩空间，相当于把ans数组向右压实，官方模拟更优美

*/