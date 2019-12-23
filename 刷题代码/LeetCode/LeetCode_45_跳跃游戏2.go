package main
func jump(nums []int) int {
	return solve(nums,0)
}

func solve(nums []int,nowPlace int) int{

	if nowPlace == len(nums)-1{
		return 0
	}
	if nowPlace + nums[nowPlace] >= len(nums)-1{
		return 1
	}

	mx:=0
	index:=0
	for i:=1;i<=nums[nowPlace];i++{
		if mx< nums[nowPlace + i]+i{
			mx = nums[nowPlace + i]+i
			index = i
		}
	}

	return solve(nums,nowPlace + index)+1

}

func main() {

}
/*
	总结
	1. 第一次用DP，超时了
	2. 第二次用贪心，AC了，贪心策略是，每次都找能够到达最远的下一个位置。 但是我写的算法的复杂度不算很好..
*/