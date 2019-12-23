package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	return step_2(path)
}

func step_1(path string) string{
	ans:=""
	flag:=0
	for i:=0;i< len(path);i++  {

		if flag==0 {
			if path[i]=='/' {
				flag=1
			}
			ans+=string(path[i])
		}else{
			if path[i]!='/' {
				flag=0
				ans+=string(path[i])
			}
		}
	}
	return ans
}

func step_2(path string) string{
	slice := strings.Split(path,"/")
	stack := make([]string,0)

	for i:=0;i< len(slice);i++  {
		if slice[i]=="" {
			continue
		}
		if slice[i]==".." {
			if len(stack)!=0 {
				stack = stack[:len(stack)-1]
			}
		}else{
			if slice[i]!="." {
				stack = append(stack, slice[i])
			}
		}
	}
	ans:=""
	for i:=0;i<len(stack) ;i++  {
		ans+= "/" + stack[i]
	}
	//fmt.Println(slice,stack, len(slice), len(stack))
	if ans=="" {
		ans="/"
	}
	return ans
}

func main() {
	s:=""
	for ;true ;  {
		fmt.Scan(&s)
		fmt.Println(simplifyPath(s))
	}

}

/*
	总结
	1. 该题考查字符串处理 + 栈
	2. 可以不经过step_1处理！

*/
