package main

import "fmt"

func search_r_l(slice []int, target int) int {
	l := 0
	r := len(slice) - 1

	for ; l <= r; {
		mid := (l + r) / 2
		if slice[mid] == target {
			r = r - 1 // 可修改处1
		} else {
			if slice[mid] > target {
				r = r - 1
			} else {
				l = l + 1
			}
		}
	}
	return l // 可修改处2
}

func sortedSquares(A []int) []int {
	B := make([]int, 0)
	r := search_r_l(A, 0)
	l := r - 1
	for ; r < len(A) && l >= 0; {
		a,b := A[r]*A[r],A[l]*A[l]
		if a < b {
			B = append(B, a)
			r++
		}else{
			B = append(B, b)
			l--
		}
	}
	for ; r < len(A); {
		B = append(B, A[r]*A[r])
		r++
	}
	for ; l >= 0; {
		B = append(B, A[l]*A[l])
		l--
	}
	return B
}

func main() {
	fmt.Println(sortedSquares([]int{-1050,-898,-450,-1,0,5,10,80}))
}

/*
	总结
	1. 该题目用二分优化了一下
	2. 最主要找到最小的平方值对应的下标，然后用双指针法左右拓展
*/