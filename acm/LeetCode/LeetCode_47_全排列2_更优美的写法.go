package main

func permuteUnique(nums []int) [][]int {
	return solve(nums)
}
func solve(nums []int) [][]int{


	ans:=make([][]int,0)
	if len(nums)==1{
		ans = append(ans,[]int{nums[0]})
		return ans
	}

	mark:=make(map[int]int)

	for i:=0;i<len(nums);i++{
		// 不允许重复
		if mark[nums[i]]==1{
			continue
		}
		mark[nums[i]]=1
		tmp:=make([]int,0)
		tmp = append(tmp,nums[:i]...)
		tmp = append(tmp,nums[i+1:]...)

		// 解决n-1的全排列
		lessAns := solve(tmp)
		for t:=0;t<len(lessAns);t++{
			lessAns[t] = append(lessAns[t],nums[i])
			elm:=make([]int,len(lessAns[t]))
			copy(elm,lessAns[t])
			ans = append(ans,elm)
		}

	}

	return ans
}

func main() {

}
/*
	总结
	1. 这和替换的很像，但是这次我不用替换了，我直接截取子串全排列 0.0
*/
