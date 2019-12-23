package main

import (
	"fmt"
	"math/rand"
	"sort"
)



func orderSearch(A []int, val int) (int, int) {
	times := 0
	for i := 0; i < len(A); i++ {
		times++
		if A[i] == val {
			return i, times
		}
	}
	return -1, times
}
func binarySearch(A []int, val int) (int, int) {
	l, r := 0, len(A)-1
	times := 0
	for l <= r {
		mid := l + (r-l)/2
		//mid := (l+r)/2
		times++
		if A[mid] == val {
			return mid, times
		} else {
			if A[mid] > val {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return -1,times

}

// 计算平均查找次数
func count(A []int,searchFunction func([]int,int) (int,int)) float64{
	if len(A)==0 {
		return 0
	}

	ans:=0.0
	for i:=0;i< len(A);i++  {
		_,times:=searchFunction(A,A[i])
		ans+=float64(times)
		//fmt.Println(index,i)
	}
	return ans/ float64(len(A))
}
const limit  = 500000000
func generate(n int) []int{
	sli:=make([]int,0)
	for i:=0;i<n ;i++  {
		sli = append(sli, rand.Intn(limit)-limit/2)
	}
	return sli
}



func main() {
	A:=generate(10000)
	sort.Ints(A)
	fmt.Println("顺序查找：",count(A,orderSearch))
	fmt.Println("二分查找：",count(A,binarySearch))
}
