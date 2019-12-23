package main

import (
	"fmt"
	//"sort"
)

//type arr [][]int
//
//func (a arr)Len() int{
//	return len(a)
//}
//func (a arr)Less(i,j int) bool{
//	return a[i][0]<a[j][0]
//}
//
//func (a arr)Swap(i,j int) {
//	a[i],a[j] = a[j],a[i]
//}


func reconstructQueue(people [][]int) [][]int {
	//sort.Sort(arr(people))
	temp:=make([][]int, 0)
	for i:=0;i< len(people);i++  {
		temp = append(temp, []int{})
		temp[i] =make([]int,len(people[i]))
		copy(temp[i], people[i])
	}

	ans:=make([][]int,0)
	x:=len(people)
	for t:=0;t< x;t++  {
		value:= 1000000
		index:=-1
		for i:=0;i< len(people);i++  {
			if people[i][1]==0 {
				if people[i][0]<value {
					value = people[i][0]
					index=i
				}
			}
		}
		if index!=-1 {
			tmp:=make([]int, len(temp[index]))
			copy(tmp,temp[index])
			ans = append(ans, tmp)
			for i:=0;i<len(people) ;i++  {
				if people[i][0]<=people[index][0] && people[i][1]!=0  {
					people[i][1]--
				}
			}
			people[index][0] = 1000000000
		}

	}
	return ans
}
func main() {
	fmt.Println(reconstructQueue([][]int{
		{5,0},{6,1},{5,2},{7,0},{4,4},{7,1},
	}))
}
/*
	总结
	1. 该题代码写多了，而且步骤复杂了，其实还有可以优化的地方，比如使用优先队列0.0..(但是我在Go还没用过 ...)
	2. 官方评论的思路和我差不多，不过我写得复杂了 0.0..
*/