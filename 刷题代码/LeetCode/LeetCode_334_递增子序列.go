package main

import "fmt"



func increasingTriplet(nums []int) bool {
	first,second:=100000000000,100000000000
	for i:=0;i< len(nums);i++  {
		if nums[i]<first {
			first = nums[i]
		}else{
			if nums[i]>first && nums[i]<second {
				second = nums[i]
			}
			if nums[i]>second {
				return true
			}
		}
	}
	return false
}
func main() {
	fmt.Println(increasingTriplet([]int{
		2147483646,2147483647,2147483647,
	}))
}
/*
	总结
	1. 第一次想的时候想出了O(n)时空复杂度的，就是记录左边最小和右边最大，之后与自己做对比就OK了，然后感觉和接雨水有相似之处
	2. 看了题解后，题解是只需要找出该数组左边的两个数，满足比该数小就OK了
	3. 第一次解答错是因为inf定义的不够大 0.0.
	4. 官方还有通用解，这是一个特殊解0.0. 它是使用数组来存储first,second，空间复杂度也满足要求 0.0.
*/