package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices)==0 {
		return 0
	}

	now:=prices[0]
	ans:=0

	for i:=1;i< len(prices);i++  {
		if prices[i]<=now {
			now = prices[i]
		}else{
			d:=0
			for ;i<len(prices) && prices[i]>=prices[i-1] ;i++  {
				d=prices[i] - now
			}
			ans+=d
			if i<len(prices) {
				now = prices[i]
			}
		}
	}
	return  ans

}

func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}

/*
	总结
	1. 这是一道简答题，要注意数组空的时候！

*/