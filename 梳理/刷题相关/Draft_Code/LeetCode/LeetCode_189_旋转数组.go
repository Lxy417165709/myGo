package main

import "fmt"

func rotate(nums []int, k int)  {

	if len(nums)==0 {
		return
	}
	k = k% len(nums)
	length:= len(nums)
	index:=0
	for i:=length-k;i<length ;i++  {
		tmp:=nums[i]
		for t:=length-k+index;t>=index+1 ;t--  {
			nums[t] = nums[t-1]
		}
		nums[index] = tmp
		index++
	}
	fmt.Println(nums)
}
func main() {
	rotate([]int{
		-1,23,3,10,12,15,
	},2)
}
/*
	总结
	1. 在没开始做的时候想到了反转，但是反转后的具体操作错了
	2. AC这题使用了O(nk)空间O(1)的算法，这是暴力中最麻烦的方法 ....
	3. 官方有一些更好的方法 0.0.比如反转
	4. 反转思想：先让k位元素到正确的位置(不是每一位元素)，之后对k位反转，和后面的n-k位转，就可以达到目的了
*/