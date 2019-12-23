package main

import (
	"fmt"
	"sort"
)

type arr [][]int


func (s arr) Len() int {
	return len(s)
}
func (s arr) Less(i, j int) bool {
	return   s[i][1] - s[i][0] < s[j][1] - s[j][0]
}
func (s arr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func twoCitySchedCost(costs [][]int) int {
	n:= len(costs)
	sum:=0
	for i:=0;i<n ;i++  {
		sum+=costs[i][0]
	}
	sort.Sort(arr(costs))

	for i:=0;i<n/2 ;i++  {
		sum+=costs[i][1] - costs[i][0]
	}
	return sum


}

func main() {
	fmt.Println(twoCitySchedCost([][]int{{10,20},{30,200}}))
}

/*
	总结
	1.  审题之后感觉是贪心，然后发现是总和，想到了动态规划..动态规划想到了2地，又回到贪心
		在贪心的时候不知道排序指标是什么..想了之后就去看答案了...
	2.  注意整体与局部，我们先让所有人去A，然后再在这群人选出一些人去B，排序指标是A->B花费费用的减小量
	3 	评论启发我合理猜测贪心的指标~(我不知道怎么证明555)

*/