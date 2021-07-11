package main

import "fmt"

// 空间复杂度大的
//func minimumTotal(triangle [][]int) int {
//	if len(triangle)==0 {
//		return 0
//	}
//	slice:=make([][]int,0)
//	for i:=len(triangle)-1;i>=0 ;i--  {
//		if i == len(triangle)-1{
//			slice=append(slice, triangle[i])
//			fmt.Println(slice)
//			continue
//		}
//		x:=len(triangle)-i-1
//
//		for t:=0;t<len(triangle[i]) ;t++  {
//
//			if len(slice)==x {
//				slice = append(slice, []int{})
//			}
//
//			if slice[x-1] [t] < slice[x-1] [t+1] {
//				slice[x] = append(slice[x], slice[x-1] [t] + triangle[i][t])
//			}else{
//				slice[x] = append(slice[x], slice[x-1] [t+1]+ triangle[i][t])
//			}
//		}
//	}
//	fmt.Println(slice)
//	return slice[len(triangle)-1][0]
//}

// 空间复杂度O(n)的
func minimumTotal(triangle [][]int) int {
	if len(triangle)==0 {
		return 0
	}
	slice:=make([]int,len(triangle))
	copy(slice, triangle[len(triangle)-1])

	for i:=len(triangle)-2;i>=0 ;i--  {
		for t:=0;t<len(triangle[i]) ;t++  {

			if slice[t] < slice[t+1] {
				slice[t] =  slice[t] + triangle[i][t]
			}else{
				slice[t] =  slice[t+1]+ triangle[i][t]
			}
		}
	}
	fmt.Println(slice)
	return slice[0]
}


func main() {
}


/*
	总结
	1. 该题是数塔问题
	2. 题目要求空间复杂度O(n)，第一种方法显然大了，所以我要思考怎么缩小空间复杂度
	3. 采用第二种方法，类似滚轮的赋值(因为该层只与上层有关)，可以把空间复杂度降低到O(n)
*/