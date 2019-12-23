package main

import (
	"fmt"
	"sort"
)

type arr [][]int
func (a arr)Len() int{
	return len(a)
}
func (a arr)Swap(i,j int){
	a[i],a[j] = a[j],a[i]
}
func (a arr)Less(i,j int) bool{
	return a[i][0]<a[j][0]
}

func merge(intervals [][]int) [][]int {
	if len(intervals)==0 {
		return [][]int{}
	}
	sort.Sort(arr(intervals))
	fmt.Println(intervals)
	slice:=make([][]int,0)

	b:=intervals[0][0]
	e:=intervals[0][1]

	for i:=0;i< len(intervals);i++  {
		if e>=intervals[i][0] {
			if intervals[i][1]>e {
				e = intervals[i][1]
			}
		}else{
			slice = append(slice, []int{b,e})
			b = intervals[i][0]
			e = intervals[i][1]
		}
	}
	slice = append(slice, []int{b,e})
	return slice

}

func main() {
	fmt.Println(merge([][]int{
	}))
}
/*

	总结
	1. 这是一道简单的区间合并，先排序，再判断，e大于下一个区间的左值时说明可能拓宽....时间复杂度为O(nlogn)
	2. 官方有大佬利用c++的map写出了O(n)的算法，但是我不太懂..(没认真理解)
*/