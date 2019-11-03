package main

import "fmt"

func countCharacters(words []string, chars string) int {
	m:=make(map[uint8]int)
	for i:=0;i< len(chars);i++  {
		m[chars[i]]++
	}
	ans:=0
	for i:=0;i< len(words);i++  {
		temp :=make(map[uint8]int)
		word := words[i]
		flag:=0
		for t:=0;t< len(word);t++  {
			temp[word[t]]++
			if m[word[t]] - temp[word[t]]<0 {
				flag=1
				break
			}
		}
		if flag==0 {
			ans+= len(word)
		}
	}
	return ans
}

func main() {
	fmt.Println(countCharacters([]string{"hello","world","leetcode"},"welldonehoneyr"))
}

/*
	总结
	1. 这是一道简单的哈希映射题

*/