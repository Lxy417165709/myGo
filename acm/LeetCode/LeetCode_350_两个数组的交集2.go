package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	i,j:=0,0
	sort.Ints(nums1)
	sort.Ints(nums2)
	ans:=make([]int,0)
	for ;i< len(nums1) && j<len(nums2);  {
		if nums1[i]==nums2[j] {
			ans = append(ans, nums1[i])
			i++
			j++
		}else{
			if nums1[i]>nums2[j]{
				j++
			}else{
				i++
			}
		}
	}
	return ans
}



func main() {
	fmt.Println(intersect([]int{
		5,
	},[]int{
		4,
	}))
}
/*
	总结
	1. 这是一道简单题
	2. 不知道写啥..
*/