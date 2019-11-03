package main

import (
	"fmt"
)

//func balancedString(s string) int {
//	m:=make(map[uint8]int)
//	m['Q']=0
//	m['W']=1
//	m['E']=2
//	m['R']=3
//	arr:=make([]int,4)
//
//	for i:=0;i< len(s);i++  {
//		if judge(int(s[i])) {
//			arr[m[s[i]]]++
//		}
//	}
//	times:= len(s)/4
//	sort.Ints(arr)
//	for i:=0;i<len(arr)/2 ;i++  {
//		arr[i],arr[len(arr)-1-i] = arr[len(arr)-1-i],arr[i]
//	}
//	fmt.Println(arr)
//	ans:=0
//	rongyu:=0
//	for i:=0;i< len(arr);i++  {
//		if arr[i]==times {
//			ans+=0
//		}else{
//			if arr[i]<times {
//				ans+=times-arr[i]
//			}else{
//				rongyu+=arr[i]-times
//				ans+=arr[i]-times
//			}
//		}
//	}
//
//	return ans-rongyu
//
//
//
//}
func balancedString(s string) int {
	m:=make(map[uint8]int)
	m['Q']=0
	m['W']=1
	m['E']=2
	m['R']=3
	arr:=make([]int,4)
	tmp:=make([]int,4)
	for i:=0;i< len(s);i++  {
		arr[m[s[i]]]++

	}

	times:= len(s)/4

	mark:=make([]int,4)
	flag:=0
	for i:=0;i<4 ;i++  {
		if arr[i]>times {
			tmp[i] = arr[i]-times
			mark[i] = 1
			flag=1
		}
	}
	if flag==0 {
		return 0
	}
	y:=make([]int,4)
	fmt.Println(tmp,arr)
	for i:=0;i< len(s);i++  {
		y[m[s[i]]]++

		if judge2(y) {
			for t:=0;t< len(y);t++  {
				y[t]-=1
			}
		}
		fmt.Println(y)
		if judge(tmp,y,mark) {
			ans:=0
			for g:=0;g< len(y);g++  {
				ans+=y[g]
			}
			return ans
		}
	}

	return 0
}



func min(a,b int) int{
	if a>b {
		return b
	}
	return a
}


func judge(a []int,b []int,mark []int) bool{
	for i:=0;i<4 ;i++  {
		if mark[i]==1 && a[i]>b[i] {
			return false
		}
	}
	return true
}
func judge2(a []int) bool{
	for i:=0;i<4 ;i++  {
		if a[i]==0 {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(balancedString("QWER"))
}
/*
	总结
	1. 策略错误了！！这题没AC
*/