package main

import "fmt"

func countVowelPermutation(n int) int {
	if n==0 {
		return 0
	}

	var m = [150][20050]int{}
	m['a'][1] = 1
	m['e'][1] = 1
	m['i'][1] = 1
	m['o'][1] = 1
	m['u'][1] = 1
	sum :=5
	for i:=2;i<=n ;i++  {
		m['a'][i] = m['e'][i-1]%1000000007
		m['e'][i] = (m['i'][i-1]+m['a'][i-1])%1000000007
		m['i'][i] = (m['a'][i-1]+m['e'][i-1]+m['o'][i-1]+m['u'][i-1])%1000000007
		m['o'][i] = (m['u'][i-1]+m['i'][i-1])%1000000007
		m['u'][i] = m['a'][i-1]%1000000007
		sum=(m['a'][i]+m['e'][i]+m['i'][i]+m['o'][i]+m['u'][i])%1000000007
		fmt.Println(sum)
	}
	return sum %1000000007


}


func main() {
	fmt.Println(countVowelPermutation(20000))
}
/*
	总结
	1. 这是一道dp题，dp规则题目都告诉我了 0.0.
*/