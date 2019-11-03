package main

import (
	"fmt"
	"sort"
)

type stock struct {
	index int
	money int
}
type Stocks []stock

func (s Stocks)Len() int{
	return len(s)
}
func (s Stocks)Less(i,j int) bool{
	return s[i].money<s[j].money
}

func (s Stocks)Swap(i,j int){
	s[i],s[j] = s[j],s[i]
}

func maxProfit(prices []int) int {
	stocks := make([]stock,0)
	for i:=0;i< len(prices);i++  {
		stocks = append(stocks, stock{i,prices[i]})
	}
	sort.Sort(Stocks(stocks))
	//fmt.Println(stocks)
	max:=0

	for l:=0;l< len(stocks);l++  {
		for r:=len(stocks)-1;r>l ;r--  {
			if stocks[r].index>stocks[l].index {
				x := stocks[r].money-stocks[l].money
				if x>max {
					max = x
					break
				}
			}
		}
	}
	return max

}

func main() {
	fmt.Println(maxProfit([]int{-221,9,-5,-11,19}))
}

/*
	总结
	1. 该题我使用了改进的暴力算法，先排序，再找寻，复杂度O(n^2)
	2. 官方有O(n)的算法，从左到右推进，找到更低的点就替换最低的，并且每次移动计算最高收益
	3. 评论还有用高等数学DP的，他说 区间和可以转换成求差的问题，求差问题，也可以转换成区间和的问题。
		求一个数组A从左到右的值(右-左)相差最大值，等于求数组A元素之间的差(右-左)形成的数组B(长度为len(A)-1)，的最大连续子序列和
		比如 B[1] + B[2] + B[3] = (A[2] - A[1]) + (A[3] - A[2]) + (A[4] - A[3]) = A[4] - A[1]

*/
