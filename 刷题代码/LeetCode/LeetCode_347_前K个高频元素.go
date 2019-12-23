package main

import (
	"fmt"
	"sort"
)

type Node struct{
	times int
	num int
}
type arr []Node

func (a arr)Len() int{
	return len(a)
}
func (a arr)Less(i,j int) bool{
	return a[i].times<a[j].times
}
func (a arr)Swap(i,j int){
	a[i],a[j] = a[j],a[i]
}


func topKFrequent(nums []int, k int) []int {
	slice := make([]int, 0)
	m := make(map[int]int)
	timeSlice := make([]Node, 0)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == 0 {
			slice = append(slice, nums[i])
		}
		m[nums[i]]++
	}

	for i := 0; i < len(slice); i++ {
		timeSlice = append(timeSlice, Node{m[slice[i]],slice[i]})

	}
	sort.Sort(arr(timeSlice))
	length := len(timeSlice)
	ans := make([]int, 0)
	for i := length - 1; i >= length-k; i-- {
		ans = append(ans, timeSlice[i].num)
	}
	return ans

}

func main() {
	fmt.Println(topKFrequent([]int{
		1,1,1,2,2,3,
	},2))
}

/*
	总结
	1. 该题我是用哈希排序AC的
	2. 官方有用堆AC的方法 0.0.
*/