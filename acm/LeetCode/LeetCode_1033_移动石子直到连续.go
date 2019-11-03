package main

import (
	"fmt"
	"sort"
)

func numMovesStones(a int, b int, c int) []int {
	slice :=[]int{a,b,c}
	sort.Ints(slice)
	mx := slice[2]-slice[0]-2
	mn:=0

	if slice[0]+2 == slice[1] || slice[1]+2 == slice[2] {
		mn++
	}else{
		if slice[0]+1 < slice[1] {
			mn++
		}
		if slice[1]+1 < slice[2] {
			mn++
		}
	}

	return []int{mn,mx}
}
func main() {
	fmt.Println( numMovesStones(4,1,1000))
}

/*
	总结
	1. 注意移动规则，(只能从最右或最左边移)
*/