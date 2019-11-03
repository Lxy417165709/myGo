package main

import "fmt"

// 这种方法只能处理重复数字只有一个,即出现2次的情况
//func findDuplicate(nums []int) int {
//	n := len(nums)-1
//	sum :=(n+1)*(n)/2
//	sum1 :=0
//	for i:=0;i<len(nums) ;i++  {
//		sum1+=nums[i]
//	}
//	return sum1-sum
//}

// 快慢指针法
func findDuplicate(nums []int) int {

	fast,slow:=0,0
	for ;true ;  {
		fast = nums[nums[fast]]
		slow = nums[slow]
		fmt.Println(fast,slow)
		if fast == slow {
			break
		}
	}
	finder :=0
	for ;true ;  {
		finder = nums[finder]
		slow = nums[slow]
		if finder== slow {
			break
		}
	}

	return finder

}


func main() {
	fmt.Println(findDuplicate([]int{1,3,3,4,2}))
}

/*
	总结
	1. 第一次的时候没考虑全，导致了第一种方法的错误，(上面写了)
	2. 看了题解之后发现是快慢指针法... 解释可以看看解答中的 ：快慢指针的解释 [从@Damien_Undo写的题解得到启发]
	3. 还是有不太理解这个做法为什么可以成功 0.0..
*/