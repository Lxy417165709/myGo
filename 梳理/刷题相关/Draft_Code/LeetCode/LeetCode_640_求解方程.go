package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solveEquation(equation string) string {
	all:=strings.Split(equation,"=")
	left:=all[0]
	right:=all[1]
	lft:=get(left)
	rit:=get(right)
	a,b:=lft[1],lft[0]
	c,d:=rit[1],rit[0]

	up:=d-b
	down:=a-c

	if up==0 && down==0 {
		return "Infinite solutions"
	}
	if up!=0 && down==0 {
		return "No solution"
	}
	return "x="+strconv.Itoa(up/down)
}
func get(a string) []int{
	ans:=make([]int,2)
	number:=0
	flag:=0
	for i:=0;i< len(a);i++  {
		if a[i]>='0' && a[i]<='9' {
			for ;i< len(a) && a[i]>='0' && a[i]<='9' ;i++  {
				number = number*10+int(a[i]-'0')
			}
			if flag==1 {
				number=-number
			}
			if i==len(a) {
				ans[0]+=number
			}else{
				if a[i]=='x' {
					ans[1]+=number
				}else{
					ans[0]+=number
					i--
				}
			}
			flag=0
			number=0
			continue
		}
		if a[i]=='-' {
			flag=1
			continue
		}
		// x前面有数字的话，在前面已经处理了，所以能到这的x前一定没有数字
		if a[i]=='x' {
			if flag==1 {
				ans[1]+=-1
			}else{
				ans[1]+=1
			}
			flag=0
			number = 0
			continue
		}
	}
	return ans

}


func main() {
	fmt.Println(solveEquation("5x=-12+5x+x"))
}
/*
	总结
	1. 这题虽然AC了，但是我觉得代码还不够优美！
	2. 我的思路是先获取他们的系数，然后通过方程公式求解 0.0
*/