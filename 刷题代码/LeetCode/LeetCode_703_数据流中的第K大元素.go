package main

import "sort"

type KthLargest struct {
	arr []int
	k int
}


func Constructor(k int, nums []int) KthLargest {
	sort.Ints(nums)
	return KthLargest{nums,k}
}


func (this *KthLargest) Add(val int) int {
	this.arr = append(this.arr,val)
	// 将元素插入数组，最差复杂度是O(n)
	for i:=len(this.arr)-1;i>=1 && this.arr[i]<this.arr[i-1];i--{
		this.arr[i],this.arr[i-1] = this.arr[i-1],this.arr[i]
	}
	return this.arr[len(this.arr)-this.k]
}

/*
	总结
	1. 我第一次想到的思路是维护一个大小为k的小顶堆，这样每次第一个元素就是第k大元素
	2. 由于没有标准容器的支持，于是我采用另外一种方法，就是先对数组排序，插入的时候采用插入排序
		由排序后的数组，很容易得到第k大。
*/