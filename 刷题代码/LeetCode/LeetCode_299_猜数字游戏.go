package main

import (
	"fmt"
	"strconv"
)

func min(a,b int) int{
	if a>b {
		return b
	}else{
		return a
	}
}

func getHint(secret string, guess string) string {
	me,friend := make(map[uint8]int),make(map[uint8]int)
	me_slice := make([]uint8,0)
	numOfA :=0
	numOfB :=0
	for i:=0;i< len(secret);i++  {
		if me[secret[i]] == 0 {
			me_slice = append(me_slice, secret[i])
		}
		me[secret[i]]++
	}
	for i:=0;i< len(guess);i++  {
		friend[guess[i]]++
	}
	for i:=0;i<len(secret) ;i++  {
		if secret[i]==guess[i] {
			me[secret[i]]--
			friend[secret[i]]--
			numOfA++
		}
	}

	for i:=0;i< len(me_slice);i++  {
		numOfB += min(me[me_slice[i]],friend[me_slice[i]])
	}

	return strconv.Itoa(numOfA) + "A" + strconv.Itoa(numOfB) + "B"
}


func main() {
	fmt.Println(getHint("123459410","123451190"))
}
/*
	总结
	1. 这是一道中等的哈希题~
	2. A的个数通过for循环判断位置相同时字母是否相等可以获得，B的数字等于去除了A的个数后，2个字符串还有多少个字母相同(位置不同)
	3. 这题我的写法有些复杂了 0.0..
*/