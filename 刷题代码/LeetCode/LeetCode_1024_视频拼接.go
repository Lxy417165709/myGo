package main

import (
	"fmt"
	"sort"
)

type arr [][]int

func (a arr)Len() int{
	return len(a)
}

func (a arr)Less(i,j int) bool{
	if a[i][0]==a[j][0] {
		return a[i][1]>a[j][1]
	}else{
		return a[i][0]<a[j][0]
	}
}
func (a arr)Swap(i,j int) {

	a[i],a[j] = a[j],a[i]
}

func videoStitching(clips [][]int, T int) int {
	sort.Sort(arr(clips))
	slice := make([][]int,0)
	l,r:=0,0


	for i:=0;i< len(clips);i++  {
		if i==0 {
			if clips[i][0]!=0 {
				return -1
			}else{
				l = 0
				r = clips[i][1]
				slice =append(slice,[]int{l,r} )
			}
		}else{
			mx := r
			index:=-1
			// 注意这个循环,通过这个循环我们可以在[0,4],[1,5],[3,9],得出{[0,4][3,9]}而不是{[0,4],[1,5],[3,9]}
			// 第一次AC的时候是把i赋予t，然后递增t,但是这样会多了很多不必要的判断,比如[0,5]已经判断过了，但i是0,所以下一个循环
			// 开始判断i==1,而不是i==6
			for ;  i<len(clips) && r>=clips[i][0];i++  {

				if clips[i][1]>mx {
					index = i
					mx = clips[i][1]
				}
			}
			if index!=-1 {
				r = mx
				slice =append(slice,[]int{clips[index][0],clips[index][1]} )
			}

		}
		if r>=T {
			return len(slice)
		}
	}

	return -1



}


func main() {
	fmt.Println(videoStitching([][]int{
		{0,1},
		{1,2},


	},3))

}


/*
	总结
	1. 该题类似于区间合并，注意点在代码中写出了
	2. 有人用DP做出来了~
*/