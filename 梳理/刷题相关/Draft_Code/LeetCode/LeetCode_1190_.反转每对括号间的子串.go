package main

import "fmt"

func reverseParentheses(s string) string {
	return solve(s)
}
func rev(s string) string {
	by := []byte(s)
	for i := 0; i < len(by)/2; i++ {
		by[i], by[len(by)-1-i] = by[len(by)-1-i], by[i]
	}
	return string(by)
}
func solve(s string) string {
	tmp:=""
	slice:=make([]string,0)
	for i:=0;i< len(s);i++  {
		if s[i]=='(' || s[i]==')' {
			if s[i]=='(' {
				slice = append(slice, tmp)
				slice = append(slice, "(")
				tmp = ""
			}else{
				x:=""
				slice = append(slice, tmp)
				tmp = ""
				for ;slice[len(slice)-1]!="(" ;  {
					x = slice[len(slice)-1] + x
					slice = slice[:len(slice)-1]
				}
				slice = slice[:len(slice)-1]
				x = rev(x)
				slice = append(slice, x)
			}
		}else{
			tmp+=string(s[i])
		}
	}
	slice = append(slice, tmp)
	ans:=""
	for i:=0;i< len(slice);i++  {
		ans+=slice[i]
	}
	return ans

}

func main() {
	fmt.Println(reverseParentheses(""))
}
/*
	总结
	1. 该题考查栈操作，之前一直做错是因为把符号和字符串分为2个栈了，其实一个栈才能解决了
*/