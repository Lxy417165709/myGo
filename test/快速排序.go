package main

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSort(arr []int64,l int,r int){
	sort(arr,l,r)
}
func sort(arr []int64,l int,r int){

	if l >= r {
		return
	}

	index := l
	eindex := r
	for l<r{
		for arr[index]<=arr[r] && l < r{
			r--
		}
		for arr[index]>=arr[l] && l < r{
			l++
		}
		arr[l],arr[r] = arr[r],arr[l]
	}
	arr[index],arr[r] = arr[r],arr[index]
	sort(arr,index,r-1)
	sort(arr,r+1,eindex)
}
func test(arr []int64) bool{

	for i:=0;i<len(arr)-1;i++{
		if arr[i]>arr[i+1]{
			//fmt.Println(arr[i],arr[i+1])
			return false
		}
	}
	return true
}
func main() {
	rand.Seed(time.Now().Unix())	// 把秒时间戳作为种子


	var arr [1000000]int64
	for i:=0;i<1000000;i++{
		arr[i] = rand.Int63n(8000000) - 4000000
	}
	begin := time.Now().UnixNano()
	quickSort(arr[:],0,len(arr)-1)
	end := time.Now().UnixNano()
	fmt.Println(test(arr[:]))
	fmt.Println(float64(end-begin)/10e9)
}

/*
	总结
	1. 这个方法可以!
*/