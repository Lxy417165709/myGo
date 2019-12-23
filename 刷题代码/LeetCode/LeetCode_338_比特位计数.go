package main

import "fmt"


// 这个规律是错的..
//func countBits(num int) []int {
//	// 2的一次方,1<<1
//
//	if num==0 {
//		return []int{0}
//	}
//	if num==1 {
//		return []int{0,1}
//	}
//	arr := make([]int,0)
//
//	arr = append(arr, 0,1)
//	for i:=1;i<=30 ;i++  {
//		slice:=get(i)
//		for t:=0;t< len(slice);t++  {
//			arr = append(arr, slice[t])
//			if len(arr)==num+1 {
//				return arr
//			}
//		}
//		arr = append(arr, i+1)
//		if len(arr)==num+1 {
//			return arr
//		}
//	}
//
//	return arr
//}
//func get(n int) []int{
//	arr := make([]int,0)
//	for i :=0;i<n ;i++  {
//		for t:=1;t<= 1<<uint8(i) ;t++  {
//			arr = append(arr, i+1)
//		}
//	}
//	return arr
//}

func countBits(num int) []int {
	// 2的一次方,1<<1

	arr :=make([]int,0)
	arr = append(arr, 0)

	for ;num+1> len(arr);  {
		l := len(arr)
		for i:=0;i<l && num+1> len(arr);i++  {
			arr = append(arr, arr[i]+1)
		}
	}

	return arr
}

func main() {
	fmt.Println(countBits(12))
}

/*
	总结
	1. 第一次做的时候规律弄错了(列举的时候粗心了)
	2. 之后正确列举后找到了规律
	3. 官方有用动态规划做的
*/