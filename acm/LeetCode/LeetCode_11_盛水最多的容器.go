package main

import (
	"fmt"
	"math"
)

//func maxArea(height []int) int {
//	ans:=0
//
//	for i:=0;i< len(height);i++  {
//		for t:=len(height)-1;t>i ;t--  {
//			if height[t]>=height[i] {
//				ans = int(math.Max(float64(ans),float64((t-i) * height[i])))
//				break
//			}else{
//				ans = int(math.Max(float64(ans),float64((t-i) * height[t])))
//			}
//		}
//		//fmt.Println(i,ans)
//	}
//	return ans
//}

func maxArea(height []int) int {
	ans:=0

	for l,r:=0,len(height)-1; l<r ;  {
		ans = int(math.Max(float64(ans), float64(r-l)* math.Min(float64(height[l]),float64(height[r]))))
		if height[l]<height[r] {
			l++
		}else{
			r--
		}
	}
	return ans
}


func main() {
	fmt.Print(maxArea([]int{1,1}))
}


/*
	总结
	1. 该题目可以用暴力法 + 剪枝，但是时间复杂度还是O(n^2)
	2. 官方题解用了双指针法(我是看的...)
	3. 这种题一定有个必然条件，比如height[l]<height[r]时，一定是移动l，r是不能移动的，可以证明 0.0，反之是移动r

*/