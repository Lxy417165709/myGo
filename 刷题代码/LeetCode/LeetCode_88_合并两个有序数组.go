package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i := m + n - 1
	m = m - 1
	n = n - 1
	for ;m>=0 && n>=0 ;  {
		if nums1[m] > nums2[n] {
			nums1[i] = nums1[m]
			m--
		}else{
			nums1[i] = nums2[n]
			n--
		}
		i--
	}
	for ;m>=0 ;  {
		nums1[i] = nums1[m]
		m--
		i--
	}
	for ;n>=0 ;  {
		nums1[i] = nums2[n]
		n--
		i--
	}

}


func main() {
	merge([]int{1,2,3,10,15,0,0,0,0,0},5,[]int{2,5,6,8,10},5)
}

/*
	总结
	1. 该题是归并排序的基础变形，因为这题最优的话是双指针从后到前扫
*/