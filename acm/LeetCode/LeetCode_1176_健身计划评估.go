package main

import "fmt"

func get(x int,lower int,upper int) int{
	if x>upper {
		return 1
	}
	if x<lower {
		return -1
	}
	return 0
}
// 未优化
//func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
//	ans:=0
//	for i:=0;i< len(calories)-k+1;i = i + 1  {
//		sum:=0
//		for t:=0;t<k;t++  {
//
//			sum+=calories[i+t]
//
//		}
//		ans += get(sum,lower,upper)
//	}
//	return ans
//}

// 优化
func dietPlanPerformance(calories []int, k int, lower int, upper int) int {
	ans:=0
	sum:=0
	for i:=0;i<= len(calories)-k;i = i + 1  {
		if i==0{
			for t:=0;t<k;t++  {
				sum+=calories[i+t]
			}
		}else{
			sum = sum-calories[i-1] + calories[i+k-1]
		}
		ans += get(sum,lower,upper)
	}
	return ans
}


func main() {
	fmt.Println(dietPlanPerformance([]int{
		6,5,0,0,
	},2,1,5))
}
/*
	总结
	1. 第一次没看清楚题目..他说的是要计算每一天 0.0..
	2. 我真的觉得是，题目的例子误导人了
	3. 算法还能优化，据说是滑动窗口 0.0. 我写在上面了

*/