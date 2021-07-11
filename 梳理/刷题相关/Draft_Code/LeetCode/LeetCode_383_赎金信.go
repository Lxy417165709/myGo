package main

import "fmt"

var m [30]int
func canConstruct(ransomNote string, magazine string) bool {
	m := [30]int{}	//注意初始化
	for i:=0;i< len(ransomNote);i++  {
		m[ransomNote[i]-'a']++
	}

	for i:=0;i< len(magazine);i++  {
		m[magazine[i]-'a']--
	}
	for i:='a'-'a';i<= 'z'-'a';i++  {
		if m[i]>0 {
			return false
		}
	}
	return true
}


func main() {
	fmt.Println(canConstruct("aa","aab"))
}

/*
	总结
	1. 注意初始化!
	2. 这是一道简单的哈希题，空间换时间~
*/