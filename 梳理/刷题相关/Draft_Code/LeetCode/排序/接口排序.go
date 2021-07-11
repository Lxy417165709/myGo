package main

import (
	"fmt"
	"math/rand"
)

type Sorter interface {
	Len() int
	Less(i,j int) bool
	Swap(i,j int)
}
type intSlice []int

func (this intSlice) Len() int{
	return len(this)
}
func (this intSlice) Less(i,j int) bool{
	return (this)[i]<(this)[j]
}
func (this intSlice) Swap(i,j int){
	(this)[i],(this)[j] = (this)[j],(this)[i]
}

func bubbleSortSorter(arr Sorter){
	for i:=0;i<arr.Len();i++{
		for t:=i;t<arr.Len()-i-1;t++{
			if !arr.Less(t,t+1){
				arr.Swap(t,t+1)
			}
		}
	}
}

func generate(n int) []int {
	sli := make([]int, 0)
	for i := 0; i < n; i++ {
		sli = append(sli, rand.Intn(limit)-limit/2)
	}
	return sli
}
func test(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}
const limit = 500000
func main() {
	A := generate(10)
	fmt.Println(A)
	bubbleSortSorter(intSlice(A))
	fmt.Println(A)
}
/*

	总结
	1. 只要实现了Sorter接口就可以实现排序 0.0
	2. intSlice实现了Sorter接口，其中intSlice本质是[]int
*/