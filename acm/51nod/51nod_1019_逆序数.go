package main

import "fmt"

var ans = 0
var tmp [50050]int

func merge(arr []int, l1 int, r1 int, l2 int, r2 int) {
	length := 0
	i, j := l1, l2
	for ; i <= r1 && j <= r2; {
		if arr[i] > arr[j] {
			tmp[length] = arr[j]
			j++
		} else {
			tmp[length] = arr[i]
			ans += j - l2 + 1 - 1	// 注意!
			i++
		}
		length++
	}

	for ; i <= r1; {
		ans += r2-l2+1
		tmp[length] = arr[i]
		length++
		i++
	}

	for ; j <= r2; {
		tmp[length] = arr[j]
		length++
		j++
	}

	for index := l1; index <= r2; index++ {
		arr[index] = tmp[index-l1]
	}

}

func mergeSort(arr []int,l int, r int) {

	mid := (l+r) / 2

	if  l<r{
		mergeSort(arr,l,mid)
		mergeSort(arr,mid+1,r)
		merge(arr, l, mid, mid+1, r)
	}


}

var arr [50050]int
func main() {

	n:=0
	b:=0
	fmt.Scan(&n,&b)
	fmt.Println(n,b)
	for i:=0;i<n ;i++  {
		fmt.Scan(&arr[i])
	}
	mergeSort(arr[:n],0,n-1)
	fmt.Println(ans)
}



/*
	总结
	1. 对于数组的递归,一般都要用参数变量标识数组的区间

*/