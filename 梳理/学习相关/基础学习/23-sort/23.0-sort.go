package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name  string
	grade int
}

type studentSlice []Student

func (s studentSlice) Len() int {
	return len(s)
}
func (s studentSlice) Less(i, j int) bool {
	return s[i].grade < s[j].grade
}
func (s studentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	// 1. 对切片排序
	// 1.0 对int排序
	sli1 := []int{8, 12, 56, -4, 8, 7}
	sort.Ints(sli1)
	fmt.Println(sli1)
	// [-4 7 8 8 12 56]

	// 1.1 对float排序
	sli2 := []float64{887.7887, 12, 55.6, -4, 8, 7.8}
	sort.Float64s(sli2)
	fmt.Println(sli2)
	// [-4 7.8 8 12 55.6 887.7887]

	// 1.2 对string排序 (以字典序排序)
	sli3 := []string{"a", "bq", "aee", "bwe", "eee", "bwwww"}
	sort.Strings(sli3)
	fmt.Println(sli3)
	// [a aee bq bwe bwwww eee]

	// 2. 自定义排序(需要继承sort.Interface接口)
	// 1.0-1.2的另外一种写法~
	//sort.Sort(sort.IntSlice(sli1))
	//sort.Sort(sort.Float64Slice(sli2))
	//sort.Sort(sort.StringSlice(sli3))

	// 2.0 逆向排序
	sort.Sort(sort.Reverse(sort.StringSlice(sli3)))
	fmt.Println(sli3)
	// [eee bwwww bwe bq aee a]

	// 2.1 结构体排序
	sli4 := studentSlice{
		{"李", 99},
		{"我", 79},
		{"啊", 80},
		{"额", 48},
	}
	sort.Sort(sli4)
	fmt.Println(sli4)
	// [{额 48} {我 79} {啊 80} {李 99}]
}

/*
	总结
	1.  逆向排序不是sort.Reverse(sort.Sort(sort.StringSlice(sli3)))
				而是sort.Sort(sort.Reverse(sort.StringSlice(sli3)))
	2.  对结构体进行排序的时候,需要先把结构体切片类型定义为一个名字,再重写Len,Less,Swap三个方法,
		之后直接进行排序就可以了(详情看上面)
*/
