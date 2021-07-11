package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	m := make(map[int]string)
	m[1] = "1"
	for i:=2;i<=n ;i++  {
		base:=m[i-1][0]
		num:=1
		for t:=1;t< len(m[i-1]);t++  {
			if m[i-1][t]==base {
				num++
			}else{
				m[i] = m[i] + strconv.Itoa(num)+string(base)
				base=m[i-1][t]
				num=1
			}
		}
		m[i] = m[i] + strconv.Itoa(num)+string(base)
	}

	return m[n]

}

func main() {
	fmt.Println(countAndSay(30))
}

/*
	总结
	1. 这题考查理解能力，反正我是看评论才理解的~
*/