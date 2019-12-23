package main

import (
	"fmt"
	"sort"
)

func findKthLargest(nums []int, k int) int {
	//m:=make(map[int]int)

	sort.Ints(nums)
	return nums[len(nums)-k]
}
func main() {
	fmt.Println(findKthLargest([]int{
		3,
	},1))
}

/*
	总结
	1. 该题我使用排序完成
	2. 其实还可以用快速选择算法，(可以接纳重复元素，只是对于1,2,2,5,那2既是第二大也是第三大)
	3. 官方有用堆做出来的 0.0
*/
