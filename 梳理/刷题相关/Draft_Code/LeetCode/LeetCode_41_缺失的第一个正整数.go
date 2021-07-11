package main
func firstMissingPositive(nums []int) int {
	low,high:=1,len(nums)
	for i:=0;i<len(nums);i++{
		for nums[i]>=low && nums[i]<=high && nums[i]!=i+1{
			if nums[nums[i]-1] == nums[i]{
				break
			}
			nums[i],nums[nums[i]-1] = nums[nums[i]-1],nums[i]
		}
	}
	for i:=0;i<len(nums);i++{
		if nums[i]!=i+1{
			return i+1
		}
	}
	return high+1
}
/*
	总结
	1. 我的方法是先获取答案可能的范围，之后遍历整个数组，如果在范围内，则进行位置交换(注意如果交换前后值一样则不交换)
	2. 遍历完成后，再扫描一次数组，找到那些不符合位置要求的数就可以了。如果全满足，则说明答案是high+1
*/
