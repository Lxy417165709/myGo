package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {

	if len(nums)==0{
		return "0"
	}

	strSlice:=make([]string,0)
	for i:=0;i<len(nums);i++{
		strSlice =append(strSlice,strconv.Itoa(nums[i]))
	}
	sort.Slice(strSlice,func(i,j int)bool{
		return judge(strSlice[i]+strSlice[j],strSlice[j]+strSlice[i])
	})

	ans:=strings.Join(strSlice,"")

	if ans[0]=='0'{
		return "0"
	}
	return ans
}


func judge(a string,b string) bool{


	for i:=0;i<len(a);i++{
		if a[i]!=b[i]{
			return a[i]>b[i]
		}
	}
	return a == b
}

func min(a,b int) int{
	if a>b{
		return b
	}
	return a
}

/*
	总结
	1. 第一次的话，比较函数写错了，每次比较的是2个数字字符串的字典序，然后返回短的。89,8 显然这是个反例。
	2. 之后想了一下，只需要比较2个数拼成的数大小就可以了，于是写出了比较函数,AC了这个题目


*/
