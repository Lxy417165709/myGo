package 双指针

import "fmt"

func trap(height []int) int {
	left,right := 0,len(height)-1
	leftMax,rightMax := 0,0
	sum := 0
	for left<=right{
		if leftMax>=rightMax{
			rightMax = max(rightMax,height[right])
			sum+=(rightMax-height[right])
			right--
		}else{
			leftMax = max(leftMax,height[left])
			sum+=(leftMax-height[left])
			left++
		}
	}
	return sum
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

/*
	总结
	1. 之前做的时候，采用的方法时空复杂度都为O(n)
	2. 这次看了官方题解后，采用了时空复杂度O(n),O(1)的解法
		为什么可以优化成O(1)呢？ 因为当我们知道了左边的最大值和右边的最大值时，如果右边最大值更大，
		那么显然我们让左边向中心推移并计算雨水值就可以了(注意再此过程中还要更新左边最大值)，如果左边
		最大值更大，那么我们就让右边向中心靠拢.....这样我们就可以只用2个指针,left,right就能实现接雨水了。
	3. 规律: 当我们要求 sum(min(lmax[i],rmax[i]))、sum(max(lmin[i],rmin[i]))的时候，我们可以采用这种
			 渐进双指针的方式优化时间复杂度。


*/