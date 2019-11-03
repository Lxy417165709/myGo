package main

import (
	"math"
	"sort"
)

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k<0 || t<0{
		return false
	}

	mp := make(map[int]int)
	size:=t+1
	for i:=0;i<len(nums);i++{
		seq:=int(math.Floor(float64(nums[i])/float64(size))) // 让正负数统一，比如size==4, -3应该放在-1号桶，而不是0,3才是放在0号
		if _,ok:=mp[seq];ok{
			// 一个桶只存一个数，因为如果是两个数的话，那说明二者的值相差都≤t
			//	于是返回true
			return true
		}
		mp[seq] = nums[i]


		// 看看最靠近自己且比自己小的桶的值是不是在合适的范围内
		if x,ok:=mp[seq-1];ok{
			if abs(x-nums[i]) <= t{
				return true
			}
		}
		// 看看最靠近自己且比自己大的桶的值是不是在合适的范围内
		if x,ok:=mp[seq+1];ok{
			if abs(x-nums[i]) <= t{
				return true
			}
		}
		// 删除旧桶，保持我们所有的桶中的元素下标相差≤k
		if i>=k{
			delete(mp,(nums[i-k])/size)
		}
	}
	return false
}


func abs(a int) int{
	if a<0{
		return -a
	}
	return a
}
/*
	总结
	1. 这是利用桶来AC的 0.0
*/