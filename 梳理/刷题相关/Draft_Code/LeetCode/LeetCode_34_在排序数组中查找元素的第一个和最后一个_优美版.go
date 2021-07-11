package main


func searchRange(nums []int, target int) []int {
	idx := binarySearch(nums,target)
	if idx == -1{
		return []int{-1,-1}
	}
	L := firstGreaterOrEqual(nums,target)
	R := lastLessOrEqual(nums,target)
	return []int{L,R}
}

// 在数组中找到等于目标值的数的索引，没有则返回 -1
func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		// mid := (l + r)/2 这样写也可以，但是l + r很大时可能造成溢出。
		// 所以我采用 mid := l + (r-l)/2 这种写法防止溢出。
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		} else {
			if nums[mid] > target {
				r = mid - 1
			}else{
				l = mid + 1
			}
		}
	}
	return -1
}

// 在数组中找到第一个大于等于目标值的数的索引，没有则返回 数组长度
func firstGreaterOrEqual(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			r = mid - 1
		} else {
			if arr[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return l
}
// 在数组中找到最后一个小于等于目标值的数的索引，没有则返回 -1
func lastLessOrEqual(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			l = mid + 1
		} else {
			if arr[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	// firstGreater 返回 l, 而 lastLessOrEqual 返回r。
	return r
}
