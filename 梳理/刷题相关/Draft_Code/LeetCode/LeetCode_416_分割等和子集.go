package main
func canPartition(nums []int) bool {
	mark = make(map[int]int)
	mp = make(map[int]bool)

	sum:=0
	for i:=0;i<len(nums);i++{
		sum+=nums[i]
	}
	if sum&1==1{
		return false
	}else{
		return solve(nums,sum/2)
	}
}

var mark map[int]int    // 标记已取
var mp map[int]bool     // 标记可组成
func solve(nums []int,target int) bool{
	if target==0{
		return true
	}
	if x,ok:=mp[target];ok{
		return x
	}
	for i:=0;i<len(nums);i++{
		if mark[i]==1{
			continue
		}

		mark[i] = 1
		if solve(nums,target-nums[i]) {
			mp[target] = true
			return  mp[target]
		}
		mark[i] = 0
	}
	mp[target] = false
	return  mp[target]
}
func main() {

}

/*
	总结
	1. 这题我使用记忆化搜索AC的
	2. 之前也有遇到过这样的题目，题意是分成两个子集且差最小，当时
	   看了答案后，用迭代dp AC了
*/

