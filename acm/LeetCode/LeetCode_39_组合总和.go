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

	if target==0 {
		sli1 := make([]int, len(sli))
		copy(sli1,sli)	// 一定要用这个拷贝，这样sli和sli1指向的地址就是不一样的
		ans = append(ans,sli1)
		return
	}
	for i:=0;i<len(candidates) && target>=candidates[i];i++  {
		sli = append(sli, candidates[i])
		solve(candidates[i:],target-candidates[i],sli)
		sli=sli[:len(sli)-1]
	}

}


func main() {
	fmt.Println(combinationSum([]int{2,3,5},10))
}


/*
	总结
	1. 要注意切片的深拷贝和浅拷贝
		copy(a,b) 则a和b指向的地址不一样,但是内容一样
		a=b 则a和b指向的地址一样,内容也一样
	2. 回溯法要知道边界条件和剪枝
*/