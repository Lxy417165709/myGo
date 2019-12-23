package main

import (
	"fmt"
	"sort"
)

func search(A []int, target int) int {
	l, r := 0, len(A)-1
	for ; l <= r; {
		mid := (l + r) / 2
		if A[mid] == target {
			return mid
		} else {
			if A[mid] > target {
				r--
			} else {
				l++
			}
		}
	}
	return -1

}

func intersection(nums1 []int, nums2 []int) []int {
	m :=make(map[int]int)
	ans:=make([]int,0)
	sort.Ints(nums2)
	for i:=0;i< len(nums1);i++  {
		if m[nums1[i]]==0 {
			if search(nums2,nums1[i])!=-1 {
				ans = append(ans, nums1[i])
			}
		}
		m[nums1[i]]=1
	}
	return ans
}

func main() {
	fmt.Println(intersection([]int{

	},[]int{
		9,4,9,8,4,5,
	}))
}
/*
	总结
	1. 这是一道简单题
	2. 先对nums2排序，之后遍历nums1元素，看是否存在nums2中
	3. 时间复杂度O(nlogm)
*/