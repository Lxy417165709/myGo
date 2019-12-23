package main

import "fmt"

func quickSort(A []int){
	if len(A)==0 {
		return
	}
	l:=0
	r:= len(A)-1
	base := A[l]
	// 这必须是l<=r
	for ;l<=r ;  {
		// 这必须是l<=r,A[l]<=base
		for ;l<=r && A[l]<=base ;  {
			l++
		}
		// 这必须是l<=r,A[r]>=base
		for ;l<=r && A[r]>=base ;  {
			r--
		}
		// l<r,l<=r都可以
		if l<=r {
			A[l],A[r] = A[r],A[l]
			//l++
			//r--
		}
	}
	A[0],A[r] = A[r],A[0]
	quickSort(A[0:(l+r)/2])
	quickSort(A[(l+r)/2+1:])
}

// 判断是否有序
func judge(A []int) bool{
	for i:=1;i< len(A);i++  {
		if A[i]<A[i-1] {
			return false
		}
	}
	return true
}


func main() {
	A:=[]int{
		5,
	}
	quickSort(A)
	fmt.Println(A,judge(A))
}
/*
	总结
	1. 快排都用<=就好了 0.0..不然会出错

*/