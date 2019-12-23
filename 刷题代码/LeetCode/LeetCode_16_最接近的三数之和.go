package main

import (
	"fmt"
	"sort"
)

func search_r_l(slice []int,target int) int{
	l:=0
	r:= len(slice)-1

	for ;l<=r ;  {
		mid := (l+r)/2
		if slice[mid] == target {
			r = mid - 1	// 可修改处1
		}else{
			if slice[mid] > target{
				r = mid -1
			}else{
				l = mid + 1
			}
		}
	}
	return l // 可修改处2
}
func abs(a int) int{
	if a<0 {
		a = -a
	}
	return a
}

func judge(target,ans,a int) int{
	if abs(target-ans)>abs(target-a) {
		return a
	}else{
		return ans
	}
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	ans:=target+100000000
	for i:=0;i< len(nums)-2;i++  {
		for t:=i+1;t<len(nums)-1 ;t++  {
			sum:=nums[i] + nums[t]
			index:=search_r_l(nums[t+1:],target-sum)

			if index!= len(nums[t+1:]) {
				if index==0 {
					a:=nums[i] + nums[t] + nums[t+1+index]
					ans = judge(target,ans,a)
				}else{
					a:=nums[i] + nums[t] + nums[t+1+index]
					b:=nums[i] + nums[t] + nums[t+index]
					ans = judge(target,ans,a)
					ans = judge(target,ans,b)
				}
			}else{
				b:=nums[i] + nums[t] + nums[t+index]
				ans = judge(target,ans,b)
			}


		}
	}
	return ans
}


func main() {
	fmt.Println(threeSumClosest([]int{
		0,0,0,
	},1))
}
/*
	总结
	1. 我的思路是先排序，然后组合2个数，再到数组剩下的元素中用二分查找找到最适合的数 0.0.时间复杂度O(n^2logn)
	2. 官方用双指针法，做法更简洁 0.0.而且时间复杂度更低，是O(n^2)  可以先找一个数，再对剩下的数双指针 0.0.

*/