package _1nod

import (
	"fmt"
	"math"
)

func main() {
	arr := [10]int{}
	for i:=1;i<=6;i++{
		fmt.Scan(&arr[i])
	}
	cnt := arr[6]
	m :=make(map[int]int)
	cnt += arr[5]
	m[1]+=11*arr[5]

	cnt += arr[4]
	m[2]+=5*arr[4]

	cnt+= int(math.Ceil(float64(arr[3])/4))
	arr[3] = arr[3]%4
	if arr[3]!=0{
		m[1]+= 8-arr[3]
		m[2]+=(4-arr[3])*2 - 1
	}


	if m[2]<arr[2]{
		arr[2]-=m[2]
		m[2] = 0
		cnt += int(math.Ceil(float64(arr[2])/9))
		arr[2] = arr[2] % 9
		if arr[2] != 0{
			m[1]+=(9-arr[2]%9)*4
		}
	}else{
		m[1] += (m[2]-arr[2]) * 4
	}

	if m[1]<arr[1]{
		arr[1]-=m[1]
		m[1] = 0
		cnt += int(math.Ceil(float64(arr[1])/36))
	}
	fmt.Println(cnt)






}

/*
	总结
	1. 这道题是贪心模拟题...
	2. 我们可以记录空闲位置，而不是先塞东西~
*/