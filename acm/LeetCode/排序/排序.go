package main

import (
	"fmt"
	"math/rand"
)



func generate(n int) []int {
	sli := make([]int, 0)
	for i := 0; i < n; i++ {
		sli = append(sli, rand.Intn(limit)-limit/2)
	}
	return sli
}
func test(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}

func bubbleSort(A []int) ([]int, int) {
	times := 0
	for i := 0; i < len(A); i++ {
		flag := false
		for t := 0; t < len(A)-i-1; t++ {
			if A[t] > A[t+1] {
				A[t], A[t+1] = A[t+1], A[t]
				flag = true
			}
			times++
		}
		if !flag {
			break
		}
	}
	return A, times
}



func selectSort(A []int) ([]int, int) {
	times := 0
	for i := 0; i < len(A); i++ {
		index := i
		mn := A[index]
		for t := len(A) - 1; t > i; t-- {
			if mn > A[t] {
				index = t
				mn = A[index]
			}
			times++
		}
		A[index], A[i] = A[i], A[index]
	}
	return A, times
}

func insertSort(A []int) ([]int, int) {
	times := 0
	for i := 0; i < len(A); i++ {
		for t := i; t >= 1 && A[t] < A[t-1]; t-- {
			A[t], A[t-1] = A[t-1], A[t]
			times++
		}
	}

	return A, times
}

// 不稳定的插入排序 (不使用交换，而是找到位置后，推移元素)
func insertSort2(A []int) ([]int, int) {
	times := 0
	for i := 0; i < len(A); i++ {
		index := i
		base := A[i]
		for index >= 0 && A[index] >= A[i] {
			index--
			times++
		}
		for t := i - 1; t >= index+1; t-- {
			A[t+1] = A[t]
			times++
		}
		A[index+1] = base
	}

	return A, times
}

// 二分插入排序，也是不稳定的(不使用交换，而是找到位置后，推移元素)
func binaryInsertSort(A []int) ([]int, int) {
	times := 0
	for i := 0; i < len(A); i++ {
		index := findPlace(A[:i+1], A[i], &times) // 注意这里 A[:i+1]，不是对整个数组进行二分查找
		base := A[i]
		for t := i - 1; t >= index+1; t-- {
			A[t+1] = A[t]
			times++
		}
		A[index+1] = base
	}

	return A, times
}
func findPlace(A []int, target int, times *int) int {
	l := 0
	r := len(A) - 1
	for l <= r {

		mid := (l + r) / 2
		if A[mid] == target {
			r = mid - 1
		} else {
			if A[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		*times++
	}
	return r
}

func quickSort(A []int) ([]int, int) {
	times := 0
	quickSortExc(A, 0, len(A)-1, &times)
	return A, times
}
func quickSortExc(A []int, l int, r int, times *int) {
	if l >= r {
		return
	}
	L, R := l, r
	base := A[L]
	for l <= r {
		// 注意点... 里面的指针别写错了啊！
		for l <= r && A[l] <= base {
			l++
			*times++
		}
		for l <= r && A[r] >= base {
			r--
			*times++
		}
		if l <= r {
			A[l], A[r] = A[r], A[l]
		}
	}
	A[L], A[r] = A[r], A[L]
	// 注意点，这不是二分，不是取中点..
	quickSortExc(A, L, r-1, times)
	quickSortExc(A, r+1, R, times)
}

func mergeSort(A []int) ([]int, int) {
	times := 0
	mergeSortExc(A, 0, len(A)-1, &times)
	return A, times
}
func mergeSortExc(A []int, l int, r int, times *int) {
	if l == r {
		return
	}
	mid := (l + r) / 2
	mergeSortExc(A, l, mid, times)
	mergeSortExc(A, mid+1, r, times)
	merge(A, l, mid, mid+1, r, times)
}
func merge(A []int, l1 int, r1 int, l2 int, r2 int, times *int) {
	L := l1
	tmp := make([]int, 0)
	for l1 <= r1 && l2 <= r2 {
		if A[l1] < A[l2] {
			tmp = append(tmp, A[l1])
			l1++
		} else {
			tmp = append(tmp, A[l2])
			l2++
		}
		*times++
	}
	for l1 <= r1 {
		tmp = append(tmp, A[l1])
		l1++
		*times++
	}
	for l2 <= r2 {
		tmp = append(tmp, A[l2])
		l2++
		*times++
	}

	for i := 0; i < len(tmp); i++ {
		A[L+i] = tmp[i]
		*times++
	}

}


type Heap struct {
	heap  []int
	len   int
	times int // 用来测试排序效率
	flag  int // 用来指明堆的性质，0为小顶，1为大顶
}

func initHeap(A []int, flag int) *Heap {
	len := len(A)
	H := &Heap{make([]int, len), len, 0, flag}
	copy(H.heap, A)
	H.heap = append([]int{0}, H.heap...) // 由于下标要从1开始，所以加了个0占位
	H.build()
	return H
}
func (H *Heap) build() {
	for i := H.len / 2; i >= 1; i-- {
		H.downAdjust(i, H.len)
	}
}
func (H *Heap) cmp(a, b int) bool {
	if H.flag == 0 {
		return a > b
	} else {
		return a < b
	}
}

func (H *Heap) downAdjust(low, high int) {
	i, j := low, low*2
	for j <= high {
		H.times ++
		// 次序不能倒
		if j+1 <= high && H.cmp(H.heap[j+1], H.heap[j]) {
			j = j + 1
		}
		// 次序不能倒
		if H.cmp(H.heap[j], H.heap[i]) {
			H.heap[i], H.heap[j] = H.heap[j], H.heap[i]
			i = j
			j = i * 2
		} else {
			break
		}
	}
}
func (H *Heap) upAdjust(low, high int) {
	i, j := high, high/2
	for j >= low {

		// 次序不能倒,这里和downAdjust调换了
		if H.cmp(H.heap[i], H.heap[j]) {
			H.heap[i], H.heap[j] = H.heap[j], H.heap[i]
			i = j
			j = i / 2
		} else {
			break
		}
	}
}

func (H *Heap) insert(number int) {
	H.heap = append(H.heap, number)
	H.len++
	H.upAdjust(1, H.len)
}
func (H *Heap) deleteTop() {
	H.heap[1], H.heap[H.len] = H.heap[H.len], H.heap[1]
	H.downAdjust(1, H.len)
	H.len--
}
func (H *Heap) heapSort() ([]int, int) {

	for i := H.len; i > 1; i-- {
		H.times++
		H.heap[1], H.heap[i] = H.heap[i], H.heap[1]
		H.downAdjust(1, i-1)
	}
	times := H.times
	H.times = 0
	return H.heap[1:], times
}


const limit = 500000

func main() {

	A, times := make([]int, 0), 0
	A, times = bubbleSort(generate(10000))
	fmt.Println("冒泡排序：", test(A), times)
	A, times = selectSort(generate(10000))
	fmt.Println("选择排序：", test(A), times)
	A, times = insertSort(generate(10000))
	fmt.Println("插入排序：", test(A), times)
	A, times = insertSort2(generate(10000))
	fmt.Println("插入排序2：", test(A), times)
	A, times = binaryInsertSort(generate(10000))
	fmt.Println("二分插入排序：", test(A), times)

	A, times = quickSort(generate(10000))
	fmt.Println("快速排序：", test(A), times)
	A, times = mergeSort(generate(10000))
	fmt.Println("归并排序：", test(A), times)

	H := initHeap(generate(10000), 0)
	A, times = H.heapSort()
	fmt.Println("堆排序：", test(A), times)
}

/*
	总结
	1. 数据越集中，快速排序越慢...
			比如：limit = 50,     n = 1000000时，快速排序times = 8635460444
			而  ：limit = 500000, n = 1000000时，快速排序times = 26109775
    2. 数据的集中性不会影响归并排序
			比如：limit = 50,     n = 1000000时，归并排序times = 39902848
			而  ：limit = 500000, n = 1000000时，快速排序times = 39902848
	3. 数据越集中，堆排序越快
			比如：limit = 50,     n = 1000000时，归并排序times = 19119318
			而  ：limit = 500000, n = 1000000时，快速排序times = 19396678
*/
