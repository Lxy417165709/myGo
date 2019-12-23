package main

import (
	"fmt"
	"strings"
)

var m map[string]bool
func isValidSerialization(preorder string) bool {
	m=make(map[string]bool)
	slice:=strings.Split(preorder,",")
	a:=solve1(slice)
	//fmt.Println(solve(slice),solve1(slice))
	return a
}

// TLE版本
func solve(slice []string) bool{
	str:=strings.Join(slice,",")

	// 这注释了就没优化，没注释就优化了 0.0
	if x,ok:=m[str];ok {
		return x
	}

	if len(slice)==0 {
		m[str] = false
		return m[str]
	}
	if len(slice)==1{
		m[str] = slice[0]=="#"
		return m[str]
	}else{
		if slice[0]=="#" {
			m[str] = false
			return m[str]
		}
	}


	for i:=2;i< len(slice);i=i+2  {
		left:=solve(slice[1:i])
		right:=solve(slice[i:])
		//fmt.Println(slice[1:i],slice[i:],left,right)
		if left&&right{
			m[str]=true
			return m[str]
		}
	}
	m[str]=false
	return m[str]
}


func solve1(slice []string) bool{

	stack :=make([]string,0)
	for i:=0;i< len(slice);i++  {
		if slice[i]=="#" {
			for len(stack)>=2 && stack[len(stack)-1]=="#"  {
				stack = stack[:len(stack)-2]
			}
		}
		if len(stack)!=0 && stack[0]=="#" {
			return false
		}
		stack  = append(stack, slice[i])
	}
	fmt.Println(stack)
	return len(stack)==1 && stack[0]=="#"
}

func main() {
	fmt.Println(isValidSerialization( "1,#,#,#,#"))
}
/*
	总结
	1. 第一次没用哈希优化和第二次用哈希优化使用的递归都TLE了 0.0..（通过验证，我的确是优化了的 0.0..）
	2. 第二次看了官方后，知道了一些方法，然后写出来AC了，但是代码不够优美，之后改更优美一些 0.0.
	3. 官方有不用栈和使用前序遍历性质的(非叶子节点总比叶子节点少1，把#看作叶子节点然后进行遍历的)，可以借鉴 0.0，大牛啊啊啊~！！！！

*/