package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	if pattern==""&&str=="" {
		return true
	}


	s1:= process(pattern)
	s2:= process1(str)
	fmt.Println(s1,s2)
	if len(s1)!=len(s2) {
		return false
	}
	for i:=0;i< len(s1);i++  {
		if len(s1[i])!=len(s2[i]) {
			return false
		}
		for t:=0;t< len(s1[i]);t++  {
			if s1[i][t]!=s2[i][t] {
				return false
			}
		}
	}
	

	return true

}

func process(s string) [][]int{
	slice := make([][]int,0)
	m :=make(map[uint8]int)
	index:=1
	for i:=0;i< len(s);i++  {
		if m[s[i]]==0 {
			m[s[i]] = index
			index++
			slice = append(slice, []int{})
		}
		slice[m[s[i]]-1] = append(slice[m[s[i]]-1], i)
	}
	return slice
}

func process1(s string) [][]int{
	x:=strings.Split(s," ")

	slice := make([][]int,0)
	m :=make(map[string]int)
	index:=1
	for i:=0;i< len(x);i++  {
		if m[x[i]]==0 {
			m[x[i]] = index
			index++
			slice = append(slice, []int{})
		}
		slice[m[x[i]]-1] = append(slice[m[x[i]]-1], i)
	}
	return slice
}
func main() {
	fmt.Println(wordPattern("aaaacaca","wo wo wo wo shi wo shi wo"))
}

/*
	总结
	1. 这是一道简单的字符串哈希题
*/