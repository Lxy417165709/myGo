package main

import (
	"fmt"
	"sort"
)

func coinChange(coins []int, amount int) int {

	slice := make([]int,0)
	sort.Ints(coins)
	slice = append(slice, 0)

	for i:=1;i<=amount ;i++  {
		ans:=100000000
		for t:=0;t< len(coins);t++  {
			if i-coins[t]>=0 {
				if slice[i-coins[t]]+1 < ans{
					ans = slice[i-coins[t]]+1
				}
			}
		}

		slice = append(slice, ans)

	}
	if slice[amount]>=100000000 {
		slice[amount] = -1
	}
	return slice[amount]
}


func main() {
	fmt.Println(coinChange([]int{1},0))
}

/*
	总结
	1. 第一次解答错是因为amount等于0的时候返回0就可以了，而不用-1(在提交前多此一举了)
	2. 该题是动态规划~
*/