package main

import "fmt"

const (
	inf = 10000000000
)

func maxSubArray(nums []int) int {
	return solve(nums)
}

func solve(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	mid := (len(nums) - 1) / 2
	l, r := mid, mid+1

	// 从中心向左边拓展，求取包括中心的左边最大连续和
	leftSum := 0
	leftMax := -inf
	for l >= 0 {
		leftSum += nums[l]
		l--
		leftMax = max(leftMax, leftSum)
	}

	// 从中心向右边边拓展，求取包括中心的右边最大连续和
	rightSum := 0
	rightMax := -inf
	for r < len(nums) {
		rightSum += nums[r]
		r++
		rightMax = max(rightMax, rightSum)
	}

	// 最大子序和 = max (不包括中心的左边的最大连续和, 不包括中心的右边的最大连续和, 中间的最大连续和)
	return max(
		solve(nums[:mid+1]),
		solve(nums[mid+1:]),
		leftMax+rightMax,
	)
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
	1. 这是最大子序和的分治写法，时间复杂度为O(nlogn)

*/
