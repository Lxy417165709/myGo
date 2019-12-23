package main

import (
	"fmt"
	"sort"
)

var ans [][]int

func combinationSum(candidates []int, target int) [][]int {


	ans = make([][]int,0)
	sli:= make([]int,0)
	sort.Ints(candidates)
	solve(candidates,target,sli)

	return ans
}
//var length = 0
func solve(candidates []int, target int,sli []int){
	// 标记某些数字在该层已使用过
	var m map[int]int
	m = make(map[int]int)
	if target==0 {
		sli1 := make([]int, len(sli))
		copy(sli1,sli)	// 一定要用这个拷贝，这样sli和sli1指向的地址就是不一样的
		ans = append(ans,sli1)
		return
	}
	for i:=0;i<len(candidates) && target>=candidates[i];i++  {

		// 就这不同
		if m[candidates[i]]==0 {
			m[candidates[i]]=1
			sli = append(sli, candidates[i])
			solve(candidates[i+1:],target-candidates[i],sli)	//上题是[i:]
			sli=sli[:len(sli)-1]
		}
	}
}


func main() {
	fmt.Println(combinationSum([]int{10,1,1,7,6,1,5},18))
}


/*
	总结
	1. 该题目是不能重复使用同一数字，与上题相比不同之处可以看上面
	2. 回溯法要知道边界条件和剪枝
*/