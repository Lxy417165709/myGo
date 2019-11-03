package main

import "fmt"

func thirdMax(nums []int) int {
	m:=make(map[int]int)
	slice:=make([]int,0)
	for i:=0;i< len(nums);i++  {
		if m[nums[i]]==0 {
			slice = append(slice, nums[i])
		}
		m[nums[i]]++
	}

	return fastSelect(slice,0, len(slice)-1,3)
}


func fastSelect(nums []int,left int,right int,k int) int{


	base:=nums[left]
	l,r:=left,right
	for ;l<=r ;  {
		for ;l<=r && nums[l]<=base;  {
			l++
		}
		for ;l<=r && nums[r]>=base;  {
			r--
		}
		if l<r{
			nums[l],nums[r] = nums[r],nums[l]
			l++
			r--
		}
	}
	// 表示第k大不存在
	if r==-1 {
		return fastSelect(nums,0, len(nums)-1,1)
	}
	nums[r],nums[left] = nums[left],nums[r]
	if r== len(nums)-k {
		return nums[r]
	}else{
		if r>len(nums)-k {
			return fastSelect(nums,left,r-1,k)
		}else{
			return fastSelect(nums,r+1,right,k)
		}
	}


}
func main() {
	fmt.Println(thirdMax([]int{
		4,2,1,5,10,5,5,5,5,5,5,3,6,
	}))
}

/*
	总结
	1. 该题用到了快速选择算法，快速选择算法只能算无重复元素数组中的第n大/小
*/