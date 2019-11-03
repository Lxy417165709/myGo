package main

func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}
func minMoves(nums []int) int {
	mn:=nums[0]
	for i:=0;i<len(nums);i++{
		mn = min(mn,nums[i])
	}
	ans:=0
	for i:=0;i<len(nums);i++{
		ans+=nums[i]-mn
	}
	return ans
}
func main() {

}
/*
	总结
	1. n-1个数变大，我们可以认为是1个数变小 0.0.
	2. 官方还有很多解法。。
*/