package _1nod

import (
	"fmt"
	"sort"
)
var arr [10050]float64
func main() {
	n,c,d:=0,0.0,0.0
	fmt.Scan(&n,&c,&d)
	for i:=0;i<n;i++{
		fmt.Scan(&arr[i])
	}
	sort.Float64s(arr[:n])
	time:=0.0
	fastTime := c-d
	ans:=0
	for i:=0;i<n;i++{
		var spendTime float64
		if time + arr[i] <= fastTime{
			spendTime = arr[i]
		}else{
			if time <= fastTime {
				half := fastTime - time
				spendTime = half + (arr[i] - half)/2
			}else{
				spendTime = arr[i]/2
			}
		}
		time += spendTime
		if time <= c{
			ans++
		}
		//fmt.Println(time,spendTime)
	}
	fmt.Println(ans)
}

/*
	总结
	1. 这还是一道贪心模拟题~
*/