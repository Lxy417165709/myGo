package main

import (
	"fmt"
	"sort"
)





// 要求A数组是升序排序的,这个函数的返回值是能组成的 值的二元组
func twoSum(A []int,sum int) [][]int{
	sort.Ints(A)	// 注意在函数体外就应该完成排序
	l,r:=0, len(A)-1
	slice:=make([][]int,0)
	for ;l<r ;  {
		tmp:=A[l]+A[r]
		if sum==tmp {
			slice = append(slice, []int{
				A[l],A[r],
			})
			a,b:=A[l],A[r]
			for ;l<r && A[l]==a ;  {
				l++
			}
			for ;l<r && A[r]==b ;  {
				r--
			}
		}else{
			if sum>tmp {
				l++
			}else{
				r--
			}
		}
	}
	return slice
}


func main() {
	fmt.Println(twoSum([]int{
		2,7,11,15,7,7,7,7,0,-5,5,-4,4,4,-4,-2,-2,2,5,-2,0,
	},18))
}
