package main

import "fmt"

var m map[uint8]string
func letterCombinations(digits string) []string {
	if digits=="" {
		return []string{}
	}
	m = make(map[uint8]string)
	m['2'] = "abc"
	m['3'] = "def"
	m['4'] = "ghi"
	m['5'] = "jkl"
	m['6'] = "mno"
	m['7'] = "pqrs"
	m['8'] = "tuv"
	m['9'] = "wxyz"
	return solve(digits)
}

func solve(digits string) []string{

	slice:=make([]string,0)
	ss := m[digits[0]]
	if len(digits)==1 {
		for i:=0;i< len(ss);i++  {
			slice = append(slice, string(ss[i]))
		}
		return slice
	}

	for i:=0;i< len(ss);i++  {
		s_slice := solve(digits[1:])
		for t:=0;t<len(s_slice) ;t++  {
			s := string(ss[i]) + s_slice[t]
			slice = append(slice, s)
		}
	}
	return slice

}

func main() {
	fmt.Println(letterCombinations("3"))
}


/*
	总结
	1. 该题目考察回溯递归，注意递归函数的意义和返回值！！！
*/