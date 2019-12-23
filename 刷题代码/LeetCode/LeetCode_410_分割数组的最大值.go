package main


func splitArray(nums []int, m int) int {
	sum = make([]int,1)
	for i:=1;i<=len(nums);i++{
		sum = append(sum,sum[i-1] + nums[i-1])
	}
	mp = make(map[int]int)
	return solve(nums,0,m)
}

var mp map[int]int
var sum []int
const inf = 100000000000
func solve(nums []int, begin int,m int) int{


	number:=begin<<20|m
	if x,ok:=mp[number];ok{
		return x
	}

	if m==1{

		mp[number] = sum[len(nums)]-sum[begin]
		return mp[number]
	}
	sum:=0
	ans:=inf
	for i:=begin;i<=len(nums)-m;i++{
		sum+=nums[i]
		ans = min(ans,max(sum,solve(nums,i+1,m-1)) )
	}
	mp[number] = ans
	return mp[number]
}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}


func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}
func main() {


}
/*
	总结
	1. 这题我使用记忆化搜索AC的
	2. 注意inf的定义，一定要够大呀！！！
	3. 官方还使用了贪心+二分快速AC了。。。
	4. 这题目好像是分巧克力的那个题目 0.0
*/