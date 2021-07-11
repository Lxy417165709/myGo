package main

import "fmt"

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	m:=make(map[int]int)
	ans:=make([]int,0)
	for i:=0;i< len(arr1);i++  {
		m[arr1[i]]++
	}
	for i:=0;i< len(arr2);i++  {
		m[arr2[i]]++
	}
	for i:=0;i< len(arr3);i++  {
		m[arr3[i]]++
		if m[arr3[i]]==3 {
			ans = append(ans, arr3[i])
		}
	}
	return ans
}
func main() {
	fmt.Println(arraysIntersection([]int{
		1,2,3,4,5,
	},[]int{
		1,2,5,7,9,
	},[]int{
		1,3,4,5,8,
	}))
}
