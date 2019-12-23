package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {


	sort.Ints(citations)
	h:=0
	for i:= len(citations)-1;i>=0 ;i--  {
		if len(citations)-i<=citations[i] {
			h = len(citations)-i
		}
	}
	return h
}

func main() {
	fmt.Println(hIndex([]int{
		100,
	}))
}
/*
	总结
	1. 这题考查理解能力...我醉了
*/