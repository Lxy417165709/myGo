package main

import "fmt"


func search_r_l(slice []int,target int) int{
	l:=0
	r:= len(slice)-1

	for ;l<=r ;  {
		mid := (l+r)/2
		if slice[mid] == target {
			r = mid - 1	// 可修改处1
		}else{
			if slice[mid] > target{
				r = mid -1
			}else{
				l = mid + 1
			}
		}
	}
	return l // 可修改处2
}

func search_r_r(slice []int,target int) int{
	l:=0
	r:= len(slice)-1

	for ;l<=r ;  {
		mid := (l+r)/2
		if slice[mid] == target {
			r = mid - 1
		}else{
			if slice[mid] > target{
				r = mid -1
			}else{
				l = mid + 1
			}
		}
	}
	return r
}

func search_l_l(slice []int,target int) int{
	l:=0
	r:= len(slice)-1

	for ;l<=r ;  {
		mid := (l+r)/2
		if slice[mid] == target {
			l = mid + 1
		}else{
			if slice[mid] > target{
				r = mid -1
			}else{
				l = mid + 1
			}
		}
	}
	return l
}

func search_l_r(slice []int,target int) int{
	l:=0
	r:= len(slice)-1

	for ;l<=r ;  {
		mid := (l+r)/2
		if slice[mid] == target {
			l = mid + 1
		}else{
			if slice[mid] > target{
				r = mid -1
			}else{
				l = mid + 1
			}
		}
	}
	return r
}


func main() {
	slice:=[]int{1,2,2,2,3,5,5,6,6,6}
	x :=0
	for ;true ;  {
		fmt.Scanln(&x)



		fmt.Println("第一个大于等于 		[0,长度]" ,search_r_l(slice,x)) // 0 表示 target<=所有数组元素
																				// 长度表示 target>所有数组元素

		fmt.Println("最后一个小于 		[-1,长度)",search_r_r(slice,x))		// -1 表示 target<=所有数组元素
																				// 长度-1 表示 target>所有数组元素

		fmt.Println("第一个大于 			[0,长度]",search_l_l(slice,x))	// 0 表示 target<所有数组元素
																				// 长度 表示 target>=所有数组元素

		fmt.Println("最后一个小于等于 	[-1,长度)",search_l_r(slice,x))		// -1 表示 target<所有数组元素
																				// 长度-1 表示 target>=所有数组元素

		// 最后一个小于 + 1 	== 第一个大于等于
		// 最后一个小于等于 + 1 == 第一个大于
	}

}
