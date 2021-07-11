package main

import "fmt"

//type arr [][]int
//
//func (a arr)Len() int{
//	return len(a)
//
//}
//func (a arr)Less(i,j int) bool{
//	if a[i][0]==a[j][0] {
//		return a[i][1]<a[j][1]
//	}else{
//		return a[i][0]<a[j][0]
//	}
//}
//func (a arr)Swap(i,j int){
//	a[i],a[j] = a[j],a[i]
//}


func corpFlightBookings(bookings [][]int, n int) []int {

	change:=make(map[int]int)

	for i:=0;i< len(bookings);i++  {
		change[bookings[i][0]] += bookings[i][2]
		change[bookings[i][1]+1] -= bookings[i][2]
	}
	ans:=make([]int,1)
	for i:=1;i<=n ;i++  {
		ans = append(ans, ans[i-1] + change[i])
	}
	return ans[1:]
}


func main() {
	fmt.Println(corpFlightBookings([][]int{
	},4))
}

/*
	总结
	1. 第一次做的时候想到了树状数组，线段树，二分查找..但是感觉太复杂了
	2. 看了官方题解之后，我们可以记录每次的变化量 0.0 这样下一个航班的人就等于该航班加上变化量了 0.0..
*/