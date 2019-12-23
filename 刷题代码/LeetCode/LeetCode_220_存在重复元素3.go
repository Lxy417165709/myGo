package main

import (
	"fmt"
	"sort"
)

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	arr := make([]Node,0)

	for i:=0;i<len(nums);i++{
		arr =append(arr,Node{nums[i],i})
	}
	sort.Slice(arr,func (i,j int) bool{
		return arr[i].num<arr[j].num
	})
	for i:=0;i<len(arr);i++{
		target := arr[i].num+t
		idx:=search(arr[i:],target) + i
		for j:=idx;j>i;j--{
			if abs(arr[j].index-arr[i].index)<=k{
				fmt.Println(arr[j].index,arr[i].index,arr[j].index-arr[i].index,k)
				return true
			}
		}
	}
	return false


}
type Node struct{
	num int
	index int
}
func abs(a int) int{
	if a<0{
		return -a
	}
	return a
}
func search(arr []Node,target int) int{
	l:=0
	r:=len(arr)-1
	for l<=r{
		mid := (l+r)/2
		if arr[mid].num==target{
			l = mid + 1
		}else{
			if arr[mid].num>target{
				r = mid - 1
			}else{
				l = mid + 1
			}
		}
	}
	return r
}
func main() {


}

/*
	总结
	1. 这题目我是使用二分搜索AC的，思路是先把每个元素封装成一个Node,里面有数值与对应索引
		排序Node数组后，在Node数组二分查找 小于等于x+t的最大索引idx
		之后在在(i,idx] (i为Node数组的第i-1个Node)中判断是不是满足，如果不满足则判断
		下一个Node，否则返回true
	2. 第一次做的时候想到了使用暴力，我觉得也可以解，然后之后想到了二分，二分也AC了
       但是二分感觉和暴力也差不了多少 0.0 因为要在(i,idx]判断
	3. 看了官方后，它有使用二叉搜索数做的，还有用桶做的，时间复杂度降低了好多 0.0
		最快的是用桶来做 0.0..
*/
