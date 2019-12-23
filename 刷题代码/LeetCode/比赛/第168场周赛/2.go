package main

import "fmt"

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k!=0{
		return false
	}
	count := make(map[int]int)
	for i:=0;i<len(nums);i++{
		count[nums[i]]++
	}

	for len(count)!=0{
		in := 10000000000

		for k := range count{
			in = min(in,k)
		}
		fmt.Println(in)
		for i:=0;i<k;i++{
			count[in]--
			if count[in]==0{
				delete(count,in)
			}
			if count[in]<0{
				return false
			}
			in++
		}
	}
	return true
}

func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}

func main() {
	fmt.Println(isPossibleDivide([]int{1,2,3,4},3))
}