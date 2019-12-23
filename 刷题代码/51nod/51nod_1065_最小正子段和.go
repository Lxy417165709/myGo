package _1nod

import (
	"fmt"
	"sort"
)

type Node struct {
	index int
	sum   int64
}

var node [50050]Node

type nodeSlice []Node

func (s nodeSlice) Len() int {
	return len(s)
}
func (s nodeSlice) Less(i, j int) bool {
	if s[i].sum != s[j].sum {
		return s[i].sum < s[j].sum
	} else {
		return s[i].index > s[j].index
	}

}
func (s nodeSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func getMin(a int64, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	var sum int64 = 0
	var x int64
	var min int64 = 2<<62-1
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		sum = sum + x
		node[i].index = i
		node[i].sum = sum
	}
	var sli nodeSlice = node[0:n]
	sort.Sort(sli)

	for i := 0; i < len(sli); i++ {
		if sli[i].sum > 0{
			min = getMin(min,sli[i].sum)
		}
	}
	for i := 0; i < len(sli)-1; i++ {
		if sli[i].sum < sli[i+1].sum && sli[i].index < sli[i+1].index {
			if sli[i+1].sum-sli[i].sum < min {
				min = sli[i+1].sum-sli[i].sum
			}
		}
	}
	fmt.Println(min)
}

/*
	知识点
	1. 在序列中找2个差距最小的数可以先排序,再寻找
	2. 两段前缀和的差包括了(0,maxIndex]的所有子段,所以我们还要补上index从0开始的遍历.	(maxIndex == 数组长度 - 1)
*/