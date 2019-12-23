package main

import (
	"fmt"
	"math/rand"
)

// 返回索引为0的元素在数组中的位置索引（从小到大排序）
func randomPartition(A []int,l int,r int) int{
	x:=rand.Intn(r-l+1) + l
	A[l],A[x] = A[x],A[l]
	index:=l
	base :=A[index]
	for l<=r  {
		for l<=r && A[l]<=base {
			l++
		}
		for l<=r && A[r]>=base {
			r--
		}
		if l<=r {
			A[l],A[r] = A[r],A[l]
		}
	}
	A[index],A[r] = A[r],A[index]
	return r
}

// 选择第n小的数(重复的数次序不等)
func randomSelect(A[]int,n int) int{
	n = n - 1	// 第n小 对应的下标是 n-1,比如 第1小 对应的是 下标0
	if n>= len(A) && n<=0 {
		return -1
	}

	x:=randomPartition(A,0, len(A)-1)
	for x!=n  {
		if x>n {
			x=randomPartition(A,0, x-1)
		}else{
			x=randomPartition(A,x+1, len(A)-1)
		}
	}
	return A[x]
}

// 选择第n大的数(重复的数次序不等)
func randomSelect2(A[]int,n int) int{
	return randomSelect(A, len(A)-n+1)
}

// LeetCode 414 第三大的数
func thirdMax(nums []int) int {
	A:=make([]int,0)
	mp:=make(map[int]int)
	for i:=0;i< len(nums);i++  {
		if _,ok :=mp[nums[i]];!ok {
			A = append(A, nums[i])
			mp[nums[i]] = 1
		}
	}
	if len(A)<3 {
		return randomSelect2(A,1)
	}else{
		return randomSelect2(A,3)
	}

}
func main() {
	fmt.Println(thirdMax([]int{1,2,3,4,5,6,7,8,9,10,10,10,10,10,11,11,11,11,11}))
}

// 已AC，


// 第n小 == 第 len(A) + 1 - n 大
// 求得的索引x对应第x+1小，即 第len(A)-x大

// 如果求的是第n小，当随机选择求得的索引为x, 如果 x + 1 == n 则返回，如果 x + 1 > n,则向左，否则向右
// 如果求的是第n大，当随机选择求得的索引为x, 如果 len(A)-x == n 则返回，如果len(A)-x > n,则向右，否则向左
