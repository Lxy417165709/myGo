package main

func findLHS(nums []int) int {
	count := make(map[int]int)
	ans := 0
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
		if count[nums[i]-1] != 0 {
			ans = max(count[nums[i]]+count[nums[i]-1], ans)
		}
		if count[nums[i]+1] != 0 {
			ans = max(count[nums[i]]+count[nums[i]+1], ans)
		}
	}

	return ans
}

func max(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	a, b := arr[0], max(arr[1:]...)
	if a > b {
		return a
	}
	return b
}

/*
	总结
	1. 刚刚开始想复杂了..以为要先排序，然后去重，之后获得子序列..
	2. 这题直接哈希，记录数字的出现次数，再按照题目要求进行相应操作就可以了。
	3. 这题还可以用排序，把空间复杂度优化为O(1)，但是时间复杂度就变为了O(nlogn)
*/
